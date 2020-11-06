package types

import (
	"fmt"
	"strings"

	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
)

// NewReport returns a Report
func NewReport(postID string, reportType string, message string, user string) Report {
	return Report{
		PostId:  postID,
		Type:    reportType,
		Message: message,
		User:    user,
	}
}

// Validate implements validator
func (r Report) Validate() error {
	if !postsTypes.IsValidPostID(r.PostId) {
		return fmt.Errorf("invalid post id: %s", r.PostId)
	}

	if len(strings.TrimSpace(r.Type)) == 0 {
		return fmt.Errorf("report type cannot be empty")
	}

	if len(strings.TrimSpace(r.Message)) == 0 {
		return fmt.Errorf("report message cannot be empty")
	}

	if len(r.User) == 0 {
		return fmt.Errorf("invalid user address: %s", r.User)
	}

	return nil
}