package types

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/internal/types/models"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/polls"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/reactions"
	"github.com/desmos-labs/desmos/x/posts/internal/types/msgs"
)

const (
	ModuleName                      = common.ModuleName
	RouterKey                       = common.RouterKey
	StoreKey                        = common.StoreKey
	MaxPostMessageLength            = common.MaxPostMessageLength
	MaxOptionalDataFieldsNumber     = common.MaxOptionalDataFieldsNumber
	MaxOptionalDataFieldValueLength = common.MaxOptionalDataFieldValueLength
	ActionCreatePost                = common.ActionCreatePost
	ActionEditPost                  = common.ActionEditPost
	ActionAnswerPoll                = common.ActionAnswerPoll
	ActionAddPostReaction           = common.ActionAddPostReaction
	ActionRemovePostReaction        = common.ActionRemovePostReaction
	ActionRegisterReaction          = common.ActionRegisterReaction
	QuerierRoute                    = common.QuerierRoute
	QueryPost                       = common.QueryPost
	QueryPosts                      = common.QueryPosts
	QueryPollAnswers                = common.QueryPollAnswers
	QueryRegisteredReactions        = common.QueryRegisteredReactions
	PostSortByCreationDate          = common.PostSortByCreationDate
	PostSortByID                    = common.PostSortByID
	PostSortOrderAscending          = common.PostSortOrderAscending
	PostSortOrderDescending         = common.PostSortOrderDescending
)

var (
	// functions aliases
	ComputeID                  = models.ComputeID
	ParsePostID                = models.ParsePostID
	NewPost                    = models.NewPost
	NewPostResponse            = models.NewPostResponse
	PostStoreKey               = models.PostStoreKey
	PostCommentsStoreKey       = models.PostCommentsStoreKey
	PostReactionsStoreKey      = models.PostReactionsStoreKey
	ReactionsStoreKey          = models.ReactionsStoreKey
	PollAnswersStoreKey        = models.PollAnswersStoreKey
	RegisterModelsCodec        = models.RegisterModelsCodec
	NewPostMedia               = common.NewPostMedia
	ValidateURI                = common.ValidateURI
	NewPostMedias              = common.NewPostMedias
	GetEmojiByShortCodeOrValue = common.GetEmojiByShortCodeOrValue
	ParseAnswerID              = polls.ParseAnswerID
	NewPollAnswer              = polls.NewPollAnswer
	NewPollAnswers             = polls.NewPollAnswers
	NewPollData                = polls.NewPollData
	ArePollDataEquals          = polls.ArePollDataEquals
	NewUserAnswer              = polls.NewUserAnswer
	NewUserAnswers             = polls.NewUserAnswers
	NewPostReaction            = reactions.NewPostReaction
	NewPostReactions           = reactions.NewPostReactions
	NewReaction                = reactions.NewReaction
	IsEmoji                    = reactions.IsEmoji
	NewReactions               = reactions.NewReactions
	NewMsgAddPostReaction      = msgs.NewMsgAddPostReaction
	NewMsgRemovePostReaction   = msgs.NewMsgRemovePostReaction
	NewMsgAnswerPoll           = msgs.NewMsgAnswerPoll
	NewMsgCreatePost           = msgs.NewMsgCreatePost
	NewMsgEditPost             = msgs.NewMsgEditPost
	NewMsgRegisterReaction     = msgs.NewMsgRegisterReaction
	RegisterMessagesCodec      = msgs.RegisterMessagesCodec

	// variable aliases
	ModelsCdc                = models.ModelsCdc
	Sha256RegEx              = common.Sha256RegEx
	HashtagRegEx             = common.HashtagRegEx
	ShortCodeRegEx           = common.ShortCodeRegEx
	URIRegEx                 = common.URIRegEx
	ModuleAddress            = common.ModuleAddress
	PostStorePrefix          = common.PostStorePrefix
	PostCommentsStorePrefix  = common.PostCommentsStorePrefix
	PostReactionsStorePrefix = common.PostReactionsStorePrefix
	ReactionsStorePrefix     = common.ReactionsStorePrefix
	PollAnswersStorePrefix   = common.PollAnswersStorePrefix
	MsgsCodec                = msgs.MsgsCodec
)

type (
	PostMedia                = common.PostMedia
	PostMedias               = common.PostMedias
	OptionalData             = common.OptionalData
	KeyValue                 = common.KeyValue
	AnswerID                 = polls.AnswerID
	PollAnswer               = polls.PollAnswer
	PollAnswers              = polls.PollAnswers
	PollData                 = polls.PollData
	UserAnswer               = polls.UserAnswer
	UserAnswers              = polls.UserAnswers
	PostReaction             = reactions.PostReaction
	PostReactions            = reactions.PostReactions
	Reaction                 = reactions.Reaction
	Reactions                = reactions.Reactions
	MsgAddPostReaction       = msgs.MsgAddPostReaction
	MsgRemovePostReaction    = msgs.MsgRemovePostReaction
	MsgAnswerPoll            = msgs.MsgAnswerPoll
	MsgCreatePost            = msgs.MsgCreatePost
	MsgEditPost              = msgs.MsgEditPost
	MsgRegisterReaction      = msgs.MsgRegisterReaction
	PostID                   = models.PostID
	PostIDs                  = models.PostIDs
	Post                     = models.Post
	Posts                    = models.Posts
	PostQueryResponse        = models.PostQueryResponse
	PollAnswersQueryResponse = models.PollAnswersQueryResponse
)
