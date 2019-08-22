package main

import(
	"fmt"
	"github.com/go-xorm/xorm"
	"bitbucket.org/falabellafif/ExampleProject/internal/config"
	_ "github.com/go-sql-driver/mysql"
)


func CreateDatabaseConnection(c *config.Db) *xorm.Engine {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", c.User, c.Pass, c.Server, c.Port, c.Db)
	engine, errOpen := xorm.NewEngine("mysql", connectionString)
	if errOpen != nil {
		fmt.Println(errOpen)
		panic(errOpen.Error())
	}
	errPing := engine.Ping()
	if errPing != nil {
		fmt.Println(errPing)
		panic(errPing.Error())
	}
	return engine
}