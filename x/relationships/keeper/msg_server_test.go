package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreateRelationship() {
	tests := []struct {
		name                string
		isBlocked           bool
		msg                 *types.MsgCreateRelationship
		storedRelationships []types.Relationship
		expErr              error
		expEvents           sdk.Events
	}{
		{
			name: "Relationship sender been blocked from receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			isBlocked: true,
			expEvents: sdk.EmptyEvents(),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"The user with address cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns has blocked you",
			),
		},
		{
			name: "Relationship already created returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			storedRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expEvents: sdk.EmptyEvents(),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"relationship already exists with cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
		},
		{
			name: "Relationship has been saved correctly",
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipCreated,
					sdk.NewAttribute(types.AttributeRelationshipSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.storedRelationships {
				err := suite.k.StoreRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			if test.isBlocked {
				userBlock := types.NewUserBlock(test.msg.Receiver, test.msg.Sender, "test", "")
				err := suite.k.SaveUserBlock(suite.ctx, userBlock)
				suite.Require().NoError(err)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.CreateRelationship(sdk.WrapSDKContext(suite.ctx), test.msg)
			suite.RequireErrorsEqual(test.expErr, err)
			suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteRelationship() {
	store := suite.ctx.KVStore(suite.storeKey)

	// Store the initial relationships
	bz, err := suite.cdc.MarshalBinaryBare(&keeper.WrappedRelationships{Relationships: []types.Relationship{
		types.NewRelationship(
			suite.testData.user,
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRelationship(
			suite.testData.user,
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}})
	store.Set(types.RelationshipsStoreKey(suite.testData.user), bz)

	// Delete the relationship
	msg := types.NewMsgDeleteRelationship(
		suite.testData.user,
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	service := keeper.NewMsgServerImpl(suite.k)
	_, err = service.RemoveRelationship(sdk.WrapSDKContext(suite.ctx), msg)
	suite.Require().NoError(err)

	// Verify the remaining relationships
	expected := []types.Relationship{
		types.NewRelationship(
			suite.testData.user,
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}

	actual := suite.k.GetUserRelationships(suite.ctx, suite.testData.user)

	suite.Require().Equal(expected, actual)

	// Check the events
	suite.Len(suite.ctx.EventManager().Events(), 1)
	suite.Contains(suite.ctx.EventManager().Events(), sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, suite.testData.user),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
	))
}

func (suite *KeeperTestSuite) Test_handleMsgBlockUser() {
	tests := []struct {
		name             string
		msg              *types.MsgBlockUser
		storedUserBlocks []types.UserBlock
		expErr           error
		expEvent         sdk.Event
	}{
		{
			name: "Relationship already created returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)},
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"the user with address cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns has already been blocked",
			),
		},
		{
			name: "Relationship has been saved correctly",
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			storedUserBlocks: nil,
			expErr:           nil,
			expEvent: sdk.NewEvent(
				types.EventTypeBlockUser,
				sdk.NewAttribute(types.AttributeUserBlockBlocker, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				sdk.NewAttribute(types.AttributeUserBlockBlocked, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				sdk.NewAttribute(types.AttributeSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				sdk.NewAttribute(types.AttributeUserBlockReason, "reason"),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.storedUserBlocks {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.BlockUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.Require().NoError(err)

				// Check the events
				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), test.expEvent)

				blocks, err := suite.k.GetUserBlocks(suite.ctx, test.msg.Blocker)
				suite.Require().NoError(err)

				suite.Len(blocks, 1)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgUnblockUser() {
	store := suite.ctx.KVStore(suite.storeKey)

	// Store the existing blocks
	bz, err := suite.cdc.MarshalBinaryBare(&keeper.WrappedUserBlocks{Blocks: []types.UserBlock{
		types.NewUserBlock(
			suite.testData.user,
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewUserBlock(
			suite.testData.user,
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}})
	suite.Require().NoError(err)
	store.Set(types.UsersBlocksStoreKey(suite.testData.user), bz)

	// Unblock a user
	msg := types.NewMsgUnblockUser(
		suite.testData.user,
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	service := keeper.NewMsgServerImpl(suite.k)
	_, err = service.UnblockUser(sdk.WrapSDKContext(suite.ctx), msg)

	suite.Require().NoError(err)

	expected := []types.UserBlock{
		types.NewUserBlock(
			suite.testData.user,
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}

	userBlocks, err := suite.k.GetUserBlocks(suite.ctx, suite.testData.user)
	suite.Require().NoError(err)
	suite.Require().Equal(expected, userBlocks)

	// Check the events
	suite.Len(suite.ctx.EventManager().Events(), 1)
	suite.Contains(suite.ctx.EventManager().Events(), sdk.NewEvent(
		types.EventTypeUnblockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, suite.testData.user),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
		sdk.NewAttribute(types.AttributeSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
	))
}