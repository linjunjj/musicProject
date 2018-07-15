package handler

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	handlerInterfacesMap = make(map[string]interface{})
	successResponse      = BaseResponse{Status: "ok", Msg: "ok"}
	errorMethodNotExist  = errors.New("method not exist !")
)

type HandlerInterface interface {
	DBInterface
	InitInterface
	TableInterface
}

type DBInterface interface {
	Query(interface{}, Option) ([]interface{}, error)  //用于get或是post的精确字段查询，并且指定需要条数
	Search(interface{}, Option) ([]interface{}, error) //用于模糊查询，并且指定需要条数
	Update(interface{}) error                          //更新数据
	Insert(interface{}) error                          //插入数据
	Delete(interface{}) error                          //删除数据
	Count(interface{}) (int, error)                    //返回当前条件的总条数
}

type InitInterface interface {
	MajorTopic() string //订阅的topic前部分
}

func RegistHandlerInterface(h HandlerInterface) {
	if reflect.ValueOf(h).Kind() != reflect.Ptr {
		t := h.(TableInterface)
		panic(fmt.Sprintf("the HandlerInterface %s must be a pointer !", strings.TrimPrefix(t.TableName(), "TBL_")))
	}
	handlerInterfacesMap[h.MajorTopic()] = h.(interface{})
}
