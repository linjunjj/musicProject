package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"strings"
	"fmt"
	"github.com/sirupsen/logrus"
)

var mysqlDb *gorm.DB

func NewMysqlDriver(dns string) *gorm.DB{
	condition := "timeout=60s&parseTime=True&charset=utf8mb4,utf8"
	if strings.Contains(dns, "?") {
		dns = dns + "&" + condition
	} else {
		dns = dns + "?" + condition
	}
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		panic("Cannt Connecting Mysql Data Source:" + err.Error())
	}
	db.DB().SetMaxOpenConns(256)
	db.DB().SetMaxIdleConns(8)
	db.DB().SetConnMaxLifetime(360 * time.Second)
	return db
}
func InitMysql(dns string){
	mysqlDb = NewMysqlDriver(dns)
	mysqlDb.BlockGlobalUpdate(true) //禁止无condition进行更新或删除
	mysqlDb.Callback().Update().Before("gorm:update").Register("valid_primary_key", ValidPrimaryKey)
	mysqlDb.SetLogger(logger{})
}
func ValidPrimaryKey(scope *gorm.Scope) {
	fds := scope.PrimaryFields()
	errTag := ""
	for _, value := range fds {
		if value.IsBlank {
			errTag += value.Tag.Get("json") + ", "
		}
	}
	if errTag != "" {
		scope.Err(fmt.Errorf("the primary key (%s) is empty !", errTag))
	}
}
func GetMysqlInstance() *gorm.DB {
	return mysqlDb
}
type logger struct {
}

func (logger) Print(values ...interface{}) {
	logrus.Info(values)
}