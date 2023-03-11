package accounts

type createUserBody struct {
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Full_name    string `json:"full_name"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password" binding:"required"`
}

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
