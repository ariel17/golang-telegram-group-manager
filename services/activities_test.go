package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetInactives(t *testing.T) {
	activities = map[int64]map[int64]UserActivity{
		1: {
			1: {
				ID:       1,
				Username: "user1",
				LastSeen: time.Now().AddDate(0, 0, -45),
			},
			2: {
				ID:       2,
				Username: "user2",
				LastSeen: time.Now().AddDate(0, 0, -25),
			},
		},
	}
	inactives := GetInactives(30, 1)
	assert.Equal(t, 1, len(inactives))
}