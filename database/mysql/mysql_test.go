package mysql

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMysql(t *testing.T) {
	cfg := &Config{
		Username:     "root",
		Password:     "1234",
		Url:          "127.0.0.1:3306",
		Port:         "3306",
		DatabaseName: "mysql",
	}

	conn, f := NewConn(cfg)
	defer f()

	db := conn.GetDB()

	var res struct {
		User string `json:"user" gorm:"column:user"`
		Host string `json:"host" gorm:"column:host"`
	}
	err := db.Table("user").Scan(&res).Error
	require.Nil(t, err)
	require.Equal(t, "root", res.User)
	require.Equal(t, "localhost", res.Host)
}
