package auth_context

type LoginInputDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
