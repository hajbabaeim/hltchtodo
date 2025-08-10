package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ----------------------------------------Mock Repository----------------------------------//
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateItem(item *domain.TodoItem) error {
	args := m.Called(item)
	return args.Error(0)
}
func (m *MockRepository) GetItemByID(id uint64) (*domain.TodoItem, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TodoItem), args.Error(1)
}
func (m *MockRepository) GetItemByUUID(uid uuid.UUID) (*domain.TodoItem, error) {
	args := m.Called(uid)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TodoItem), args.Error(1)
}
func (m *MockRepository) UpdateItem(item *domain.TodoItem) error {
	args := m.Called(item)
	return args.Error(0)
}
func (m *MockRepository) DeleteItem(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockRepository) ListItems() ([]*domain.TodoItem, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.TodoItem), args.Error(1)
}

// ----------------------------------------Mock SQS----------------------------------//
type MockSQSClient struct {
	mock.Mock
}

func (m *MockSQSClient) SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error) {
	args := m.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sqs.SendMessageOutput), args.Error(1)
}

// ----------------------------------------TEST Cases----------------------------------//
func TestUsecase_CreateItem_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	req := &requests.CreateItemRequest{
		Description: "Test todo item",
		DueDate:     "2024-12-31 23:59:59",
	}

	expectedItem := &domain.TodoItem{
		UUID:        uuid.New(),
		Description: req.Description,
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	mockRepo.On("CreateItem", mock.AnythingOfType("*domain.TodoItem")).Return(nil)
	mockRepo.On("GetItemByUUID", mock.AnythingOfType("uuid.UUID")).Return(expectedItem, nil)
	mockSQS.On("SendMessage", mock.Anything, mock.AnythingOfType("*sqs.SendMessageInput"), mock.Anything).Return(&sqs.SendMessageOutput{}, nil)

	result, err := uc.CreateItem(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedItem.Description, result.Description)

	mockRepo.AssertExpectations(t)
	mockSQS.AssertExpectations(t)
}

func TestUsecase_CreateItem_RepositoryError(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	req := &requests.CreateItemRequest{
		Description: "Test todo item",
		DueDate:     "2024-12-31 23:59:59",
	}

	expectedError := errors.New("database error")
	mockRepo.On("CreateItem", mock.AnythingOfType("*domain.TodoItem")).Return(expectedError)

	result, err := uc.CreateItem(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
	mockSQS.AssertNotCalled(t, "SendMessage")
}

func TestUsecase_CreateItem_SQSError_ShouldFailRequest(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	req := &requests.CreateItemRequest{
		Description: "Test todo item",
		DueDate:     "2024-12-31 23:59:59",
	}

	mockRepo.On("CreateItem", mock.AnythingOfType("*domain.TodoItem")).Return(nil)
	mockSQS.On("SendMessage", mock.Anything, mock.AnythingOfType("*sqs.SendMessageInput"), mock.Anything).Return(nil, errors.New("SQS error"))

	result, err := uc.CreateItem(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "SQS error")

	mockRepo.AssertExpectations(t)
	mockSQS.AssertExpectations(t)
}

func TestUsecase_GetItem_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	testUUID := uuid.New()
	req := &requests.GetItemRequest{
		Id: testUUID.String(),
	}

	expectedItem := &domain.TodoItem{
		UUID:        testUUID,
		Description: "Test item",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	mockRepo.On("GetItemByUUID", testUUID).Return(expectedItem, nil)

	result, err := uc.GetItem(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedItem.UUID, result.UUID)
	assert.Equal(t, expectedItem.Description, result.Description)

	mockRepo.AssertExpectations(t)
}

func TestUsecase_GetItem_InvalidUUID(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	req := &requests.GetItemRequest{
		Id: "invalid-uuid",
	}

	result, err := uc.GetItem(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)

	mockRepo.AssertNotCalled(t, "GetItemByUUID")
}

func TestUsecase_UpdateItem_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	testUUID := uuid.New()
	testID := testUUID.String()
	newDescription := "Updated description"
	newDueDate := "2024-12-30 20:00:00"

	req := &requests.UpdateItemRequest{
		Id:          &testID,
		Description: &newDescription,
		DueDate:     &newDueDate,
	}

	existingItem := &domain.TodoItem{
		UUID:        testUUID,
		Description: "Old description",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	updatedItem := &domain.TodoItem{
		UUID:        existingItem.UUID,
		Description: newDescription,
		DueDate:     time.Now().Add(48 * time.Hour),
	}

	mockRepo.On("GetItemByUUID", testUUID).Return(existingItem, nil)
	mockRepo.On("UpdateItem", mock.AnythingOfType("*domain.TodoItem")).Return(nil)
	mockRepo.On("GetItemByUUID", existingItem.UUID).Return(updatedItem, nil)

	result, err := uc.UpdateItem(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newDescription, result.Description)

	mockRepo.AssertExpectations(t)
}

func TestUsecase_UpdateItem_MissingID(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	req := &requests.UpdateItemRequest{
		Id: nil,
	}

	result, err := uc.UpdateItem(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")

	mockRepo.AssertNotCalled(t, "GetItemByUUID")
}

func TestUsecase_DeleteItem_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	testUUID := uuid.New()
	req := &requests.DeleteItemRequest{
		Id: testUUID.String(),
	}

	mockRepo.On("DeleteItem", testUUID).Return(nil)

	result, err := uc.DeleteItem(context.Background(), req)

	assert.NoError(t, err)
	assert.True(t, result)

	mockRepo.AssertExpectations(t)
}

func TestUsecase_DeleteItem_EmptyID(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	req := &requests.DeleteItemRequest{
		Id: "",
	}

	result, err := uc.DeleteItem(context.Background(), req)

	assert.Error(t, err)
	assert.False(t, result)
	assert.Contains(t, err.Error(), "id is required")

	mockRepo.AssertNotCalled(t, "DeleteItem")
}

func TestUsecase_ListItems_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockSQS := new(MockSQSClient)
	mockLogger := logrus.New()
	queueURL := "http://localhost:4566/000000000000/test-queue"

	uc := &usecase{
		repo:     mockRepo,
		sqs:      mockSQS,
		queueURL: queueURL,
		logger:   mockLogger,
	}

	expectedItems := []*domain.TodoItem{
		{
			UUID:        uuid.New(),
			Description: "Item 1",
			DueDate:     time.Now().Add(24 * time.Hour),
		},
		{
			UUID:        uuid.New(),
			Description: "Item 2",
			DueDate:     time.Now().Add(48 * time.Hour),
		},
	}

	mockRepo.On("ListItems").Return(expectedItems, nil)

	result, err := uc.ListItems(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedItems[0].Description, result[0].Description)
	assert.Equal(t, expectedItems[1].Description, result[1].Description)

	mockRepo.AssertExpectations(t)
}
