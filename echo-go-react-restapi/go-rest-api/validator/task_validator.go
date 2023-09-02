package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// インターフェース
type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

// 構造体
type taskValidator struct{}

// コンストラクター（依存性の注入）
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// メソッド
func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}
