package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

type natsDataInterface interface {
	MajorTopic() string
}
type BaseResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

//具体返回给请求端的结构
type APIResponse struct {
	BaseResponse
	Data       []interface{} `json:"data"`
	Pagination Pagination    `json:"pagination"`
}

//返回的接口定义，页信息，总数等
type Pagination struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

/*
目的: 用于插入或是修改数据
方法: PUT
成功返回:
{
    "status": "ok",
    "msg": "ok"
}
失败返回:
{
    "status": "fail",
    "msg": "错误信息"
}
*/
type QueryHeader struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
type QueryResponse struct {
	QueryHeader
	BaseResponse
	Data []interface{} `json:"data"`
}

type Option struct {
	Limit int
	Skip  int
	Order string //key1 AEC, key2 DEC
}
type TableInterface interface {
	TableName() string
}

func HandlerAddOrUpdateInterface(ctx *gin.Context, natsInter natsDataInterface, method string) {
	majorTp := natsInter.MajorTopic()
	v := natsInter.(interface{})
	err := ctx.ShouldBindJSON(v)
	if err != nil {
		ctx.JSON(400, BaseResponse{Status: "fail", Msg: "参数格式有误"})
		return
	}
	majorTopic := fmt.Sprintf("%s", majorTp)
	logrus.Infof("request: %+v\n", v)
	queryH := &QueryHeader{Size: 999}
	v, ok := handlerInterfacesMap[majorTopic]
	hd, ok := v.(HandlerInterface)
	if !ok {
		err = fmt.Errorf("the handler %s is not exist !", majorTopic)
		return
	}
	var replyData interface{}
	switch method {
	case "GET", "get":
		fallthrough
	case "POST", "post":
		var rets []interface{}
		rets, err = hd.Query(v, Option{
			Limit: queryH.Size,
			Skip:  queryH.Page * queryH.Size,
		})
		resp := QueryResponse{
			BaseResponse: BaseResponse{
				Status: "ok",
				Msg:    "ok",
			},
			QueryHeader: QueryHeader{
				Page: queryH.Page,
				Size: len(rets),
			},
		}
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				resp.Data = []interface{}{}
			} else {
				return
			}
		} else {
			resp.Data = rets
		}

	case "update", "UPDATE":
		err = hd.Update(v)
		if err != nil {
			return
		}
		replyData =successResponse
	case "PUT", "put":
		err = hd.Insert(v)
		if err != nil {
			return
		}
		replyData =successResponse
	case "count", "COUNT":
		var count int
		count, err = hd.Count(v)
		if err != nil {
			return
		}
		replyData = []byte(fmt.Sprintf(`{"count":%d}`, count))
	case "delete", "DELETE":
		err = hd.Delete(v)
		if err != nil {
			return
		}
		replyData, _ = json.Marshal(successResponse)
	case "search", "SEARCH":
		var rets []interface{}
		rets, err = hd.Search(v, Option{
			Limit: queryH.Size,
			Skip:  queryH.Page * queryH.Size,
		})
		resp := QueryResponse{
			BaseResponse: BaseResponse{
				Status: "ok",
				Msg:    "ok",
			},
			QueryHeader: QueryHeader{
				Page: queryH.Page,
				Size: len(rets),
			},
		}
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				resp.Data = []interface{}{}
			} else {
				return
			}
		} else {
			resp.Data = rets
		}
		replyData = resp
	default:
		replyData, _ = json.Marshal(BaseResponse{
			Status: "fail",
			Msg:    fmt.Sprintf("%s method of  not support !", method),
		})
	}

	if err != nil {
		ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
		return
	}
	logrus.Infof("response: %+v\n", v)
	ctx.JSON(200,replyData)
}

/*
目的: 用于查询数据
方法: POST
成功返回: APIResponse结构
失败返回:
{
    "status": "fail",
    "msg": "错误信息"
}
*/
//func HandlerQueryInterface(ctx *gin.Context, natsInter natsDataInterface, method string) {
//	majorTp := natsInter.MajorTopic()
//	v := natsInter.(interface{})
//	err := ctx.ShouldBindJSON(v)
//	if err != nil {
//		ctx.JSON(400, BaseResponse{Status: "fail", Msg: "参数格式有误"})
//		return
//	}
//	logrus.Infof("request: %+v\n", v)
//	res, err := Request(fmt.Sprintf("%s.%s", majorTp, method), v)
//	if err != nil {
//		ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
//		return
//	}
//	var pageInfo Pagination
//	err = json.Unmarshal(res, &pageInfo) //得到微服务返回的结果
//	if err != nil {
//		ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
//		return
//	}
//	var resp APIResponse
//	err = json.Unmarshal(res, &resp) //得到微服务返回的结果
//	if err != nil {
//		ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
//		return
//	}
//	if method != "search" {
//		//获取总数量
//		res, err = Request(fmt.Sprintf("%s.count", majorTp), v)
//		if err != nil {
//			ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
//			return
//		}
//
//		total := &struct {
//			Count int `json:"count"`
//		}{
//			Count: 0,
//		}
//
//		err = json.Unmarshal(res, total)
//		if err != nil {
//			ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
//			return
//		}
//		pageInfo.Total = total.Count
//	} else {
//		pageInfo.Total = len(resp.Data)
//	}
//
//	resp.Pagination = pageInfo
//	ctx.JSON(200, resp)
//}

