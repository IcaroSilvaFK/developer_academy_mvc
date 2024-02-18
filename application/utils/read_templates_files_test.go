package utils_test

import (
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldReadTemplatesFiles(t *testing.T) {

	paths, err := utils.ReadTemplatesFiles("./public/pages")

	assert.Nil(t, err)
	assert.NotNil(t, paths)
}
