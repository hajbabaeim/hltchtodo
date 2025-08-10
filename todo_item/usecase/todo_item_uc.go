package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	"github.com/hajbabaeim/hltchtodo/helpers"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain"
	"github.com/hajbabaeim/hltchtodo/todo_item/domain/requests"
	"strings"
	"time"
)

func (uc *usecase) CreateItem(ctx context.Context, req *requests.CreateItemRequest) (*domain.TodoItem, error) {
	item := new(domain.TodoItem)
	layout := "2006-01-02 15:04:05" // Input request time-format "YYYY-MM-DD HH:MM:SS"
	parsedTime, err := time.Parse(layout, req.DueDate)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}
	item.DueDate = parsedTime
	item.Description = req.Description
	uid := uuid.New()
	item.UUID = uid
	if err = uc.repo.CreateItem(item); err != nil {
		return nil, err
	}
	if err = uc.sendMessageToQueue(ctx, item); err != nil {
		uc.logger.Errorf("send message to queue failed: %v", err)
		return nil, err
	}

	return uc.repo.GetItemByUUID(item.UUID)
}

func (uc *usecase) sendMessageToQueue(ctx context.Context, item *domain.TodoItem) error {
	if uc.queueURL == "" {
		return errors.New("queue URL not configured")
	}
	// build up msg for sqs
	messageData := map[string]interface{}{
		"event_type": "todo_item_created",
		"todo_item": map[string]interface{}{
			"id":          item.ID,
			"uuid":        item.UUID.String(),
			"description": item.Description,
			"due_date":    item.DueDate,
			"created_at":  item.CreatedAt,
		},
		"timestamp": item.CreatedAt,
	}

	messageBody, err := json.Marshal(messageData)
	if err != nil {
		return err
	}

	_, err = uc.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(uc.queueURL),
		MessageBody: aws.String(string(messageBody)),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"EventType": {
				DataType:    aws.String("String"),
				StringValue: aws.String("todo_item_created"),
			},
			"TodoItemUUID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(item.UUID.String()),
			},
		},
	})

	return err
}

func (uc *usecase) UpdateItem(ctx context.Context, req *requests.UpdateItemRequest) (*domain.TodoItem, error) {
	if req.Id == nil && *req.Id == "" {
		return nil, errors.New("id is required")
	}
	uid, err := uuid.Parse(*req.Id)
	if err != nil {
		return nil, err
	}
	oldItem, err := uc.repo.GetItemByUUID(uid)
	if err != nil {
		return nil, err
	}
	if req.Description != nil && oldItem.Description != *req.Description {
		oldItem.Description = *req.Description
	}
	if req.DueDate != nil && len(strings.Split(*req.DueDate, " ")) == 2 {
		parsedTime, err := helpers.ConvertStringTime(*req.DueDate)
		if err != nil {
			return nil, err
		}
		oldItem.DueDate = parsedTime
	}
	if err = uc.repo.UpdateItem(oldItem); err != nil {
		return nil, err
	}
	return uc.repo.GetItemByUUID(oldItem.UUID)
}

func (uc *usecase) DeleteItem(ctx context.Context, req *requests.DeleteItemRequest) (bool, error) {
	if req.Id == "" {
		return false, errors.New("id is required")
	}
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		return false, err
	}
	if err = uc.repo.DeleteItem(uid); err != nil {
		return false, err
	}
	return true, nil
}

func (uc *usecase) ListItems(ctx context.Context) ([]*domain.TodoItem, error) {
	items, err := uc.repo.ListItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *usecase) GetItem(ctx context.Context, req *requests.GetItemRequest) (*domain.TodoItem, error) {
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetItemByUUID(uid)
}
