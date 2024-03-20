package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	chatAPI "github.com/Shemistan/chat_server/internal/api/chat_v1"
	"github.com/Shemistan/chat_server/internal/service/mocks"
)

func TestNew(t *testing.T) {
	service := new(mocks.Chat)

	api := chatAPI.New(service)
	assert.NotNil(t, api)
}
