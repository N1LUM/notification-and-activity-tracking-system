package handler

import (
	"net/http"
	"notify-activity-tracking-system/internal/gateway/dto/user_context"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var input user_context.CreateUserInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Warnf("[UserHandler] Failed to bind JSON for CreateUser: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.services.User.CreateUser(input)
	if err != nil {
		logrus.Errorf("[UserHandler] Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.services.User.GetAllUsers()
	if err != nil {
		logrus.Errorf("[UserHandler] Failed to get users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.services.User.GetByID(parseUUID(id))
	if err != nil {
		logrus.Errorf("[UserHandler] Failed to get user by ID=%s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var input user_context.UpdateUserInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Warnf("[UserHandler] Failed to bind JSON for UpdateUser: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"error": "User not found"}})
		return
	}

	updatedUser, err := h.services.User.Update(parseUUID(id), input)
	if err != nil {
		logrus.Errorf("[UserHandler] Failed to update user by ID=%s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *Handler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.services.User.GetUserByUsername(username)
	if err != nil {
		logrus.Errorf("[Handler] Failed to get user by Username=%s: %v", username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	if err := h.services.User.DeleteUser(parseUUID(id), ctx); err != nil {
		logrus.Errorf("[UserHandler] Failed to delete user by ID=%s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func parseUUID(id string) uuid.UUID {
	uid, err := uuid.Parse(id)
	if err != nil {
		logrus.Warnf("[UserHandler] Invalid UUID: %s", id)
		return uuid.Nil
	}
	return uid
}
