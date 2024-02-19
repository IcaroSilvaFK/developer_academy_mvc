package utils_test

import (
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidateIdWhenIdIsValid(t *testing.T) {

	id := utils.NewId()

	actual := utils.IsValidId(id)

	assert.True(t, actual)
}

func TestValidateIdWhenIdIsNotValid(t *testing.T) {

	id := "invalid_id"

	actual := utils.IsValidId(id)

	assert.False(t, actual)
}
