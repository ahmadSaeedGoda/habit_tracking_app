package services

import (
	"lasting-dynamics.com/habit_tracking_app/internal/api/repositories"
)

// BaseService contains common functionality for all handlers
type BaseService struct {
	baseRepo *repositories.BaseRepository
}
