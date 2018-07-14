package db

import (
	"git.coding.net/greatLIU/utils/db"
	"fmt"
	"musicProject/config"
	"reflect"
	"strings"
	"strconv"
)

var tar *db.Conn
var err error

func Init() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	//连接到tarantool数据库
	tar, err = db.Dial(config.GetValue("tarantool_addr").ToString(), config.GetValue("tarantool_username").ToString(), config.GetValue("tarantool_password").ToString())
	if err != nil {
		panic(err)
	}

}

func Init_mysql() {
	//连接到mysql数据库
	mysqlHost := config.GetValue("mysql_host").ToString()         // mysql host
	mysqlDB := config.GetValue("mysql_db").ToString()             // mysql databases
	mysqlUser := config.GetValue("mysql_user").ToString()         // mysql user
	mysqlPassword := config.GetValue("mysql_password").ToString() // mysql password
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDB)
	InitMysql(connInfo)
}

func GetDBInstance() *db.Conn {
	return tar
}
type StructField struct {
	Struct    reflect.StructField
	Value     reflect.Value
	Name      string
	OrmTag    string
	JsonTag   string
	SearchTag string
	IsBlank   bool
	Num       string
}
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}


func ParseStruct(v interface{}) (map[string]StructField, error) {
	structMap := make(map[string]StructField)
	reflectType := reflect.TypeOf(v).Elem()
	reflectValue := reflect.ValueOf(v)
	reflectKind := reflectType.Kind()
	if reflectKind != reflect.Struct {
		return nil, fmt.Errorf("not support non-struct type: %s !", reflectKind.String())
	}
	for i := 0; i < reflectType.NumField(); i++ {
		fieldStruct := reflectType.Field(i)
		fieldValue := reflectValue.Elem().Field(i)
		if fieldStruct.Type.Kind() == reflect.Struct {
			rfValue := reflect.New(reflect.PtrTo(fieldStruct.Type))
			rfValue.Elem().Set(fieldValue.Addr())
			tagMap, err := ParseStruct(rfValue.Elem().Interface())
			if err != nil {
				return nil, err
			}
			for key, value := range tagMap {
				structMap[key] = value
			}
			continue
		}
		ormTag := fieldStruct.Tag.Get("orm")
		jsonTag := fieldStruct.Tag.Get("json")
		searchTag := fieldStruct.Tag.Get("search")
		if strings.Contains(ormTag, "-") {
			continue
		}
		field := StructField{
			Struct:    fieldStruct,
			Name:      fieldStruct.Name,
			OrmTag:    ormTag,
			JsonTag:   jsonTag,
			SearchTag: searchTag,
			Value:     fieldValue,
			IsBlank:   isBlank(fieldValue),
			Num:       strconv.Itoa(i),
		}
		//		fmt.Println("field:", field)
		if searchTag != "" {
			structMap[searchTag] = field
		} else if ormTag != "" && ormTag != "primary_key" {
			structMap[ormTag] = field
		} else {
			structMap[jsonTag] = field
		}
	}
	return structMap, nil
}