package utils_test

import (
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewId(t *testing.T) {

	id := utils.NewId()

	assert.NotNil(t, id)
	assert.Equal(t, 36, len(id))
}
