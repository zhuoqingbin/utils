package gormv2

import (
	"time"

	"github.com/zhuoqingbin/utils/flags"
)

var (
	mysqlHost            = flags.String("mysql_host", "mysql", "mysql server address. default: mysql")
	mysqlPort            = flags.Int("mysql_port", 3306, "mysql server port. default: 3306")
	mysqlUser            = flags.StringRequired("mysql_user", "mysql user.")
	mysqlPasswd          = flags.StringRequired("mysql_passwd", "mysql passwd.")
	mysqlMaxIdleConns    = flags.Int("mysql_max_idle_conns", 10, "mysql max idle conns")
	mysqlMaxOpenConns    = flags.Int("mysql_max_open_conns", 100, "mysql max open conns")
	mysqlConnMaxLifetime = flags.Duration("mysql_conn_max_lifetime", 300*time.Second, "mysql max open conns")
)
