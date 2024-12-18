package user

type ResponseDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewResponseDto(user User) *ResponseDto {
	return &ResponseDto{ID: user.ID, Name: user.Name, Email: user.Email}
}
