package Models

type Ms_user struct {
	User_Id  int    `json:"user_id  "`
	Username string `json:"username"`
	Userpass string `json:"userpass"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
}

func (b *Ms_user) TableName() string {
	return "ms_user"
}
