package handler

func init() {
	RegistHandlerInterface(&RelateMusic{})
}

type RelateMusic struct {
	Id      string `json:"id" gorm:"column:ID;primary_key"`
	UserId  string `json:"user_id" gorm:"column:USER_ID" `
	MusicId string `json:"music_id" gorm:"column:MUSIC_ID"`
}

func (*RelateMusic) TableName() string {
	return "TBL_RELATEMUSIC"
}

func (*RelateMusic) MajorTopic() string {
	return "relateMusic"
}

func (*RelateMusic) Query(v interface{}, op Option) ([]interface{}, error) {
	return defaultHandler.Query(v, op)
}

func (*RelateMusic) Search(v interface{}, op Option) ([]interface{}, error) {
	return defaultHandler.Search(v, op)
}

func (*RelateMusic) Update(v interface{}) error {
	return defaultHandler.Update(v)
}

func (*RelateMusic) Insert(v interface{}) error {
	return defaultHandler.Insert(v)
}

func (*RelateMusic) Delete(v interface{}) error {
	return defaultHandler.Delete(v)
}

func (*RelateMusic) Count(v interface{}) (int, error) {
	return defaultHandler.Count(v)
}
