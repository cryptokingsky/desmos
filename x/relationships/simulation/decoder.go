package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/desmos-labs/desmos/x/relationships/keeper"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// RelationshipsUnmarshaler defines the expected encoding store functions.
type RelationshipsUnmarshaler interface {
	UnmarshalRelationships([]byte) ([]types.Relationship, error)
	UnmarshalUserBlocks([]byte) ([]types.UserBlock, error)
}

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc codec.BinaryMarshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.RelationshipsStorePrefix):
			var relationshipsA, relationshipsB keeper.WrappedRelationships
			cdc.MustUnmarshalBinaryBare(kvA.Value, &relationshipsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &relationshipsB)
			return fmt.Sprintf("Relationships: %s\nRelationships: %s\n",
				relationshipsA.Relationships, relationshipsB.Relationships)

		case bytes.HasPrefix(kvA.Key, types.UsersBlocksStorePrefix):
			var userBlocksA, userBlocksB keeper.WrappedUserBlocks
			cdc.MustUnmarshalBinaryBare(kvA.Value, &userBlocksA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &userBlocksB)
			return fmt.Sprintf("UsersBlocks: %s\nUsersBlocks: %s\n",
				userBlocksA.Blocks, userBlocksB.Blocks)

		default:
			panic(fmt.Sprintf("invalid relationships key %X", kvA.Key))
		}
	}
}
