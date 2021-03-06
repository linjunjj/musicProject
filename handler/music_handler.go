package handler

func init() {
	RegistHandlerInterface(&Music{})

}

type Music struct {
	Id  string `json:"id" gorm:"column:ID;primary_key"`
	Src string `json:"src" gorm:"column:SRC"`
	Name string `json:"name" gorm:"column:NAME"`
}

func (*Music) TableName() string {
	return "TBL_MUSIC"
}

func (*Music) MajorTopic() string {
	return "music"
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

func InsertMusic(v interface{}) error {
	return defaultHandler.Insert(v)
}
