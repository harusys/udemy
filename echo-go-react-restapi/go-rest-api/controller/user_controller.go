package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// インターフェース
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

// 実装クラス
type userController struct {
	uu usecase.IUserUsecase
}

// コンストラクター（依存性の注入）
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// メソッド
func (uc *userController) SignUp(c echo.Context) error {
	// 前処理
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// 実行
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	// 前処理
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// 実行
	token, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 後処理
	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour), // 有効期限は24時間
		Path:    "/",
		Domain:  os.Getenv("API_DOMAIN"),
		// Secure:  true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // クロスドメイン間でのCookie送信を許可
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(), // 有効期限なし
		Path:    "/",
		Domain:  os.Getenv("API_DOMAIN"),
		// Secure:  true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // クロスドメイン間でのCookie送信を許可
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
