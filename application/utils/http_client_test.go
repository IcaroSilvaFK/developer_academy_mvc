package utils_test

import (
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldExecuteGetRequestAndPassFromStructResult(t *testing.T) {

	c := utils.NewHttpClient()

	var result struct {
		UserId    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	_, err := c.Get("https://jsonplaceholder.typicode.com/todos/1", &result, nil)

	assert.Nil(t, err)
}

func TestShouldReturnHttStatusOnExecuteRequest(t *testing.T) {

	c := utils.NewHttpClient()

	r, err := c.Get("https://jsonplaceholder.typicode.com/todos/1", nil, nil)

	assert.Equal(t, 200, r.StatusCode)
	assert.Nil(t, err)
}

func TestShouldExecutePostOnApi(t *testing.T) {

	c := utils.NewHttpClient()

	type example struct {
		UserId int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
	}

	payload := example{

		UserId: 1,
		ID:     1,
		Title:  "delectus aut autem",
	}

	r := example{}

	req, err := c.Post("https://jsonplaceholder.typicode.com/posts", &payload, &r)

	assert.Nil(t, err)
	assert.Equal(t, "delectus aut autem", r.Title)
	assert.Equal(t, 1, r.UserId)
	assert.Equal(t, 201, req.StatusCode)
}

func TestShouldAddParamsInUrl(t *testing.T) {

	c := utils.NewHttpClient()

	params := map[string]string{

		"param1": "value1",
		"param2": "value2",
	}

	url := c.CreateUrl("https://jsonplaceholder.typicode.com/posts", params)

	expected := "https://jsonplaceholder.typicode.com/posts?param1=value1&param2=value2"

	assert.Equal(t, expected, url)
}
