// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package ante

import (
	"math"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	ethermint "github.com/evmos/ethermint/types"
	"github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// VerifyEthAccount validates checks that the sender balance is greater than the total transaction cost.
// The account will be set to store if it doesn't exis, i.e cannot be found on store.
// This AnteHandler decorator will fail if:
// - any of the msgs is not a MsgEthereumTx
// - from address is empty
// - account balance is lower than the transaction cost
func VerifyEthAccount(
	ctx sdk.Context, tx sdk.Tx,
	evmKeeper EVMKeeper, ak evmtypes.AccountKeeper, evmDenom string,
) error {
	if !ctx.IsCheckTx() {
		return nil
	}

	for i, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return errorsmod.Wrapf(errortypes.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return errorsmod.Wrapf(err, "failed to unpack tx data any for tx %d", i)
		}

		// sender address should be in the tx cache from the previous AnteHandle call
		from := msgEthTx.GetFrom()
		if from.Empty() {
			return errorsmod.Wrap(errortypes.ErrInvalidAddress, "from address cannot be empty")
		}

		// check whether the sender address is EOA
		fromAddr := common.BytesToAddress(from)
		acct := evmKeeper.GetAccount(ctx, fromAddr)

		if acct == nil {
			acc := ak.NewAccountWithAddress(ctx, from)
			ak.SetAccount(ctx, acc)
		} else if acct.IsContract() {
			return errorsmod.Wrapf(errortypes.ErrInvalidType,
				"the sender is not EOA: address %s, codeHash <%s>", fromAddr, acct.CodeHash)
		}

		balance := evmKeeper.GetBalance(ctx, sdk.AccAddress(fromAddr.Bytes()), evmDenom)
		if err := keeper.CheckSenderBalance(sdkmath.NewIntFromBigInt(balance), txData); err != nil {
			return errorsmod.Wrap(err, "failed to check sender balance")
		}
	}
	return nil
}

// CheckEthGasConsume validates that the Ethereum tx message has enough to cover intrinsic gas
// (during CheckTx only) and that the sender has enough balance to pay for the gas cost.
//
// Intrinsic gas for a transaction is the amount of gas that the transaction uses before the
// transaction is executed. The gas is a constant value plus any cost incurred by additional bytes
// of data supplied with the transaction.
//
// This AnteHandler decorator will fail if:
// - the message is not a MsgEthereumTx
// - sender account cannot be found
// - transaction's gas limit is lower than the intrinsic gas
// - user doesn't have enough balance to deduct the transaction fees (gas_limit * gas_price)
// - transaction or block gas meter runs out of gas
// - sets the gas meter limit
// - gas limit is greater than the block gas meter limit
func CheckEthGasConsume(
	ctx sdk.Context, tx sdk.Tx,
	ethCfg *params.ChainConfig,
	evmKeeper EVMKeeper,
	baseFee *big.Int,
	maxGasWanted uint64,
	evmDenom string,
) (sdk.Context, error) {
	gasWanted := uint64(0)

	blockHeight := big.NewInt(ctx.BlockHeight())
	homestead := ethCfg.IsHomestead(blockHeight)
	istanbul := ethCfg.IsIstanbul(blockHeight)
	shanghai := ethCfg.IsShanghai(big.NewInt(ctx.BlockHeader().Height), uint64(ctx.BlockHeader().Time.Unix()))
	var events sdk.Events

	// Use the lowest priority of all the messages as the final one.
	minPriority := int64(math.MaxInt64)

	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errorsmod.Wrapf(errortypes.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return ctx, errorsmod.Wrap(err, "failed to unpack tx data")
		}

		priority := evmtypes.GetTxPriority(txData, baseFee)

		if priority < minPriority {
			minPriority = priority
		}

		if ctx.IsCheckTx() && maxGasWanted != 0 {
			// We can't trust the tx gas limit, because we'll refund the unused gas.
			if txData.GetGas() > maxGasWanted {
				gasWanted += maxGasWanted
			} else {
				gasWanted += txData.GetGas()
			}
		} else {
			gasWanted += txData.GetGas()
		}

		// user balance is already checked during CheckTx so there's no need to
		// verify it again during ReCheckTx
		if ctx.IsReCheckTx() {
			continue
		}

		fees, err := keeper.VerifyFee(txData, evmDenom, baseFee, homestead, istanbul, shanghai, ctx.IsCheckTx())
		if err != nil {
			return ctx, errorsmod.Wrapf(err, "failed to verify the fees")
		}

		err = evmKeeper.DeductTxCostsFromUserBalance(ctx, fees, common.BytesToAddress(msgEthTx.From))
		if err != nil {
			return ctx, errorsmod.Wrapf(err, "failed to deduct transaction costs from user balance")
		}

		events = append(events,
			sdk.NewEvent(
				sdk.EventTypeTx,
				sdk.NewAttribute(sdk.AttributeKeyFee, fees.String()),
			),
		)
	}

	ctx.EventManager().EmitEvents(events)

	blockGasLimit := ethermint.BlockGasLimit(ctx)

	// return error if the tx gas is greater than the block limit (max gas)

	// NOTE: it's important here to use the gas wanted instead of the gas consumed
	// from the tx gas pool. The later only has the value so far since the
	// EthSetupContextDecorator so it will never exceed the block gas limit.
	if gasWanted > blockGasLimit {
		return ctx, errorsmod.Wrapf(
			errortypes.ErrOutOfGas,
			"tx gas (%d) exceeds block gas limit (%d)",
			gasWanted,
			blockGasLimit,
		)
	}

	// Set tx GasMeter with a limit of GasWanted (i.e gas limit from the Ethereum tx).
	// The gas consumed will be then reset to the gas used by the state transition
	// in the EVM.

	// FIXME: use a custom gas configuration that doesn't add any additional gas and only
	// takes into account the gas consumed at the end of the EVM transaction.
	newCtx := ctx.
		WithGasMeter(ethermint.NewInfiniteGasMeterWithLimit(gasWanted)).
		WithPriority(minPriority)

	// we know that we have enough gas on the pool to cover the intrinsic gas
	return newCtx, nil
}

