package auth_context

type RefreshInputDTO struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
