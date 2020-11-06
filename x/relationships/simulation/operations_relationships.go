package simulation

// DONTCOVER

import (
	"math/rand"

	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// SimulateMsgCreateRelationship tests and runs a single msg create relationships
// nolint: funlen
func SimulateMsgCreateRelationship(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		sender, relationship, skip := randomRelationshipFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgCreateRelationship(relationship.Creator, relationship.Recipient, relationship.Subspace)
		err = sendMsgCreateRelationship(r, app, ak, bk, msg, ctx, chainID, []crypto.PrivKey{sender.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreateRelationship sends a transaction with a Relationship from a provided random account
func sendMsgCreateRelationship(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgCreateRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, err := simtypes.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)
	if err != nil {
		return err
	}

	_, _, err = app.Deliver(txGen.TxEncoder(), tx)
	if err != nil {
		return err
	}

	return nil
}

// randomRelationshipFields returns random relationships fields
func randomRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, types.Relationship, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, types.Relationship{}, true
	}

	// Get random accounts
	sender, _ := simtypes.RandomAcc(r, accs)
	receiver, _ := simtypes.RandomAcc(r, accs)

	subspace := RandomSubspace(r)

	// skip if the two relationship are equals
	if sender.Equals(receiver) {
		return simtypes.Account{}, types.Relationship{}, true
	}

	if k.IsUserBlocked(ctx, receiver.Address.String(), sender.Address.String()) {
		return simtypes.Account{}, types.Relationship{}, true
	}

	rel := types.NewRelationship(sender.Address.String(), receiver.Address.String(), subspace)

	// Skip if relationships already exists
	relationships := k.GetUserRelationships(ctx, sender.Address.String())
	for _, relationship := range relationships {
		if relationship.Equal(rel) {
			return simtypes.Account{}, types.Relationship{}, true
		}
	}

	return sender, rel, false
}

// SimulateMsgDeleteRelationship tests and runs a single msg delete relationships
// nolint: funlen
func SimulateMsgDeleteRelationship(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		user, counterparty, subspace, skip := randomDeleteRelationshipFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgDeleteRelationship(user.Address.String(), counterparty, subspace)
		err = sendMsgDeleteRelationship(r, app, ak, bk, msg, ctx, chainID, []crypto.PrivKey{user.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDeleteRelationship sends a transaction with a MsgDenyBidirectionalRelationship from a provided random account
func sendMsgDeleteRelationship(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgDeleteRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, err := simtypes.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)
	if err != nil {
		return err
	}

	_, _, err = app.Deliver(txGen.TxEncoder(), tx)
	if err != nil {
		return err
	}

	return nil
}

// randomDeleteRelationshipFields returns random delete relationships fields
func randomDeleteRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (user simtypes.Account, counterparty string, subspace string, skip bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, "", "", true
	}

	// Get a random account
	user, _ = simtypes.RandomAcc(r, accs)

	// Skip the test if the user has no relationships
	relationships := k.GetUserRelationships(ctx, user.Address.String())
	if len(relationships) == 0 {
		return simtypes.Account{}, "", "", true
	}

	// Get a randomg relationship
	relationship := RandomRelationship(r, relationships)

	// Get the counterparty
	if user.Address.String() == relationship.Creator {
		counterparty = relationship.Recipient
	} else {
		counterparty = relationship.Creator
	}

	return user, counterparty, relationship.Subspace, false
}