// CheckEthCanTransfer creates an EVM from the message and calls the BlockContext CanTransfer function to
// see if the address can execute the transaction.
func CheckEthCanTransfer(
	ctx sdk.Context, tx sdk.Tx,
	baseFee *big.Int,
	ethCfg *params.ChainConfig,
	evmKeeper EVMKeeper,
	evmParams *evmtypes.Params,
) error {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return errorsmod.Wrapf(errortypes.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		coreMsg, err := msgEthTx.AsMessage(baseFee)
		if err != nil {
			return errorsmod.Wrapf(
				err,
				"failed to create an ethereum core.Message",
			)
		}

		if evmtypes.IsLondon(ethCfg, ctx.BlockHeight()) {
			if baseFee == nil {
				return errorsmod.Wrap(
					evmtypes.ErrInvalidBaseFee,
					"base fee is supported but evm block context value is nil",
				)
			}
			if coreMsg.GasFeeCap.Cmp(baseFee) < 0 {
				return errorsmod.Wrapf(
					errortypes.ErrInsufficientFee,
					"max fee per gas less than block base fee (%s < %s)",
					coreMsg.GasFeeCap, baseFee,
				)
			}
		}

		// check that caller has enough balance to cover asset transfer for **topmost** call
		// NOTE: here the gas consumed is from the context with the infinite gas meter
		if coreMsg.Value.Sign() > 0 && !canTransfer(ctx, evmKeeper, evmParams.EvmDenom, coreMsg.From, coreMsg.Value) {
			return errorsmod.Wrapf(
				errortypes.ErrInsufficientFunds,
				"failed to transfer %s from address %s using the EVM block context transfer function",
				coreMsg.Value,
				coreMsg.From,
			)
		}
	}

	return nil
}

// canTransfer adapted the core.CanTransfer from go-ethereum
func canTransfer(ctx sdk.Context, evmKeeper EVMKeeper, denom string, from common.Address, amount *big.Int) bool {
	balance := evmKeeper.GetBalance(ctx, sdk.AccAddress(from.Bytes()), denom)
	return balance.Cmp(amount) >= 0
}

// CheckEthSenderNonce handles incrementing the sequence of the signer (i.e sender). If the transaction is a
// contract creation, the nonce will be incremented during the transaction execution and not within
// this AnteHandler decorator.
func CheckEthSenderNonce(
	ctx sdk.Context, tx sdk.Tx, ak evmtypes.AccountKeeper,
) error {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return errorsmod.Wrapf(errortypes.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return errorsmod.Wrap(err, "failed to unpack tx data")
		}

		// increase sequence of sender
		acc := ak.GetAccount(ctx, msgEthTx.GetFrom())
		if acc == nil {
			return errorsmod.Wrapf(
				errortypes.ErrUnknownAddress,
				"account %s is nil", common.BytesToAddress(msgEthTx.GetFrom().Bytes()),
			)
		}
		nonce := acc.GetSequence()

		// we merged the nonce verification to nonce increment, so when tx includes multiple messages
		// with same sender, they'll be accepted.
		if txData.GetNonce() != nonce {
			return errorsmod.Wrapf(
				errortypes.ErrInvalidSequence,
				"invalid nonce; got %d, expected %d", txData.GetNonce(), nonce,
			)
		}

		if err := acc.SetSequence(nonce + 1); err != nil {
			return errorsmod.Wrapf(err, "failed to set sequence to %d", acc.GetSequence()+1)
		}

		ak.SetAccount(ctx, acc)
	}

	return nil
}
