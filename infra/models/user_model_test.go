package models_test

import (
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewUserWhenPassValidData(t *testing.T) {

	u := models.NewUserModel("test@test.com", "test", "https://", "http://", "test create new user")

	assert.NotNil(t, u.ID)
	assert.NotNil(t, u.Email)
	assert.NotNil(t, u.Url)
	assert.NotNil(t, u.Bio)
	assert.NotNil(t, u.AvatarUrl)
	assert.NotNil(t, u.Name)
}
