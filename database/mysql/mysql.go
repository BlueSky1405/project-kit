package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	// db username
	Username string `json:"username"`
	// db password
	Password string `json:"password"`
	// db url
	Url string `json:"url"`
	// db port
	Port string `json:"port"`
	// database name
	DatabaseName string `json:"database_name"`
	// db MaxIdleConn
	//MaxIdleConn int
	// db MaxOpenConn
	//MaxOpenConn int
	// db ConnMaxLifeTime
	//ConnMaxLifeTime string
}

// Conn mysql连接封装对象，方便后期扩展方法
type Conn struct {
	cfg  *Config
	conn *gorm.DB
}

func (c *Config) dialector(initDB bool) gorm.Dialector {
	if initDB {
		return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/",
			c.Username, c.Password, c.Url))
	}
	// mysql: url 示例: 127.0.0.1:3306(包含端口)
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Url, c.DatabaseName))
}

func NewConn(cfg *Config) (*Conn, func()) {
	return newConn(cfg, false)
}

func NewMigrateConn(cfg *Config) (*Conn, func()) {
	return newConn(cfg, true)
}

// NewConn 初始化mysql连接对象, func() - 关闭方法
func newConn(cfg *Config, initDb bool) (*Conn, func()) {
	db, err := gorm.Open(cfg.dialector(initDb))
	if err != nil {
		panic(any(err))
	}
	// test ping
	sqlDb, err := db.DB()
	if err != nil {
		panic(any(err))
	}
	cleanUp := func() {
		// no handler err
		_ = sqlDb.Close()
	}
	return &Conn{
		cfg:  cfg,
		conn: db,
	}, cleanUp
}

func (c *Conn) GetDB() *gorm.DB {
	return c.conn
}

func (c *Conn) CreateDB() {
	createDbSQL := "CREATE DATABASE IF NOT EXISTS " + c.cfg.DatabaseName + " DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;"
	err := c.GetDB().Exec(createDbSQL).Error
	if err != nil {
		panic(any("创建失败：" + err.Error() + " sql:" + createDbSQL))
	}
	fmt.Println(fmt.Sprintf("database: %v create success", c.cfg.DatabaseName))
}

func (c *Conn) DropDB() {
	dropDbSQL := "DROP DATABASE IF EXISTS " + c.cfg.DatabaseName + ";"
	err := c.GetDB().Exec(dropDbSQL).Error
	if err != nil {
		panic(any("删除失败：" + err.Error() + " sql:" + dropDbSQL))
	}
	fmt.Println(fmt.Sprintf("database: %v drop success", c.cfg.DatabaseName))
}
