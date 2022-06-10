package post_service

import (
	"github.com/abaron10/Posts-API-Golang/mocks"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post is empty")
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.PostRepository)
	var identifier int64 = 1
	post := model.Post{Id: identifier, Title: "A", Text: "B"}
	// Setup expectations
	mockRepo.On("FindAll").Return([]model.Post{post}, nil)
	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()
	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, identifier, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)

}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := model.Post{Id: 1, Text: "Test-api"}
	testService := NewPostService(nil)
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post title id empty")

}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.PostRepository)
	post := model.Post{Title: "A", Text: "B"}

	//Setup expectations
	mockRepo.On("Save", &post).Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.Id)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}
