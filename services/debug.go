package services

// Debug sets/exposes internal data structure.
func Debug(text string) (string, error) {
	if text == "" {
		return repository.Dump(), nil
	}
	return "", repository.Set(text)
}