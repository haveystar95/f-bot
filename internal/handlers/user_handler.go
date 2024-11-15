package handlers

import (
	"f-bot/internal/models"
	"f-bot/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService services.UserService
}

type UserResponse struct {
	UserInfo     models.User `json:"user_info"`
	RandomString string      `json:"random_string"`
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// Get user ID from the URL
	id := c.Param("id")

	// Convert the ID to an integer
	// This will return an error if the ID is not a valid integer
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID sdsdsds"})
		return
	}

	// Call the GetUserByID method on the user service
	user, err := h.UserService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}

	// Generate a random string
	randomString := "random string"
	userResponse := UserResponse{
		UserInfo:     *user,
		RandomString: randomString,
	}
	// Return the user as JSON
	c.JSON(http.StatusOK, userResponse)
}
