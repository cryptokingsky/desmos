package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	parentPost := types.Post{
		PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		Message:      "Post without medias",
		Created:      suite.testData.post.Created,
		LastEdited:   time.Time{},
		Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData: nil,
		Creator:      suite.testData.post.Creator,
		Attachments:  suite.testData.post.Attachments,
		PollData:     suite.testData.post.PollData,
	}

	commentPost := types.Post{
		PostID:         "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
		ParentID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		Message:        "Post without medias",
		AllowsComments: false,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   nil,
		Created:        suite.testData.post.Created.Add(time.Hour),
		Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	}

	postReaction := types.NewPostReaction(":like:", "+1", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	reaction := types.NewRegisteredReaction(suite.testData.post.Creator, ":like:", "+1", suite.testData.post.Subspace)
	answer := types.NewUserAnswer([]string{"1", "2"}, suite.testData.post.Creator)

	tests := []struct {
		name         string
		posts        types.Posts
		answers      *types.UserAnswer
		postReaction *types.PostReaction
		reaction     *types.RegisteredReaction
		expStop      bool
	}{
		{
			name:         "All invariants are not violated",
			posts:        types.Posts{parentPost, commentPost},
			answers:      &answer,
			postReaction: &postReaction,
			reaction:     &reaction,
			expStop:      true,
		},
		{
			name: "ValidPosts Invariants violated",
			posts: types.Posts{
				types.Post{
					PostID:       "1234",
					Message:      "Message",
					Created:      suite.testData.post.Created,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				}},
			answers:      nil,
			postReaction: nil,
			reaction:     nil,
			expStop:      true,
		},
		{
			name: "ValidCommentsDate Invariants violated",
			posts: types.Posts{parentPost,
				types.Post{
					PostID:       commentPost.PostID,
					ParentID:     parentPost.PostID,
					Message:      "Message",
					Created:      suite.testData.postEndPollDateExpired,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			answers:      nil,
			postReaction: nil,
			reaction:     nil,
			expStop:      true,
		},
		{
			name:         "ValidPostForReactions Invariants violated",
			posts:        types.Posts{},
			answers:      nil,
			postReaction: &postReaction,
			reaction:     &reaction,
			expStop:      true,
		},
		{
			name:         "ValidPollForPollAnswers Invariants violated",
			posts:        types.Posts{commentPost},
			answers:      &answer,
			postReaction: nil,
			reaction:     nil,
			expStop:      true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.DefaultParams())

			for _, post := range test.posts {
				suite.k.SavePost(suite.ctx, post)
			}

			if test.reaction != nil && test.postReaction != nil {
				suite.k.SaveRegisteredReaction(suite.ctx, *test.reaction)

				err := suite.k.SavePostReaction(suite.ctx, parentPost.PostID, *test.postReaction)
				suite.Require().NoError(err)
			}
			if test.answers != nil {
				suite.k.SavePollAnswers(suite.ctx, test.posts[0].PostID, *test.answers)
			}

			_, stop := keeper.AllInvariants(suite.k)(suite.ctx)
			suite.Require().Equal(test.expStop, stop)
		})
	}
}
