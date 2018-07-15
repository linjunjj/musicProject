package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const (
	AccessKey = "WNiSKaCk1gu5mt3JtW5cqwhtBMD0pCvitUhzKRwI"
	SecretKey = "g9wBcCWARvYjEjyA8dHqsocB5qolKir-mqIvv-9l"
	Bucket    = "sancheng"
	Origin    = "http://os3kbkwao.bkt.clouddn.com/"
	imgPath   = "/Users/linjun/go/src/shangcheng-api/gowork-api/image/"
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
		replyData = resp
	case "update", "UPDATE":
		err = hd.Update(v)
		if err != nil {
			return
		}
		replyData = successResponse
	case "PUT", "put":
		err = hd.Insert(v)
		if err != nil {
			return
		}
		replyData = successResponse
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
	ctx.JSON(200, replyData)
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
func HandlerQueryInterface(ctx *gin.Context, natsInter natsDataInterface, method string) {
	majorTp := natsInter.MajorTopic()
	v := natsInter.(interface{})
	err := ctx.ShouldBindJSON(v)
	if err != nil {
		ctx.JSON(400, BaseResponse{Status: "fail", Msg: "参数格式有误"})
		return
	}
	logrus.Infof("request: %+v\n", v)
	queryH := &QueryHeader{Size: 999}
	v, ok := handlerInterfacesMap[majorTp]
	hd, ok := v.(HandlerInterface)
	if !ok {
		err = fmt.Errorf("the handler %s is not exist !", majorTp)
		return
	}

	var rets []interface{}
	rets, err = hd.Query(v, Option{
		Limit: queryH.Size,
		Skip:  queryH.Page * queryH.Size,
	})
	resps := QueryResponse{
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
			resps.Data = []interface{}{}
		} else {
			return
		}
	} else {
		resps.Data = rets
	}

	if err != nil {
		ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
		return
	}
	var pageInfo Pagination

	var resp APIResponse

	if method != "search" {
		//获取总数量
		count, err := hd.Count(v)
		if err != nil {
			ctx.JSON(400, BaseResponse{Status: "fail", Msg: err.Error()})
			return
		}

		total := &struct {
			Count int `json:"count"`
		}{
			Count: 0,
		}
		total.Count = count
		pageInfo.Total = total.Count
	} else {
		pageInfo.Total = len(resp.Data)
	}

	resp.Pagination = pageInfo
	ctx.JSON(200, resp)
}

func UploadImageInterface(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": 1,
			"msg":    "suc",
			"result": "shibai",
		})
		ctx.Abort()
		return
	}
	filePath := imgPath + file.Filename
	fileName := file.Filename
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": 1,
			"msg":    "文传上传失败",
			"result": "",
		})
		ctx.Abort()
		return
	}
	err1 := formUploader.PutFile(context.Background(), &ret, upToken, fileName, filePath, nil)
	if err1 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "文件上传失败",
			"result": "",
		})
		ctx.Abort()
		return
	}
	data := &Music{Src: Origin + fileName,Name:fileName}
	err = InsertMusic(data)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "文件上传成功，但插入数据库失败",
			"result": "",
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "suc",
		"result": Origin + fileName,
	})
	os.Remove(filePath)
}
