package types_test

import (
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types"

	"testing"
	"time"
)

func TestPollAnswer_Validate(t *testing.T) {
	answer := types.NewPollAnswer("0", "")
	require.Equal(t, "answer text must be specified and cannot be empty", answer.Validate().Error())
}

func TestPollAnswer_Equals(t *testing.T) {
	tests := []struct {
		name        string
		answer      types.PollAnswer
		otherAnswer types.PollAnswer
		expEquals   bool
	}{
		{
			name:        "Different answers ID",
			answer:      types.NewPollAnswer("1", "Yes"),
			otherAnswer: types.NewPollAnswer("2", "Yes"),
			expEquals:   false,
		},
		{
			name:        "Different answers Text",
			answer:      types.NewPollAnswer("1", "Yes"),
			otherAnswer: types.NewPollAnswer("1", "No"),
			expEquals:   false,
		},
		{
			name:        "Equals answers",
			answer:      types.NewPollAnswer("1", "yes"),
			otherAnswer: types.NewPollAnswer("1", "yes"),
			expEquals:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.answer.Equal(test.otherAnswer))
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestPollData_Validate(t *testing.T) {
	var timeZone, _ = time.LoadLocation("UTC")
	var pollEndDate = time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)

	tests := []struct {
		pollData types.PollData
		expError string
	}{
		{
			pollData: types.NewPollData("", pollEndDate, types.PollAnswers{}, true, true),
			expError: "missing poll question",
		},
		{
			pollData: types.NewPollData("title", time.Time{}, types.PollAnswers{}, true, true),
			expError: "invalid poll's end date",
		},
		{
			pollData: types.NewPollData("title", pollEndDate, types.PollAnswers{}, true, true),
			expError: "poll answers must be at least two",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expError, test.pollData.Validate().Error())
	}
}

func TestUserAnswer_Validate(t *testing.T) {
	tests := []struct {
		name            string
		userPollAnswers types.UserAnswer
		expErr          string
	}{
		{
			name:            "Empty user returns error",
			userPollAnswers: types.NewUserAnswer([]string{"1", "2"}, ""),
			expErr:          "user cannot be empty",
		},
		{
			name:            "Empty answers returns error",
			userPollAnswers: types.NewUserAnswer(nil, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expErr:          "answers cannot be empty",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if err := test.userPollAnswers.Validate(); err != nil {
				require.Equal(t, test.expErr, err.Error())
			}
		})
	}
}

func TestUserAnswer_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.UserAnswer
		second    types.UserAnswer
		expEquals bool
	}{
		{
			name:      "Different users returns false",
			first:     types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			second:    types.NewUserAnswer([]string{"1", "2"}, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			expEquals: false,
		},
		{
			name:      "Different answers lengths returns false",
			first:     types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			second:    types.NewUserAnswer([]string{"1"}, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			expEquals: false,
		},
		{
			name:      "Different answers return false",
			first:     types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			second:    types.NewUserAnswer([]string{"1"}, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			expEquals: false,
		},
		{
			name:      "Equals userPollAnswers returns true",
			first:     types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			second:    types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equal(test.second))
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestUserAnswers_AppendIfMissingOrIfUsersEquals(t *testing.T) {
	tests := []struct {
		name             string
		answers          types.UserAnswers
		answer           types.UserAnswer
		expectedSlice    types.UserAnswers
		expectedAppended bool
	}{
		{
			name: "Missing user answers details appended correctly",
			answers: types.UserAnswers{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			answer: types.NewUserAnswer([]string{"1", "2"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			expectedSlice: types.UserAnswers{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			expectedAppended: true,
		},
		{
			name: "Same user with different answers replace previous ones",
			answers: types.UserAnswers{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			answer: types.NewUserAnswer([]string{"3"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expectedSlice: types.UserAnswers{
				types.NewUserAnswer([]string{"3"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			expectedAppended: true,
		},
		{
			name: "Equals user answers details returns the same users answers details",
			answers: types.UserAnswers{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			answer: types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expectedSlice: types.UserAnswers{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			expectedAppended: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.answers.AppendIfMissingOrIfUsersEquals(test.answer)
			require.Equal(t, test.expectedSlice, actual)
			require.Equal(t, test.expectedAppended, appended)
		})
	}
}
