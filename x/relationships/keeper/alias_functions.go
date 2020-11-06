package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// IterateRelationships iterates through the relationships and perform the provided function
func (k Keeper) IterateRelationships(ctx sdk.Context, fn func(index int64, relationship types.Relationship) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		var wrapped WrappedRelationships
		err := k.cdc.UnmarshalBinaryBare(iterator.Value(), &wrapped)
		if err != nil {
			panic(err)
		}

		var stop = false
		for _, relationship := range wrapped.Relationships {
			stop = fn(i, relationship)

			if stop {
				break
			}

			i++
		}

		if stop {
			break
		}
	}
}
