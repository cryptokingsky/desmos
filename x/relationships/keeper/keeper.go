package keeper

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc      codec.BinaryMarshaler // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// StoreRelationship allows to store the given relationship returning an error if he's already present.
func (k Keeper) StoreRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships := MustUnmarshalRelationships(k.cdc, store.Get(key))
	for _, rel := range relationships {
		if rel.Equal(relationship) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"relationship already exists with %s", relationship.Recipient)
		}
	}
	relationships = append(relationships, relationship)

	bz := MustMarshalRelationships(k.cdc, relationships)
	store.Set(key, bz)
	return nil
}

// GetUserRelationships allows to list all the stored that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user string) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		if relationship.Creator == user || relationship.Recipient == user {
			relationships = append(relationships, relationship)
		}
		return false
	})
	return relationships
}

// GetAllRelationships allows to returns the map of all the users and their associated stored
func (k Keeper) GetAllRelationships(ctx sdk.Context) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		relationships = append(relationships, relationship)
		return false
	})
	return relationships
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) DeleteRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	var wrapped WrappedRelationships
	err := k.cdc.UnmarshalBinaryBare(store.Get(key), &wrapped)
	if err != nil {
		return err
	}

	for index, rel := range wrapped.Relationships {
		if rel.Recipient == relationship.Recipient && rel.Subspace == relationship.Subspace {
			relationships := append(wrapped.Relationships[:index], wrapped.Relationships[index+1:]...)
			if len(relationships) == 0 {
				store.Delete(key)
			} else {
				bz, err := k.cdc.MarshalBinaryBare(&WrappedRelationships{Relationships: relationships})
				if err != nil {
					return err
				}

				store.Set(key, bz)
			}
			break
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// SaveUserBlock allows to store the given block inside the store, returning an error if
// something goes wrong.
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(userBlock.Blocker)

	var wrapped WrappedUserBlocks
	err := k.cdc.UnmarshalBinaryBare(store.Get(key), &wrapped)
	if err != nil {
		return err
	}

	for _, ub := range wrapped.Blocks {
		if ub == userBlock {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the user with address %s has already been blocked", userBlock.Blocked)
		}
	}

	wrapped = WrappedUserBlocks{Blocks: append(wrapped.Blocks, userBlock)}
	bz, err := k.cdc.MarshalBinaryBare(&wrapped)
	if err != nil {
		return err
	}

	store.Set(key, bz)
	return nil
}

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user.
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspace string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(blocker)

	var wrapped WrappedUserBlocks
	err := k.cdc.UnmarshalBinaryBare(store.Get(key), &wrapped)
	if err != nil {
		return err
	}

	for index, ub := range wrapped.Blocks {
		if ub.Blocker == blocker && ub.Blocked == blocked && ub.Subspace == subspace {
			usersBlocks := append(wrapped.Blocks[:index], wrapped.Blocks[index+1:]...)
			if len(usersBlocks) == 0 {
				store.Delete(key)
			} else {
				bz, err := k.cdc.MarshalBinaryBare(&WrappedUserBlocks{Blocks: usersBlocks})
				if err != nil {
					return err
				}
				store.Set(key, bz)
			}
			return nil
		}
	}

	return fmt.Errorf("blocked user with address %s not found", blocked)
}

// GetUserBlocks returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlocks(ctx sdk.Context, user string) ([]types.UserBlock, error) {
	store := ctx.KVStore(k.storeKey)

	var wrapped WrappedUserBlocks
	err := k.cdc.UnmarshalBinaryBare(store.Get(types.UsersBlocksStoreKey(user)), &wrapped)
	if err != nil {
		return nil, err
	}
	return wrapped.Blocks, nil
}

// GetUsersBlocks returns a list of all the users blocks inside the given context.
func (k Keeper) GetUsersBlocks(ctx sdk.Context) []types.UserBlock {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UsersBlocksStorePrefix)
	defer iterator.Close()

	var usersBlocks []types.UserBlock
	for ; iterator.Valid(); iterator.Next() {
		var wrapped WrappedUserBlocks
		err := k.cdc.UnmarshalBinaryBare(iterator.Value(), &wrapped)
		if err != nil {
			panic(err)
		}

		usersBlocks = append(usersBlocks, wrapped.Blocks...)
	}

	return usersBlocks
}

// CheckForBlockedUser checks if the given user address is present inside the blocked users array
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string) bool {
	blockedUsers, _ := k.GetUserBlocks(ctx, blocker)

	for _, user := range blockedUsers {
		if user.Blocked == blocked {
			return true
		}
	}

	return false
}
