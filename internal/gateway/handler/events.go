package handler

import (
	"net/http"
	"notify-activity-tracking-system/internal/gateway/dto/event_context"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostEvents(c *gin.Context) {
	var input event_context.PostEventsDTO
	//ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//user, accessToken, refreshToken, err := h.services.Auth.Login(input.Username, input.Password, ctx)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"id":           user.ID,
	//	"username":     user.Username,
	//	"accessToken":  accessToken,
	//	"refreshToken": refreshToken,
	//})
}
