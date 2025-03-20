package repositories_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	mocks_test "github.com/IcaroSilvaFK/developer_academy_mvc/mocks"
	"github.com/stretchr/testify/assert"
)

func setupUserRepo() repositories.UserRepositoryInterface {
	mockDb := mocks_test.NewGormDbMock()
	repo := repositories.NewUserRepository(mockDb)

	return repo
}

var ctx = context.Background()

func TestShouldCreateUser(t *testing.T) {
	repo := setupUserRepo()

	u := models.NewUserModel("test@test.com", "test", "https://", "http://", "test create new user", nil)

	err := repo.Create(ctx, u)

	assert.Nil(t, err)
}

func TestShouldFindByEmailUser(t *testing.T) {

	repo := setupUserRepo()
	uEmail := "test2@test.com"
	u := models.NewUserModel(uEmail, "test", "https://", "http://", "test create new user", nil)

	err := repo.Create(ctx, u)

	assert.Nil(t, err)

	user, err := repo.FindByEmail(ctx, uEmail)

	assert.Nil(t, err)

	assert.Equal(t, u.ID, user.ID)
}

func TestShouldFindByIdUser(t *testing.T) {

	repo := setupUserRepo()
	u := models.NewUserModel("test@test.com", "test", "https://", "http://", "test create new user", nil)

	err := repo.Create(ctx, u)

	assert.Nil(t, err)

	user, err := repo.FindById(ctx, u.ID)

	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
}

func TestShouldFindAllUsers(t *testing.T) {

	repo := setupUserRepo()

	expectedQuantityUsers := 10

	for i := 0; i < expectedQuantityUsers; i++ {
		u := models.NewUserModel(fmt.Sprintf("test%d@test.com", i), "test", "https://", "http://", "test create new user", nil)
		repo.Create(ctx, u)
	}

	users, err := repo.FindAll(ctx)

	assert.Nil(t, err)
	assert.Equal(t, expectedQuantityUsers, len(users))
}

func TestShouldFindTenAndCountUsers(t *testing.T) {

	repo := setupUserRepo()

	expectedQuantityUsers := 20

	for i := 0; i < expectedQuantityUsers; i++ {
		u := models.NewUserModel(fmt.Sprintf("test%d@test.com", i), "test", "https://", "http://", "test create new user", nil)
		repo.Create(ctx, u)
	}

	users, quantity, err := repo.FindFirstTenAndCount(ctx)

	assert.Nil(t, err)
	assert.Equal(t, 10, len(users))
	assert.Equal(t, 10, quantity)
}

func TestShouldDeleteUser(t *testing.T) {

	repo := setupUserRepo()

	u := models.NewUserModel("test@test.com", "test", "https://", "http://", "test create new user", nil)

	repo.Create(ctx, u)

	err := repo.Delete(ctx, u.ID)

	assert.Nil(t, err)
}

func TestShouldCountUsers(t *testing.T) {
	repo := setupUserRepo()

	expectedQuantityUsers := 20

	for i := 0; i < expectedQuantityUsers; i++ {
		u := models.NewUserModel(fmt.Sprintf("test%d@test.com", i), "test", "https://", "http://", "test create new user", nil)
		repo.Create(ctx, u)
	}

	c, err := repo.Count(ctx)

	assert.Nil(t, err)
	assert.Equal(t, int64(expectedQuantityUsers), c)
}
