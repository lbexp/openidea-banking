package user_model

type UserData struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type UserResponse struct {
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}
