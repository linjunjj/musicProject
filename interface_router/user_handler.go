package interface_router


func init() {
	RegistHandlerInterface(&UserMaster{})

}

type UserMaster struct {
	Id string `json:"id" gorm:"column:ID;primary_key"`
	UserName string `json:"user_name" gorm:"column:USER_NAME"`
	Age string `json:"age" gorm:"column:AGE"`
	Account string `json:"account" gorm:"column:ACCOUNT"`
	Password string `json:"password" gorm:"column:PASSWORD"`
}
func (*UserMaster) TableName() string {
	return "TBL_USER"
}

func (*UserMaster) MajorTopic() string {
	return "user"
}
func (*UserMaster) Query(v interface{}, op Option) ([]interface{}, error) {
	return defaultHandler.Query(v, op)
}

func (*UserMaster) Search(v interface{}, op Option) ([]interface{}, error) {
	return defaultHandler.Search(v, op)
}

func (*UserMaster) Update(v interface{}) error {
	return defaultHandler.Update(v)
}

func (*UserMaster) Insert(v interface{}) error {
	return defaultHandler.Insert(v)
}

func (*UserMaster) Delete(v interface{}) error {
	return defaultHandler.Delete(v)
}

func (*UserMaster) Count(v interface{}) (int, error) {
	return defaultHandler.Count(v)
}

