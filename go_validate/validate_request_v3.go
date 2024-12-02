package main

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	// "strings"
)

type CreateDataDeletionRequest struct {
	ID      string              `json:"id" validate:"required"`
	Subject DataDeletionSubject `json:"subject" validate:"required"`
	DryRun  bool                `json:"dry_run"`
	Action  DataDeletionAction  `json:"action" validate:"required"`
}

type DataDeletionSubject struct {
	ID   string `json:"id" validate:"required"`
	Type string `json:"type" validate:"required,eq=Organization"`
}

type DataDeletionAction struct {
	Name string `json:"name" validate:"required,oneof=PreCheck SoftDelete CancelSoftDelete HardDelete"`
}

type ErrDetail struct {
	Path string `json:"path"`
	Info string `json:"info"`
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
			Name: "InvalidAction", // Change this to test different values
		},
	}

	err := validate.Struct(request)
	if err != nil {
		var errDetails []ErrDetail
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			for _, validatorErr := range validationErrs {
				fieldName := validatorErr.Field()
				tag := validatorErr.Tag()
				var info string
				switch tag {
				case "required":
					info = fmt.Sprintf("%s is required", fieldName)
				case "oneof":
					info = fmt.Sprintf("%s must be one of [%s]", fieldName, validatorErr.Param())
				case "eq":
					info = fmt.Sprintf("%s must be equal to %s", fieldName, validatorErr.Param())
				default:
					info = fmt.Sprintf("%s failed on the %s tag", fieldName, tag)
				}
				fmt.Println(info)
				errDetails = append(errDetails, ErrDetail{
					Path: validatorErr.Namespace(),
					// Info: info,
					Info: validatorErr.Error(),
				})
			}
		} else {
			fmt.Println("Unexpected error:", err)
			return
		}

		for _, detail := range errDetails {
			fmt.Printf("Validation error: %s - %s\n", detail.Path, detail.Info)
		}
		return
	}

	fmt.Println("Validation passed")
}

/*
为了优化校验错误信息的展示，使其更加明确和可读，我们可以自定义错误信息的格式。通过解析 validator.ValidationErrors，可以提取出具体的字段和错误类型，并将其转换为更友好的错误信息。
*/
