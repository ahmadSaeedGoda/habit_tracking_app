package controllers

import (
	"lasting-dynamics.com/habit_tracking_app/internal/api/services"
)

// BaseHandler contains common functionality for all handlers
type BaseHandler struct {
	baseService *services.BaseService
}
