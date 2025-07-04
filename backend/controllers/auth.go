package controllers

import (
	"github.com/Gylmynnn/go-gin-flutter-bloc/database"
	"github.com/Gylmynnn/go-gin-flutter-bloc/dto"
	"github.com/Gylmynnn/go-gin-flutter-bloc/helper"
	"github.com/Gylmynnn/go-gin-flutter-bloc/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

/*
Handler untuk authentication user register dan login
*/

// Register
func Register(c *gin.Context) {
	var req = dto.UserCreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FailedResponse{
			Success: false,
			Status:  http.StatusUnprocessableEntity,
			Message: helper.ValidatorErrorMsg(err),
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: helper.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if helper.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, helper.FailedResponse{
				Success: false,
				Status:  http.StatusConflict,
				Message: helper.ValidatorErrorMsg(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse{
				Success: false,
				Status:  http.StatusInternalServerError,
				Message: helper.ValidatorErrorMsg(err),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, helper.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "register berhasil",
		Data: dto.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: helper.DateFormatter(user.CreatedAt),
			UpdatedAt: helper.DateFormatter(user.UpdatedAt),
		},
	})
}

// Login
func Login(c *gin.Context) {

	var req = dto.UserLoginRequest{}
	var user = models.User{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FailedResponse{
			Success: false,
			Status:  http.StatusUnprocessableEntity,
			Message: helper.ValidatorErrorMsg(err),
		})
		return
	}

	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "username tidak ditemukan",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, helper.FailedResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Message: helper.ValidatorErrorMsg(err),
		})
		return
	}

	token := helper.GenerateToken(user.Username)

	c.JSON(http.StatusOK, helper.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "login berhasil",
		Data: dto.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: helper.DateFormatter(user.CreatedAt),
			UpdatedAt: helper.DateFormatter(user.UpdatedAt),
			Token:     &token,
		},
	})
}
