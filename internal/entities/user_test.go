package entities_test

import (
	"testing"

	"github.com/resyahrial/go-template/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	input := entities.CreateUserRequest{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name           string
		expectedOutput *entities.User
		expectedError  []error
		opts           []entities.UserOption
	}{
		{
			name: "should create basic user",
			expectedOutput: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
	}

	for _, tc := range testCases {
		user, errs := entities.NewUser(input, tc.opts...)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedError, errs)
			assert.EqualValues(t, tc.expectedOutput, user)
		})
	}
}
