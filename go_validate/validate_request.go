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
	Action  DataDeletionAction  `json:"action" validate:"required,actionname"`
}

// DataDeletionSubject 数据删除主体
type DataDeletionSubject struct {
	ID   string `json:"id" validate:"required"`
	Type string `json:"type" validate:"required,eq=Organization"`
}

// DataDeletionAction 数据删除操作
type DataDeletionAction struct {
	Name string `json:"name" validate:"required,actionname"`
}

// 自定义验证器
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("actionname", validateActionName)
}

// 自定义验证函数
func validateActionName(fl validator.FieldLevel) bool {
	actionName := fl.Field().String()
	allowedActions := []string{"PreCheck", "SoftDelete", "CancelSoftDelete", "HardDelete"}
	for _, action := range allowedActions {
		if actionName == action {
			return true
		}
	}
	return false
}

func main() {
	request := CreateDataDeletionRequest{
		ID: "123",
		Subject: DataDeletionSubject{
			ID:   "456",
			Type: "Organization",
		},
		Action: DataDeletionAction{
			Name: "PreCheck",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("Validation passed")
}

/*
虽然可以通过自定义方法实现复杂验证逻辑，但对于简单的枚举值验证，实际上可以通过组合使用 oneof 标签来实现，而不需要自定义验证函数。

参考: https://darjun.github.io/2020/04/04/godailylib/validator/
*/
