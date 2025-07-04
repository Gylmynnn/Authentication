package dto

type UserResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omitempty"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Username string `json:"username" binding:"required,min=3,max=20" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Username string `json:"username" binding:"required,min=3,max=20" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password string `json:"password,omitempty" binding:"min=6,max=50"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20" gorm:"unique;not null"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}
