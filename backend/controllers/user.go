package controllers

import (
	"net/http"

	"github.com/Gylmynnn/go-gin-flutter-bloc/database"
	"github.com/Gylmynnn/go-gin-flutter-bloc/dto"
	"github.com/Gylmynnn/go-gin-flutter-bloc/helper"
	"github.com/Gylmynnn/go-gin-flutter-bloc/models"
	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/gin-gonic/gin"
)

// mendapat semua data user
func FindAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	var response []dto.UserResponse
	for _, user := range users {
		res := dto.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: helper.DateFormatter(user.CreatedAt),
			UpdatedAt: helper.DateFormatter(user.UpdatedAt),
		}
		response = append(response, res)
	}

	c.JSON(http.StatusOK, helper.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "berhasil mendapatkan semua data user",
		Data:    response,
	})
}

// mendapat data dengan id user
func FindUserById(c *gin.Context) {

	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "user tidak ditemukan",
		})
		return
	}
	c.JSON(http.StatusOK, helper.SuccessResponse{
		Success: true,
		Message: "user ditemukan",
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

func CreateUser(c *gin.Context) {
	var req = dto.UserCreateRequest{}

	// Bind JSON request ke struct UserRequest
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
		c.JSON(http.StatusInternalServerError, helper.FailedResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, helper.SuccessResponse{
		Success: true,
		Message: "User created successfully",
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

// update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "user tidak ditemukan",
		})
		return
	}

	var req = dto.UserUpdateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FailedResponse{
			Success: false,
			Status:  http.StatusUnprocessableEntity,
			Message: helper.ValidatorErrorMsg(err),
		})
		return
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.Password = helper.HashPassword(req.Password)

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: helper.ValidatorErrorMsg(err),
		})
		return
	}

	// Kirimkan response sukses
	c.JSON(http.StatusOK, helper.SuccessResponse{
		Success: true,
		Message: "User updated successfully",
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

// delete data user
func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "user tidak ditemukan",
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: helper.ValidatorErrorMsg(err),
		})
		return
	}

	c.JSON(http.StatusOK, helper.SuccessResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}
