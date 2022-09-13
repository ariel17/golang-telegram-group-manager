package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/golang-telegram-group-manager/repositories"
)

func TestGetInactives(t *testing.T) {
	var chatID int64 = 1
	repository.SetActivityForUser(chatID, 1, repositories.UserActivity{
		ID:       1,
		Username: "user1",
		LastSeen: time.Now().AddDate(0, 0, -45),
	})
	repository.SetActivityForUser(chatID, 2, repositories.UserActivity{
		ID:       2,
		Username: "user2",
		LastSeen: time.Now().AddDate(0, 0, -25),
	})
	inactives := GetInactives(30, chatID)
	assert.Equal(t, 1, len(inactives))
}