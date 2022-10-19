package mysql

import "fmt"

func InitMigrateCmd(config *Config, do string, tables ...interface{}) {
	migrateORM, cleanup := NewMigrateConn(config)
	defer cleanup()
	var orm *Conn
	if do == "migrate" {
		tmpORM, cleanup2 := NewConn(config)
		orm = tmpORM
		defer cleanup2()
	}

	switch do {
	case "create":
		migrateORM.CreateDB()
	case "drop":
		migrateORM.DropDB()
	case "migrate":
		if err := orm.GetDB().AutoMigrate(tables...); err != nil {
			panic(any(err))
		}
	default:
		panic(any(fmt.Sprintf("no support do:%s", do)))
	}
}
