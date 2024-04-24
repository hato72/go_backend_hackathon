package controller

import (
	"backend/src/model"
	"backend/src/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ICuisineController interface {
	GetAllCuisines(c echo.Context) error
	GetCuisineById(c echo.Context) error
	CreateCuisine(c echo.Context) error
	UpdateCuisine(c echo.Context) error
	DeleteCuisine(c echo.Context) error
}

type cuisineController struct {
	cu usecase.ICuisineUsecase
}

func NewCuisineController(tu usecase.ICuisineUsecase) ICuisineController {
	return &cuisineController{tu}
}

func (cc *cuisineController) GetAllCuisines(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)    //コンテキストからjwtをデコードした値を読み込む
	claims := user.Claims.(jwt.MapClaims) //その中のデコードされたclaimsを取得
	userId := claims["user_id"]           //claimsの中のuserIdを取得

	taskRes, err := cc.cu.GetAllCuisines(uint(userId.(float64))) //一度floatにしてからuintに型変換
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (cc *cuisineController) GetCuisineById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")       //リクエストパラメーターからtaskIdを取得
	taskId, _ := strconv.Atoi(id) //stringからintに
	taskRes, err := cc.cu.GetCuisineById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (cc *cuisineController) CreateCuisine(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := model.Cuisine{}
	if err := c.Bind(&task); err != nil { //リクエストボディに含まれる内容をtask構造体に代入
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId.(float64))
	taskRes, err := cc.cu.CreateCuisine(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (cc *cuisineController) UpdateCuisine(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Cuisine{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := cc.cu.UpdateCuisine(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (cc *cuisineController) DeleteCuisine(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Cuisine{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := cc.cu.DeleteCuisine(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
