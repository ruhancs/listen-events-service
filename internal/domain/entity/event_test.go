package entity_test

import (
	"testing"

	"github.com/ruhancs/listen-events/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	event, err := entity.NewEvent("payment", "order-created", "success", 10, 2, 2019)

	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "payment", event.Service)
	assert.Equal(t, 10, event.Day)
	assert.Equal(t, 2, event.Month)
	assert.Equal(t, 2019, event.Year)
	assert.Equal(t, "order-created", event.Type)
	assert.Equal(t, "success", event.Status)
}
