package users

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name" validate:"required"`
    Password string `json:"password" validate:"required"`
    Username string `json:"username" validate:"required"`
}

type UserRegistrationRequest struct {
    Name     string `json:"name" validate:"required"`
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
}

type UserResponse struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}