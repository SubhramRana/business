package main

import (
	"business/Config"
	"business/Entities/CustomerEntity"
	"business/Entities/OrderEntity"
	"business/Entities/ProductEntity"
	"business/Router"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
func main(){
	//make a Mysql database connection
	dsn := Config.DBURL(Config.BuildDBConfig())
	var dbErr error
	if Config.DB , dbErr = gorm.Open(mysql.Open(dsn)) ; dbErr!=nil{
		log.Fatal(dbErr.Error())
	}
	//migrate the database with pre-defined schemas
	Config.DB.AutoMigrate(
		&CustomerEntity.Customer{},
		&ProductEntity.Product{},
		&OrderEntity.Order{},
	)

	//creating *gin.Engine
	server:=gin.Default()

	//setting up the routers
	Router.SetupRouter(server)

	//Run the server for accepting HTTP req
	if serverErr := server.Run(":8080") ; serverErr!=nil{
		log.Fatal(serverErr)
	}
}
