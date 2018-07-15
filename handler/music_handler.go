package handler

func init() {
	RegistHandlerInterface(&Music{})

}


type Music struct {
	Id string `json:"id" gorm:"column:ID;primary_key"`
	Src string `json:"src" gorm:"column:SRC"`
}
func (*Music) TableName() string {
	return "TBL_ACCOUNT_MASTER"
}

func (*Music) MajorTopic() string {
	return "merchant.account_master"
}

func (*Music) Query(v interface{}, op Option) ([]interface{}, error) {
	return defaultHandler.Query(v, op)
}

func (*Music) Search(v interface{}, op Option) ([]interface{}, error) {
	return defaultHandler.Search(v, op)
}

func (*Music) Update(v interface{}) error {
	return defaultHandler.Update(v)
}

func (*Music) Insert(v interface{}) error {
	return defaultHandler.Insert(v)
}

func (*Music) Delete(v interface{}) error {
	return defaultHandler.Delete(v)
}

func (*Music) Count(v interface{}) (int, error) {
	return defaultHandler.Count(v)
}
