package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// CreateDataDeletionRequest 创建数据删除请求
type CreateDataDeletionRequest struct {
	ID      string              `json:"id" validate:"required"`
	Subject DataDeletionSubject `json:"subject" validate:"required"`
	DryRun  bool                `json:"dry_run"`
	Action  DataDeletionAction  `json:"action" validate:"required"`
}

// DataDeletionSubject 数据删除主体
type DataDeletionSubject struct {
	ID   string `json:"id" validate:"required"`
	Type string `json:"type" validate:"required,eq=Organization"`
}

// DataDeletionAction 数据删除操作
type DataDeletionAction struct {
	Name string `json:"name" validate:"required,oneof=PreCheck SoftDelete CancelSoftDelete HardDelete"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func main() {
	request := CreateDataDeletionRequest{
		ID: "123",
		Subject: DataDeletionSubject{
			ID:   "456",
			Type: "Organization",
		},
		Action: DataDeletionAction{
			Name: "PreCheck1", // Change this to test different values
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("Validation passed")
}
