package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

func MustUnmarshalRelationships(codec codec.BinaryMarshaler, bz []byte) []types.Relationship {
	var wrapped WrappedRelationships
	err := codec.UnmarshalBinaryBare(bz, &wrapped)
	if err != nil {
		panic(err)
	}
	return wrapped.Relationships
}

func MustMarshalRelationships(cdc codec.BinaryMarshaler, relationships []types.Relationship) []byte {
	wrapped := WrappedRelationships{Relationships: relationships}
	bz, err := cdc.MarshalBinaryBare(&wrapped)
	if err != nil {
		panic(err)
	}
	return bz
}
