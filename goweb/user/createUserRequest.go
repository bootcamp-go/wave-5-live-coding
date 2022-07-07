package user

type CreateUserRequest struct {
	Id        int     `json:"id"`
	FirstName string  `json:"firstName"  binding:"required"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Age       int     `json:"age"`
	Height    float64 `json:"height"`
}
