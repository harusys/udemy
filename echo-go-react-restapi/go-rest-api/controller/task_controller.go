package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// インターフェース
type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

// 構造体
type taskController struct {
	tu usecase.ITaskUsecase
}

// コンストラクター（依存性の注入）
func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

// メソッド
func (tc *taskController) GetAllTasks(c echo.Context) error {
	// 前処理
	user := c.Get("user").(*jwt.Token) // コンテキストから取得
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64) // any -> float64に型アサーション
	// 実行
	tasksRes, err := tc.tu.GetAllTasks(uint(userId)) // uint型へキャスト
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	// 前処理
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	id := c.Param("taskId")       // パスパラメータから取得 `tasks/:taskId`
	taskId, _ := strconv.Atoi(id) // string型からint型へキャスト
	// 実行
	taskRes, err := tc.tu.GetTaskById(uint(userId), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	// 前処理
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	task := model.Task{}
	if err := c.Bind(&task); err != nil { // リクエストボディから取得
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId) // JWTから取得したユーザIDにセット
	// 実行
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	// 前処理
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// 実行
	taskRes, err := tc.tu.UpdateTask(task, uint(userId), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	// 前処理
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	// 実行
	if err := tc.tu.DeleteTask(uint(userId), uint(taskId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
