package Models

type Tb_todo struct {
	Todo_Id    int    `json:"todo_id "`
	Todo_Title string `json:"todo_title"`
	Todo_Desc  string `json:"todo_desc"`
	User_Id    string `json:"user_id"`
}

func (b *Tb_todo) TableName() string {
	return "tb_todo"
}
