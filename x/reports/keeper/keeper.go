package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc      codec.BinaryMarshaler // The wire codec for binary encoding/decoding.

	postKeeper postskeeper.Keeper // Post's keeper to perform checks on the postIDs
}

// NewKeeper creates new instances of the stored Keeper
func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey, pk postskeeper.Keeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		postKeeper: pk,
	}
}

// CheckPostExistence checks if a post with the given postID is present inside
// the current context and returns a boolean indicating that.
func (k Keeper) CheckPostExistence(ctx sdk.Context, postID string) bool {
	_, exist := k.postKeeper.GetPost(ctx, postID)
	return exist
}

// SaveReport allows to save the given stored inside the current context.
// It assumes that the given stored has already been validated.
// If the same stored has already been inserted, nothing will be changed.
func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) error {
	store := ctx.KVStore(k.storeKey)
	key := types.ReportStoreKey(report.PostId)

	// Get the list of stored related to the given postID
	var wrapped WrappedReports
	err := k.cdc.UnmarshalBinaryBare(store.Get(key), &wrapped)
	if err != nil {
		return err
	}

	// Append the given report
	if newWrapped, appended := wrapped.AppendIfMissing(report); appended {
		bz, err := k.cdc.MarshalBinaryBare(&newWrapped)
		if err != nil {
			return err
		}

		store.Set(key, bz)
	}

	return nil
}

// GetPostReports returns the list of stored associated with the given postID.
// If no stored is associated with the given postID the function will returns an empty list.
func (k Keeper) GetPostReports(ctx sdk.Context, postID string) ([]types.Report, error) {
	store := ctx.KVStore(k.storeKey)

	var wrapped WrappedReports
	err := k.cdc.UnmarshalBinaryBare(store.Get(types.ReportStoreKey(postID)), &wrapped)
	if err != nil {
		return nil, err
	}

	return wrapped.Reports, nil
}

// GetReports returns the list of all the stored that have been stored inside the given context
func (k Keeper) GetReports(ctx sdk.Context) ([]types.Report, error) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ReportsStorePrefix)
	defer iterator.Close()

	var reports []types.Report
	for ; iterator.Valid(); iterator.Next() {
		var wrapped WrappedReports
		err := k.cdc.UnmarshalBinaryBare(iterator.Value(), &wrapped)
		if err != nil {
			return nil, err
		}

		reports = append(reports, wrapped.Reports...)
	}

	return reports, nil
}
