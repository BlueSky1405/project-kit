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

// NewConn 初始化mysql连接对象, func() - 关闭方法
func NewConn(cfg *Config) (*Conn, func()) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Url, cfg.DatabaseName)
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
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
