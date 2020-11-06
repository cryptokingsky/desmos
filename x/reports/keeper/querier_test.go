package keeper_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) Test_queryReports() {
	tests := []struct {
		name          string
		path          []string
		storedReports []types.Report
		expErr        error
		expResponse   []types.Report
	}{
		{
			name:          "Invalid Post ID",
			path:          []string{types.QueryReports, "1234"},
			storedReports: nil,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid postID: 1234"),
		},
		{
			name: "Valid request returns correctly",
			path: []string{types.QueryReports, suite.testData.postID},
			storedReports: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.creator,
				),
				types.NewReport(
					"other_post",
					"type",
					"message",
					suite.testData.creator,
				),
			},
			expResponse: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.creator,
				),
			},
		},
		{
			name:          "Empty stored and valid ID",
			path:          []string{types.QueryReports, suite.testData.postID},
			storedReports: nil,
			expResponse:   nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, rep := range test.storedReports {
				err := suite.keeper.SaveReport(suite.ctx, rep)
				suite.Require().NoError(err)
			}

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Require().Nil(err)

				var reports []types.Report
				err := suite.legacyAminoCdc.UnmarshalJSON(result, &reports)
				suite.Require().NoError(err)
				suite.Require().Equal(test.expResponse, reports)
			}

			if result == nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}
		})
	}
}
