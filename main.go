package main

import (
	"musicProject/config"
	"musicProject/db"
	"musicProject/gin_router"
	_"musicProject/router"
)

func init() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	db.Init_mysql()
}
func main() {

	gin_router.GetEngine().Run(":8080")

}
