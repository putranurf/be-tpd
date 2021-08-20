package migration

type Password struct {
	Id          int     `json:"id" `
	Newpassword *string `json:"new_password"`
}
