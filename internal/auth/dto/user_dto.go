package dto

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
