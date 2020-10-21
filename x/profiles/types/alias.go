package types

// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"github.com/desmos-labs/desmos/x/profiles/types/msgs"
)

const (
	ModuleName                = models.ModuleName
	RouterKey                 = models.RouterKey
	StoreKey                  = models.StoreKey
	ActionSaveProfile         = models.ActionSaveProfile
	ActionDeleteProfile       = models.ActionDeleteProfile
	ActionRequestDtag         = models.ActionRequestDtag
	ActionAcceptDtagTransfer  = models.ActionAcceptDtagTransfer
	RefuseDTagTransferRequest = models.RefuseDTagTransferRequest
	CancelDTagTransferRequest = models.CancelDTagTransferRequest
	QuerierRoute              = models.QuerierRoute
	QueryProfile              = models.QueryProfile
	QueryProfiles             = models.QueryProfiles
	QueryDTagRequests         = models.QueryDTagRequests
	QueryParams               = models.QueryParams
)

var (
	// functions aliases
	NewDTagTransferRequest          = models.NewDTagTransferRequest
	RegisterModelsCodec             = models.RegisterModelsCodec
	ProfileStoreKey                 = models.ProfileStoreKey
	DtagStoreKey                    = models.DtagStoreKey
	DtagTransferRequestStoreKey     = models.DtagTransferRequestStoreKey
	NewProfile                      = models.NewProfile
	NewProfiles                     = models.NewProfiles
	NewPictures                     = models.NewPictures
	NewMsgSaveProfile               = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile             = msgs.NewMsgDeleteProfile
	NewMsgRequestDTagTransfer       = msgs.NewMsgRequestDTagTransfer
	NewMsgAcceptDTagTransfer        = msgs.NewMsgAcceptDTagTransfer
	NewMsgRefuseDTagTransferRequest = msgs.NewMsgRefuseDTagTransferRequest
	NewMsgCancelDTagTransferRequest = msgs.NewMsgCancelDTagTransferRequest
	RegisterMessagesCodec           = msgs.RegisterMessagesCodec

	// variable aliases
	ModelsCdc                  = models.ModelsCdc
	ProfileStorePrefix         = models.ProfileStorePrefix
	DtagStorePrefix            = models.DtagStorePrefix
	DTagTransferRequestsPrefix = models.DTagTransferRequestsPrefix
	MsgsCodec                  = msgs.MsgsCodec
)

type (
	DTagTransferRequest          = models.DTagTransferRequest
	Profile                      = models.Profile
	Profiles                     = models.Profiles
	Pictures                     = models.Pictures
	MsgSaveProfile               = msgs.MsgSaveProfile
	MsgDeleteProfile             = msgs.MsgDeleteProfile
	MsgRequestDTagTransfer       = msgs.MsgRequestDTagTransfer
	MsgAcceptDTagTransferRequest = msgs.MsgAcceptDTagTransferRequest
	MsgRefuseDTagTransferRequest = msgs.MsgRefuseDTagTransferRequest
	MsgCancelDTagTransferRequest = msgs.MsgCancelDTagTransferRequest
)