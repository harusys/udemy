package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// インターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// 実装クラス
type userRepository struct {
	db *gorm.DB
}

// コンストラクター（依存性の注入）
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// メソッド（ポインタレシーバ）
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
