package user

type User struct {
	Id       int    `json:"id"`       //nolint:structtag
	Name     string `json:"name"`     //nolint:structtag
	Email    string `json:"email"`    //nolint:structtag
	Password string `json:"password"` //nolint:structtag
}
