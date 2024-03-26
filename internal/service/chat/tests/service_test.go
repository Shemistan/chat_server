package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/chat_server/internal/service/chat"
	"github.com/Shemistan/chat_server/internal/storage/mocks"
)

func TestNewService(t *testing.T) {
	storage := new(mocks.Chat)
	service := chat.NewService(storage)

	assert.NotNil(t, service)
}
