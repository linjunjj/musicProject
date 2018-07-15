package handler

import (
	"database/sql"
	"fmt"
	"musicProject/db"
	"reflect"
	"strings"
)

var defaultHandler handler
var _ DBInterface = defaultHandler

func DefaultHandler() DBInterface {
	return defaultHandler
}

type handler struct {
}

func (handler) Search(v interface{}, op Option) ([]interface{}, error) {
	fdMap, err := db.ParseStruct(v)
	if err != nil {
		return nil, err
	}
	conditon := ""
	var values []interface{}
	for tag, value := range fdMap {
		conditon += tag + " like ? and "
		values = append(values, value.Value.Interface())
	}
	if conditon != "" {
		conditon = conditon[:strings.LastIndex(conditon, "and")]
	}
	typ := reflect.TypeOf(v).Elem()
	rows, err := db.GetMysqlInstance().Debug().Model(reflect.New(typ).Interface()).Where(conditon, values...).Rows()
	if err != nil {
		return nil, err
	}
	var rets []interface{}
	for rows.Next() {
		val := reflect.New(typ).Interface()
		db.GetMysqlInstance().Debug().ScanRows(rows, val)
		rets = append(rets, val)
	}
	return rets, nil
}

func (handler) Query(v interface{}, op Option) ([]interface{}, error) {
	typ := reflect.TypeOf(v).Elem()
	rows, err := db.GetMysqlInstance().Debug().Model(reflect.New(typ).Interface()).Where(v).Offset(op.Skip).Limit(op.Limit).Rows()
	if err != nil {
		return nil, err
	}
	var rets []interface{}
	for rows.Next() {
		val := reflect.New(typ).Interface()
		db.GetMysqlInstance().Debug().ScanRows(rows, val)
		rets = append(rets, val)
	}
	return rets, nil
}

func (handler) Update(v interface{}) error {
	ret := db.GetMysqlInstance().Debug().Model(v).Update(v)
	if ret.Error != nil {
		return ret.Error
	}
	if ret.RowsAffected < 1 {
		return fmt.Errorf("update fail !")
	}
	return nil
}

func (handler) Insert(v interface{}) error {
	return db.GetMysqlInstance().Debug().Model(v).Create(v).Error
}

func (handler) Delete(v interface{}) error {
	ret := db.GetMysqlInstance().Debug().Model(v).Delete(v)
	if ret.Error != nil {
		return ret.Error
	}
	if ret.RowsAffected < 1 {
		return fmt.Errorf("delete fail !")
	}
	return nil
}

func (handler) Count(v interface{}) (int, error) {
	count := 0
	typ := reflect.TypeOf(v).Elem()
	err := db.GetMysqlInstance().Debug().Model(reflect.New(typ).Interface()).Where(v).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (handler) querySQL(sql string, resultType interface{}) ([]interface{}, error) {
	rows, err := db.GetMysqlInstance().Debug().Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	typ := reflect.TypeOf(resultType).Elem()
	var rets []interface{}
	for rows.Next() {
		val := reflect.New(typ).Interface()
		err := db.GetMysqlInstance().Debug().ScanRows(rows, val)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		rets = append(rets, val)
	}
	return rets, nil

}
func (handler) ScanRows(rows *sql.Rows, v interface{}) ([]interface{}, error) {
	typ := reflect.TypeOf(v).Elem()
	var rets []interface{}
	for rows.Next() {
		rv := reflect.New(typ)
		val := rv.Interface()
		err := db.GetMysqlInstance().Debug().ScanRows(rows, val)
		if err != nil {
			return nil, err
		}
		rets = append(rets, val)
	}
	return rets, nil
}

func (handler) GetCondition(interMap map[string]interface{}, bSearch bool) (string, []interface{}, error) {
	//var fdMap map[string]map[string]DB.StructField
	condition := ""
	var values []interface{}
	for k, v := range interMap {
		fdMap, err := db.ParseStruct(v)
		if err != nil {
			return condition, values, err
		}
		for tag, value := range fdMap {
			if value.IsBlank {
				continue
			}
			//用于获取string数组类的值
			if value.Struct.Type.Kind() == reflect.Ptr && value.Struct.Type.Elem().Kind() == reflect.Slice {
				vals := value.Value.Elem()
				if vals.Len() < 1 {
					continue
				}
				if bSearch {
					condition += k + "." + tag + " like ? and "
				} else {
					condition += k + "." + tag + " = ? and "
				}
				values = append(values, vals.Index(0).String())
				continue
			}
			if bSearch {
				condition += k + "." + tag + " like ? and "
			} else {
				condition += k + "." + tag + " = ? and "
			}
			values = append(values, value.Value.Interface())
		}
	}
	//fmt.Println(condition)
	return condition, values, nil
}
