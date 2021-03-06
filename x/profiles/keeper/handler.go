package keeper

import (
	"fmt"
	"regexp"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewHandler returns a handler for "profile" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgSaveProfile:
			return handleMsgSaveProfile(ctx, keeper, msg)
		case types.MsgDeleteProfile:
			return handleMsgDeleteProfile(ctx, keeper, msg)
		case types.MsgRequestDTagTransfer:
			return handleMsgRequestDTagTransfer(ctx, keeper, msg)
		case types.MsgAcceptDTagTransferRequest:
			return handleMsgAcceptDTagTransfer(ctx, keeper, msg)
		case types.MsgRefuseDTagTransferRequest:
			return handleMsgRefuseDTagTransfer(ctx, keeper, msg)
		case types.MsgCancelDTagTransferRequest:
			return handleMsgCancelDTagTransfer(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Profiles message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// ValidateProfile checks if the given profile is valid according to the current profile's module params
func ValidateProfile(ctx sdk.Context, keeper Keeper, profile types.Profile) error {
	params := keeper.GetParams(ctx)

	minMonikerLen := params.MonikerParams.MinMonikerLen.Int64()
	maxMonikerLen := params.MonikerParams.MaxMonikerLen.Int64()

	if profile.Moniker != nil {
		nameLen := int64(len(*profile.Moniker))
		if nameLen < minMonikerLen {
			return fmt.Errorf("profile moniker cannot be less than %d characters", minMonikerLen)
		}
		if nameLen > maxMonikerLen {
			return fmt.Errorf("profile moniker cannot exceed %d characters", maxMonikerLen)
		}
	}

	dTagRegEx := regexp.MustCompile(params.DtagParams.RegEx)
	minDtagLen := params.DtagParams.MinDtagLen.Int64()
	maxDtagLen := params.DtagParams.MaxDtagLen.Int64()
	dtagLen := int64(len(profile.DTag))

	if !dTagRegEx.MatchString(profile.DTag) {
		return fmt.Errorf("invalid profile dtag, it should match the following regEx %s", dTagRegEx)
	}

	if dtagLen < minDtagLen {
		return fmt.Errorf("profile dtag cannot be less than %d characters", minDtagLen)
	}

	if dtagLen > maxDtagLen {
		return fmt.Errorf("profile dtag cannot exceed %d characters", maxDtagLen)
	}

	maxBioLen := params.MaxBioLen.Int64()
	if profile.Bio != nil && int64(len(*profile.Bio)) > maxBioLen {
		return fmt.Errorf("profile biography cannot exceed %d characters", maxBioLen)
	}

	if err := profile.Validate(); err != nil {
		return err
	}

	return nil
}

// handleMsgSaveProfile handles the creation/edit of a profile
func handleMsgSaveProfile(ctx sdk.Context, keeper Keeper, msg types.MsgSaveProfile) (*sdk.Result, error) {
	profile, found := keeper.GetProfile(ctx, msg.Creator)

	// Create a new profile if not found
	if !found {
		profile = types.NewProfile(msg.Dtag, msg.Creator, ctx.BlockTime())
	}

	// If the DTag changes, delete all the previous DTag transfer requests
	if profile.DTag != msg.Dtag {
		keeper.DeleteAllDTagTransferRequests(ctx, msg.Creator)
	}

	// Replace all editable fields (clients should autofill existing values)
	// We do not replace the tag since we do not want it to be editable
	profile = profile.
		WithDTag(msg.Dtag).
		WithMoniker(msg.Moniker).
		WithBio(msg.Bio).
		WithPictures(msg.ProfilePic, msg.CoverPic)
	err := ValidateProfile(ctx, keeper, profile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the profile
	if err := keeper.SaveProfile(ctx, profile); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeProfileSaved,
		sdk.NewAttribute(types.AttributeProfileDtag, profile.DTag),
		sdk.NewAttribute(types.AttributeProfileCreator, profile.Creator.String()),
		sdk.NewAttribute(types.AttributeProfileCreationTime, profile.CreationDate.Format(time.RFC3339)),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(profile.DTag),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgDeleteProfile handles the deletion of a profile
func handleMsgDeleteProfile(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteProfile) (*sdk.Result, error) {
	profile, found := keeper.GetProfile(ctx, msg.Creator)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("no profile associated with this address: %s", msg.Creator))
	}

	keeper.DeleteProfile(ctx, profile.Creator, profile.DTag)

	createEvent := sdk.NewEvent(
		types.EventTypeProfileDeleted,
		sdk.NewAttribute(types.AttributeProfileDtag, profile.DTag),
		sdk.NewAttribute(types.AttributeProfileCreator, profile.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(profile.DTag),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgRequestDTagTransfer handles the request of a dTag transfer
func handleMsgRequestDTagTransfer(ctx sdk.Context, keeper Keeper, msg types.MsgRequestDTagTransfer) (*sdk.Result, error) {
	// check if the request's receiver has blocked the sender before
	if keeper.IsUserBlocked(ctx, msg.Receiver, msg.Sender) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("The user with address %s has blocked you", msg.Receiver))
	}

	dtagToTrade := keeper.GetDtagFromAddress(ctx, msg.Receiver)
	if len(dtagToTrade) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("The user with address %s doesn't have a profile yet so their DTag cannot be transferred",
				msg.Receiver))
	}
	transferRequest := types.NewDTagTransferRequest(dtagToTrade, msg.Receiver, msg.Sender)

	if err := keeper.SaveDTagTransferRequest(ctx, transferRequest); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferRequest,
		sdk.NewAttribute(types.AttributeDTagToTrade, dtagToTrade),
		sdk.NewAttribute(types.AttributeRequestReceiver, transferRequest.Receiver.String()),
		sdk.NewAttribute(types.AttributeRequestSender, transferRequest.Sender.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(transferRequest),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgAcceptDTagTransfer handles the acceptance of a dTag transfer request
func handleMsgAcceptDTagTransfer(ctx sdk.Context, keeper Keeper, msg types.MsgAcceptDTagTransferRequest) (*sdk.Result, error) {
	requests := keeper.GetUserDTagTransferRequests(ctx, msg.Receiver)

	// Check if the receiving user request is present, if not return error
	found := false
	var dTagWanted string
	for _, req := range requests {
		if req.Sender.Equals(msg.Sender) {
			dTagWanted = req.DTagToTrade
			found = true
			break
		}
	}

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("no request made from %s", msg.Sender))
	}

	// Get the current owner profile
	currentOwnerProfile, exist := keeper.GetProfile(ctx, msg.Receiver)
	if !exist {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("profile of %s doesn't exist", msg.Receiver))
	}

	// Get the DTag to trade
	dTagToTrade := currentOwnerProfile.DTag

	if dTagWanted != dTagToTrade {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner's DTag is different from the one to be exchanged")
	}

	// Save the current owner profile with his new dTag
	currentOwnerProfile = currentOwnerProfile.WithDTag(msg.NewDTag)
	err := keeper.SaveProfile(ctx, currentOwnerProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// check for an existent profile of the receiving user
	receiverProfile, exist := keeper.GetProfile(ctx, msg.Sender)
	if !exist {
		receiverProfile = types.NewProfile(dTagToTrade, msg.Sender, ctx.BlockTime())
	} else {
		receiverProfile = receiverProfile.WithDTag(dTagToTrade)
	}

	err = keeper.SaveProfile(ctx, receiverProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	keeper.DeleteAllDTagTransferRequests(ctx, msg.Receiver)
	keeper.DeleteAllDTagTransferRequests(ctx, msg.Sender)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferAccept,
		sdk.NewAttribute(types.AttributeDTagToTrade, dTagToTrade),
		sdk.NewAttribute(types.AttributeNewDTag, msg.NewDTag),
		sdk.NewAttribute(types.AttributeRequestReceiver, msg.Receiver.String()),
		sdk.NewAttribute(types.AttributeRequestSender, msg.Sender.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(dTagToTrade),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil

}

// DeleteDTagTransferRequest delete the request made by the sender and returns the event with the given eventType.
func DeleteDTagTransferRequest(ctx sdk.Context, keeper Keeper, owner, sender sdk.AccAddress, eventType string) (*sdk.Result, error) {
	if err := keeper.DeleteDTagTransferRequest(ctx, owner, sender); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		eventType,
		sdk.NewAttribute(types.AttributeRequestReceiver, owner.String()),
		sdk.NewAttribute(types.AttributeRequestSender, sender.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(sender),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgRefuseDTagTransfer handles the reject of a dTag transfer request
func handleMsgRefuseDTagTransfer(ctx sdk.Context, keeper Keeper, msg types.MsgRefuseDTagTransferRequest) (*sdk.Result, error) {
	return DeleteDTagTransferRequest(ctx, keeper, msg.Sender, msg.Receiver, types.EventTypeDTagTransferRefuse)
}

// handleMsgCancelDTagTransfer handles the deletion of a dTag transfer request
func handleMsgCancelDTagTransfer(ctx sdk.Context, keeper Keeper, msg types.MsgCancelDTagTransferRequest) (*sdk.Result, error) {
	return DeleteDTagTransferRequest(ctx, keeper, msg.Receiver, msg.Sender, types.EventTypeDTagTransferCancel)
}
