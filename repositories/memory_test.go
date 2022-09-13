package repositories

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	chatID  int64 = 1
	userID1 int64 = 1
	userID2 int64 = 2
)

func TestMemoryRepository_GetActivityForUser(t *testing.T) {
	activity := UserActivity{
		ID:       userID1,
		Username: "username",
		LastSeen: time.Now(),
		Count:    1,
	}
	r := memoryRepository{}
	r[chatID] = chat{
		Activities: map[int64]UserActivity{
			userID1: activity,
		},
	}

	t.Run("user exists", func(t *testing.T) {
		v, found := r.GetActivityForUser(chatID, userID1)
		assert.True(t, found)
		assert.Equal(t, activity, v)
	})

	t.Run("user does not exist but chat exists", func(t *testing.T) {
		v, found := r.GetActivityForUser(chatID, 99)
		assert.False(t, found)
		assert.Equal(t, UserActivity{}, v)
	})

	t.Run("nor user or chat exists", func(t *testing.T) {
		v, found := r.GetActivityForUser(99, 99)
		assert.False(t, found)
		assert.Equal(t, UserActivity{}, v)
	})
}

func TestMemoryRepository_SetActivityForUser(t *testing.T) {
	t.Run("user exists", func(t *testing.T) {
		activity1 := UserActivity{
			ID:       userID1,
			Username: "username",
			LastSeen: time.Now(),
			Count:    1,
		}
		activity2 := UserActivity{
			ID:       userID1,
			Username: "username",
			LastSeen: time.Now(),
			Count:    2,
		}
		r := memoryRepository{}
		r[chatID] = chat{
			Activities: map[int64]UserActivity{
				userID1: activity1,
			},
		}

		r.SetActivityForUser(chatID, userID1, activity2)
		v, found := r.GetActivityForUser(chatID, userID1)
		assert.True(t, found)
		assert.Equal(t, activity2, v)
	})

	t.Run("user does not exist but chat exists", func(t *testing.T) {
		activity := UserActivity{
			ID:       userID1,
			Username: "username",
			LastSeen: time.Now(),
			Count:    1,
		}
		r := memoryRepository{}
		r[chatID] = chat{
			Activities: map[int64]UserActivity{},
		}

		r.SetActivityForUser(chatID, userID1, activity)
		v, found := r.GetActivityForUser(chatID, userID1)
		assert.True(t, found)
		assert.Equal(t, activity, v)
	})

	t.Run("nor user or chat exists", func(t *testing.T) {
		activity := UserActivity{
			ID:       userID1,
			Username: "username",
			LastSeen: time.Now(),
			Count:    1,
		}
		r := memoryRepository{}

		r.SetActivityForUser(chatID, userID1, activity)
		v, found := r.GetActivityForUser(chatID, userID1)
		assert.True(t, found)
		assert.Equal(t, activity, v)
	})
}

func TestMemoryRepository_GetActivities(t *testing.T) {

	t.Run("chat exists and has activities", func(t *testing.T) {
		activity1 := UserActivity{
			ID:       userID1,
			Username: "username1",
			LastSeen: time.Now(),
			Count:    1,
		}
		activity2 := UserActivity{
			ID:       userID2,
			Username: "username2",
			LastSeen: time.Now(),
			Count:    2,
		}
		r := memoryRepository{}
		r[chatID] = chat{
			Activities: map[int64]UserActivity{
				userID1: activity1,
				userID2: activity2,
			},
		}

		activities := r.GetActivities(chatID)
		assert.Equal(t, 2, len(activities))
		assert.Equal(t, []UserActivity{activity1, activity2}, activities)
	})

	t.Run("chat exists and has no activities", func(t *testing.T) {
		r := memoryRepository{}
		r[chatID] = chat{}

		activities := r.GetActivities(chatID)
		assert.Equal(t, 0, len(activities))
	})

	t.Run("chat does not exist", func(t *testing.T) {
		r := memoryRepository{}

		activities := r.GetActivities(chatID)
		assert.Equal(t, 0, len(activities))
	})
}

func TestMemoryRepository_GetWelcomeForChat(t *testing.T) {

	t.Run("chat exists and has welcome message", func(t *testing.T) {
		r := memoryRepository{}
		text := "Hello there"
		r[chatID] = chat{
			Welcome: text,
		}

		v, found := r.GetWelcomeForChat(chatID)
		assert.True(t, found)
		assert.Equal(t, text, v)
	})

	t.Run("chat exists and has no welcome message", func(t *testing.T) {
		r := memoryRepository{}
		r[chatID] = chat{}

		v, found := r.GetWelcomeForChat(chatID)
		assert.False(t, found)
		assert.Equal(t, "", v)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		r := memoryRepository{}

		v, found := r.GetWelcomeForChat(chatID)
		assert.False(t, found)
		assert.Equal(t, "", v)
	})
}

func TestMemoryRepository_SetWelcomeForChat(t *testing.T) {
	text := "Hello there"

	t.Run("chat exists and has welcome message", func(t *testing.T) {
		r := memoryRepository{}
		text2 := "Bye bye"
		r[chatID] = chat{
			Welcome: text,
		}

		r.SetWelcomeForChat(chatID, text2)
		v, found := r.GetWelcomeForChat(chatID)
		assert.True(t, found)
		assert.Equal(t, text2, v)
	})

	t.Run("chat exists and has no welcome message", func(t *testing.T) {
		r := memoryRepository{}
		r[chatID] = chat{}

		r.SetWelcomeForChat(chatID, text)
		v, found := r.GetWelcomeForChat(chatID)
		assert.True(t, found)
		assert.Equal(t, text, v)
	})

	t.Run("chat does not exist", func(t *testing.T) {
		r := memoryRepository{}

		r.SetWelcomeForChat(chatID, text)
		v, found := r.GetWelcomeForChat(chatID)
		assert.True(t, found)
		assert.Equal(t, text, v)
	})
}

func TestMemoryRepository_Set(t *testing.T) {
	r := memoryRepository{}
	r[chatID] = chat{
		Activities: map[int64]UserActivity{
			userID1: {
				ID:       userID1,
				Username: "username",
				LastSeen: time.Now(),
				Count:    1,
			},
		},
		Welcome: "Hello there",
	}

	v := `{}`
	err := r.Set(v)
	assert.Nil(t, err)
	assert.Equal(t, v, r.Dump())
}

func TestMemoryRepository_Dump(t *testing.T) {
	r := memoryRepository{}
	r[chatID] = chat{
		Activities: map[int64]UserActivity{
			userID1: {
				ID:       userID1,
				Username: "username",
				LastSeen: time.Date(2000, 1, 1, 17, 00, 00, 0, time.UTC),
				Count:    1,
			},
		},
		Welcome: "Hello there",
	}

	v := `{"1":{"welcome":"Hello there","activities":{"1":{"id":1,"username":"username","last_seen":"2000-01-01T17:00:00Z","count":1}}}}`
	assert.Equal(t, v, r.Dump())
}