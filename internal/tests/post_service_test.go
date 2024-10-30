package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vladovsiychuk/microservice-demo-go/internal/post"
	"github.com/vladovsiychuk/microservice-demo-go/mocks"
)

var anyValidPostRequest = post.PostRequest{Content: "foo", IsPrivate: false}
var errFoo = errors.New("Error")

func TestCreatePost(t *testing.T) {
	repository := mocks.NewPostRepositoryI(t)
	eventbus := mocks.NewEventBusI(t)
	service := post.NewService(repository, eventbus)

	tests := []struct {
		name          string
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Success",
			setupMocks: func() {
				repository.On("Create", mock.Anything).Return(nil).Once()
				eventbus.On("Publish", mock.Anything)
			},
			expectedError: nil,
		},
		{
			name: "With DB error",
			setupMocks: func() {
				repository.On("Create", mock.Anything).Return(errFoo).Once()
			},
			expectedError: errFoo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := service.CreatePost(anyValidPostRequest)

			assert.Equal(t, tt.expectedError, err)
			if tt.expectedError != nil {
				assert.Nil(t, response)
			} else {
				assert.NotNil(t, response)
			}

			repository.AssertExpectations(t)
			eventbus.AssertExpectations(t)
		})
	}
}
