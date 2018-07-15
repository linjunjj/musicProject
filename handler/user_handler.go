package handler

import (
	"strings"
	"fmt"
	"musicProject/db"
)

func init() {
	RegistHandlerInterface(&User{})

}
type StringArray []string

type User struct {
	UserMaster
	JionData
}


type UserMaster struct {
	Id string `json:"id" gorm:"column:ID;primary_key"`
	UserName string `json:"user_name" gorm:"column:USER_NAME"`
	Age string `json:"age" gorm:"column:AGE"`
	Account string `json:"account" gorm:"column:ACCOUNT"`
	Password string `json:"password" gorm:"column:PASSWORD"`
}
type JionData struct {
	LikeMusic               *StringArray `json:"like_music" gorm:"column:LIKE_SRC"`                               //交易订单号
}


func (*UserMaster) TableName() string {
	return "TBL_USER"
}

func (*User) MajorTopic() string {
	return "user"
}
func (*User) Query(v interface{}, op Option) ([]interface{}, error) {
	req := v.(*User)
	sql := `select ci.*, m.src
			from TBL_USER as u
		    left join TBL_MUSIC as m on m.id = rm.MUSIC_ID
            left join TBL_RELATEMUSIC as rm on rm.USER_ID = u.id`
	structMap := map[string]interface{}{
		"u":    &req.UserMaster,
	}
	condition, values, err := defaultHandler.GetCondition(structMap, false)
	if condition != "" {
		condition = condition[:strings.LastIndex(condition, "and")]
		condition = strings.Replace(condition, "u.id = ?", "instr(u.id , ?)", 1)
		sql += " where " + condition
	}
	sql += fmt.Sprintf(" limit %d,%d", op.Skip, op.Limit)
	rows, err := db.GetMysqlInstance().Debug().Raw(sql, values...).Rows()
	if err != nil {
		return nil, err
	}
	return defaultHandler.ScanRows(rows, &User{})
}

func (*User) Search(v interface{}, op Option) ([]interface{}, error) {
	return defaultHandler.Search(v, op)
}

func (*User) Update(v interface{}) error {
	return defaultHandler.Update(v)
}

func (*User) Insert(v interface{}) error {
	return defaultHandler.Insert(v)
}

func (*User) Delete(v interface{}) error {
	return defaultHandler.Delete(v)
}

func (*User) Count(v interface{}) (int, error) {
	return defaultHandler.Count(v)
}

