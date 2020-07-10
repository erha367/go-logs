package database

import (
	"eeo_webcast_service/application/library"
	"eeo_webcast_service/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
)

const Mater = "master"
const Slave = "slave"

//数据库操作日志
type databaseLogger struct{}

//实现gorm定义的日志处理接口
func (log *databaseLogger) Print(v ...interface{}) {
	var (
		level    = v[0]
		position = v[1]
	)
	if level == "sql" {
		library.Logger.Debug("数据库日志：",
			zap.String("SQL", v[3].(string)),
			zap.String("position", position.(string)))
	} else {
		library.Logger.Error("数据库错误：", zap.String("log", fmt.Sprintf("%v",v[2])))
	}
}

//定义数据库存放容器
var databaseContainer = make(map[string]*gorm.DB)

//初始化数据库
func InitDatabases() {
	dialect := config.App.Databases.Driver
	for _, database := range config.App.Databases.List {
		dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
			database.Username,
			database.Password,
			database.Host,
			database.Port,
			database.Name)
		conn, openErr := gorm.Open(dialect, dns)
		if openErr != nil {
			panic("数据库链接失败，原因：" + openErr.Error())
		}
		conn.DB().SetMaxOpenConns(database.MaxOpenConnections)
		conn.DB().SetMaxIdleConns(database.MaxIdleConnections)
		if config.Environment != config.EnvProd {
			conn.LogMode(true)
			conn.SetLogger(&databaseLogger{})
		}
		conn.SingularTable(true)
		databaseContainer[database.Name+"."+database.Type] = conn
	}
	log.Info(`数据库连接成功`)
}

//获取数据库主库链接
func GetMasterDB(dbName string) *gorm.DB {
	return databaseContainer[dbName+"."+Mater]
}

//获取数据库从库链接,如果从库不存在，则从主库获取，解决只有主库的情况
func GetSlaveDB(dbName string) *gorm.DB {
	if conn, ok := databaseContainer[dbName+"."+Slave]; ok {
		return conn
	} else {
		return databaseContainer[dbName+"."+Mater]
	}
}

//从主库开启一个事务
func Begin(dbName string) *gorm.DB {
	return GetMasterDB(dbName).Begin()
}

//关闭所有链接
func CloseDatabases() {
	for key, conn := range databaseContainer {
		if err := conn.DB().Close(); err != nil {
			library.Logger.Error("关闭数据库链接失败:", zap.String("database", key), zap.Error(err))
		} else {
			library.Logger.Info("成功关闭数据库", zap.String("database", key))
		}
	}
}
