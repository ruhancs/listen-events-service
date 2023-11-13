package entity_test

import (
	"testing"

	"github.com/ruhancs/listen-events/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewLogError(t *testing.T) {
	logErr, err := entity.NewLogError("error to finish the sale", "payment-service", "so:wind,ip:123..", 403,1,3,2018)

	assert.Nil(t, err)
	assert.NotNil(t, logErr)
	assert.Equal(t, "error to finish the sale", logErr.Message)
	assert.Equal(t, "payment-service", logErr.Service)
	assert.Equal(t, 1, logErr.Day)
	assert.Equal(t, 3, logErr.Month)
	assert.Equal(t, 2018, logErr.Year)
	assert.Equal(t, "so:wind,ip:123..", logErr.UserInfo)
	assert.Equal(t, 403, logErr.StatusCode)
}
