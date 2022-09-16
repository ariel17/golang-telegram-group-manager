package services

import "log"

// Debug sets/exposes internal data structure.
func Debug(text string) (string, error) {
	if text == "" {
		text = repository.Dump()
		log.Printf(text)
		return text, nil
	}
	return "", repository.Set(text)
}