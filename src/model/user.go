package model

type User struct {
	Id   int
	Name string
	Pass string
	Del  int
}

//表示配置操作数据库名称
func (u User) TableName() string {
	return "user"
}
