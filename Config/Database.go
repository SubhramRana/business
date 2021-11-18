package Config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type DBConfig struct{
	Host string
	Port uint
	Username string
	Password string
	DBName string
}

func BuildDBConfig() *DBConfig {
	return & DBConfig{
		Host: "localhost",
		Port: 3306,
		Username: "root",
		Password: "",
		DBName: "business_database",
		//DBName: "users_database",
	}
}

func DBURL(cnfg *DBConfig)(url string){

	url = fmt.Sprintf(
	"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cnfg.Username,
		cnfg.Password,
		cnfg.Host,
		cnfg.Port,
		cnfg.DBName,
	)

	return
}
