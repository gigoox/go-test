package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestSpec_db(t *testing.T){
	Convey("Test load database config", t, func() {
		Convey("Test load database config ok", func() {
			driver, server, user, pass, port, db := "psql", "128.342.345.1", "admin_root", "root123", "9876", "gestor_prod"
			os.Setenv("GO_DB_DRIVER", driver)
			os.Setenv("GO_DB_SERVER", server)
			os.Setenv("GO_DB_USER", user)
			os.Setenv("GO_DB_PASS", pass)
			os.Setenv("GO_DB_PORT", port)
			os.Setenv("GO_DB_DB", db)

			dbConfig, _ := NewDb()

			So(dbConfig.Driver, ShouldEqual, driver)
			So(dbConfig.Server, ShouldEqual, server)
			So(dbConfig.User, ShouldEqual, user)
			So(dbConfig.Pass, ShouldEqual, pass)
			So(dbConfig.Port, ShouldEqual, port)
			So(dbConfig.Db, ShouldEqual, db)
		})
	})
}