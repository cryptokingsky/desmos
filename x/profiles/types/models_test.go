package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/profiles/types"

	"github.com/stretchr/testify/require"
)

func TestPictures_Validate(t *testing.T) {
	tests := []struct {
		name     string
		pictures types.Pictures
		expErr   error
	}{
		{
			name:     "Valid Pictures",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			expErr:   nil,
		},
		{
			name:     "Invalid Pictures profile uri",
			pictures: types.NewPictures("invalid", "https://shorturl.at/cgpyF"),
			expErr:   fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name:     "Invalid Pictures cover uri",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "invalid"),
			expErr:   fmt.Errorf("invalid profile cover uri provided"),
		},
	}

	for _, test := range tests {
		actErr := test.pictures.Validate()
		require.Equal(t, test.expErr, actErr)
	}
}

// ___________________________________________________________________________________________________________________

func TestProfile_Validate(t *testing.T) {
	tests := []struct {
		name    string
		account types.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Unix(100, 0),
				"",
			),
			expErr: fmt.Errorf("invalid creator address: "),
		},
		{
			name: "Empty profile DTag returns error",
			account: types.NewProfile(
				"",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile DTag: "),
		},
		{
			name: "Invalid profile pictures returns error",
			account: types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("pic", "cov"),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Valid profile returns no error",
			account: types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.account.Validate())
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestDTagTransferRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request types.DTagTransferRequest
		expErr  error
	}{
		{
			name: "Empty DTag to trade returns error",
			request: types.NewDTagTransferRequest(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("invalid DTag to trade "),
		},
		{
			name: "Empty request sender returns error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("invalid sender address: "),
		},
		{
			name: "Empty request receiver returns error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
			),
			expErr: fmt.Errorf("invalid receiver address: "),
		},
		{
			name: "Equals request receiver and request sender addresses return error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("the sender and receiver must be different"),
		},
		{
			name: "Valid request returns no error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.request.Validate())
		})
	}
}
