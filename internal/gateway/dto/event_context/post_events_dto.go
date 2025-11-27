package event_context

import (
	"notify-activity-tracking-system/internal/gateway/models"
)

type PostEventsDTO struct {
	Events []models.Event `json:"events" binding:"required,min=1""`
}
