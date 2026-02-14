package controller

import (
	repository "mampuio-project/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepository repository.UserRepository
}

type WithdrawRequest struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

func NewUserController(UserRepository repository.UserRepository) *UserController {
	return &UserController{
		UserRepository: UserRepository,
	}
}

func (userController *UserController) Withdraw(c *gin.Context) {

	var req WithdrawRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	user := userController.UserRepository.GetUserByName(req.Name)

	err := userController.UserRepository.WithdrawBalance(user, req.Amount)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "withdraw success"})
}

func (userController *UserController) Balance(c *gin.Context) {

	name := c.Param("name")

	getBalance := userController.UserRepository.GetBalance(name)

	c.JSON(200, gin.H{
		"balance": getBalance,
	})
}
