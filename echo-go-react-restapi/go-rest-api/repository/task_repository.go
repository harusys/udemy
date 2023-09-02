package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// インターフェース
type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// 実装クラス
type taskRepository struct {
	db *gorm.DB
}

// コンストラクター（依存性の注入）
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// メソッド（ポインタレシーバ）
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	// .Model: モデルのポインタ渡し
	// .Clauses: 更新後モデルをポインタのアドレスに書き込み
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { // 更新対象なし(0件)はエラー発生なし
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { // 削除対象なし(0件)はエラー発生なし
		return fmt.Errorf("object does not exist")
	}
	return nil
}
