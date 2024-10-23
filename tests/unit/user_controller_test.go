package unit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/RianAsmara/personal-finance-advisor-api/model"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetUsers(t *testing.T) {
	// authentication
	tokenResponse := AuthenticationTest()

	request := httptest.NewRequest("GET", "/v1/api/users", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+tokenResponse["token"].(string))

	response, _ := appTest.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := io.ReadAll(response.Body)

	var webResponse model.GeneralResponse
	err := json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, webResponse.Success)
	assert.Equal(t, "Success", webResponse.Message)

	// Assuming the Data field is a JSON array of users
	// var users []model.User
	// dataStr, ok := webResponse.Data.(string)
	// assert.Equal(t, true, ok)
	// err = json.Unmarshal([]byte(dataStr), &users)
	// assert.Equal(t, nil, err)
	// assert.NotEqual(t, 0, len(users))

	// // Check the first user in the list
	// if len(users) > 0 {
	// 	assert.Equal(t, "admin@gmail.com", users[0].Email)
	// }
}

func TestCreateUserWithDuplicateEmail(t *testing.T) {
	// Authentication
	tokenResponse := AuthenticationTest()

	// Create a user with a specific email
	firstUser := model.UserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Create the first user
	createUserRequest(t, firstUser, tokenResponse["token"].(string))

	// Attempt to create a second user with the same email
	duplicateUser := model.UserRequest{
		Email:    "test@example.com",
		Password: "anotherpassword",
	}

	request := httptest.NewRequest("POST", "/v1/api/users", createJSONBody(duplicateUser))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+tokenResponse["token"].(string))

	response, _ := appTest.Test(request)

	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := io.ReadAll(response.Body)

	var webResponse model.GeneralResponse
	err := json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, webResponse.Success)
	assert.Equal(t, "Email already exists", webResponse.Message)
}

func createUserRequest(t *testing.T, user model.UserRequest, token string) {
	request := httptest.NewRequest("POST", "/v1/api/users", createJSONBody(user))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, _ := appTest.Test(request)
	fmt.Println("===================================	")
	fmt.Println(response.StatusCode)
	fmt.Println("===================================	")
	assert.Equal(t, 201, response.StatusCode)
}

func createJSONBody(v interface{}) io.Reader {
	jsonBytes, _ := json.Marshal(v)
	return io.NopCloser(io.Reader(bytes.NewBuffer(jsonBytes)))
}
