package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// インターフェース
type IUserValidator interface {
	UserValidate(user model.User) error
}

// 構造体
type userValidator struct{}

// コンストラクター（依存性の注入）
func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// メソッド
func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}
