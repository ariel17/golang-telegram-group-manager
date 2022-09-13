package repositories

import "encoding/json"

type memoryRepository map[int64]map[string]interface{}

func (m memoryRepository) GetActivityForUser(chatID, userID int64) (UserActivity, bool) {
	chat, exists := m[chatID]
	if !exists {
		return UserActivity{}, false
	}
	activities, exists := chat["activities"]
	if !exists {
		return UserActivity{}, false
	}
	activity, exists := activities.(map[int64]UserActivity)[userID]
	if !exists {
		return UserActivity{}, false
	}
	return activity, true
}

func (m memoryRepository) SetActivityForUser(chatID, userID int64, activity UserActivity) {
	chat, exists := m[chatID]
	if !exists {
		chat = map[string]interface{}{}
	}
	activities, exists := chat["activities"]
	if !exists {
		activities = map[int64]UserActivity{}
	}
	activities.(map[int64]UserActivity)[userID] = activity
	chat["activities"] = activities
	m[chatID] = chat
}

func (m memoryRepository) GetActivities(chatID int64) []UserActivity {
	chat, exists := m[chatID]
	if !exists {
		return []UserActivity{}
	}
	activities, exists := chat["activities"]
	if !exists {
		return []UserActivity{}
	}
	l := []UserActivity{}
	for _, a := range activities.(map[int64]UserActivity) {
		l = append(l, a)
	}
	return l
}

func (m memoryRepository) GetWelcomeForChat(chatID int64) (string, bool) {
	chat, exists := m[chatID]
	if !exists {
		return "", false
	}
	text, exists := chat["welcome"]
	if !exists {
		return "", false
	}
	return text.(string), true
}

func (m memoryRepository) SetWelcomeForChat(chatID int64, text string) {
	chat, exists := m[chatID]
	if !exists {
		chat = map[string]interface{}{}
	}
	chat["welcome"] = text
	m[chatID] = chat
}

func (m memoryRepository) Set(value string) error {
	var temp memoryRepository
	err := json.Unmarshal([]byte(value), &temp)
	if err != nil {
		return err
	}
	for k, v := range temp {
		m[k] = v
	}
	for k := range m {
		_, exists := temp[k]
		if !exists {
			delete(m, k)
		}
	}
	return nil
}

func (m memoryRepository) Dump() string {
	b, _ := json.Marshal(m)
	return string(b)
}