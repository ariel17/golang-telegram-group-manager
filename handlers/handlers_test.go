package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

func TestRemoveCommandFromText(t *testing.T) {

	t.Run("with params", func(t *testing.T) {
		v := removeCommandFromText("/debug 1 2 3", config.Debug)
		assert.Equal(t, "1 2 3", v)
	})

	t.Run("without params", func(t *testing.T) {
		v := removeCommandFromText("/debug", config.Debug)
		assert.Equal(t, "", v)
	})
}