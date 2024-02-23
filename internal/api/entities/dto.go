package dto

// HabitDTO represents the habit information needed by other packages
type HabitDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      uint   `json:"-"`
}

// UserDTO represents the user information needed by other packages
type UserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
