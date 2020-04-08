package msgs

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models"
)

// ----------------------
// --- MsgAddPostReaction
// ----------------------

// MsgAddPostReaction defines the message to be used to add a reaction to a post
type MsgAddPostReaction struct {
	PostID models.PostID  `json:"post_id"` // Id of the post to react to
	Value  string         `json:"value"`   // Value of the reaction
	User   sdk.AccAddress `json:"user"`    // Address of the user reacting to the post
}

// NewMsgAddPostReaction is a constructor function for MsgAddPostReaction
func NewMsgAddPostReaction(postID models.PostID, value string, user sdk.AccAddress) MsgAddPostReaction {
	return MsgAddPostReaction{
		PostID: postID,
		User:   user,
		Value:  value,
	}
}

// Route should return the name of the module
func (msg MsgAddPostReaction) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgAddPostReaction) Type() string { return models.ActionAddPostReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddPostReaction) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id")
	}

	if msg.User.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid user address: %s", msg.User))
	}

	if !models.ShortCodeRegEx.MatchString(msg.Value) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Reaction short code must be an emoji short code")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddPostReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAddPostReaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.User}
}

// ----------------------
// --- MsgRemovePostReaction
// ----------------------

// MsgRemovePostReaction defines the message to be used when wanting to remove
// an existing reaction from a specific user having a specific value
type MsgRemovePostReaction struct {
	PostID   models.PostID  `json:"post_id"`  // Id of the post to unlike
	User     sdk.AccAddress `json:"user"`     // Address of the user that has previously liked the post
	Reaction string         `json:"reaction"` // Value of the reaction to be removed
}

// MsgUnlikePostPost is the constructor of MsgRemovePostReaction
func NewMsgRemovePostReaction(postID models.PostID, user sdk.AccAddress, reaction string) MsgRemovePostReaction {
	return MsgRemovePostReaction{
		PostID:   postID,
		User:     user,
		Reaction: reaction,
	}
}

// Route should return the name of the module
func (msg MsgRemovePostReaction) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgRemovePostReaction) Type() string { return models.ActionRemovePostReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgRemovePostReaction) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id")
	}

	if msg.User.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid user address: %s", msg.User))
	}

	if len(strings.TrimSpace(msg.Reaction)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Reaction value cannot be empty nor blank")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRemovePostReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRemovePostReaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.User}
}