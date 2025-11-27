package models

import "time"

type Event struct {
	ID        string                 `json:"id" binding:"required"`
	Type      string                 `json:"type" binding:"required"`
	UserID    string                 `json:"user_id" binding:"required"`
	Timestamp time.Time              `json:"timestamp" binding:"required"`
	Payload   map[string]interface{} `json:"payload" binding:"omitempty"`
}
