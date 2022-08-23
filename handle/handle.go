package handle

import (
	"fmt"

	"github.com/MasterJoyHunan/mysql_proxy/config"
	"github.com/MasterJoyHunan/mysql_proxy/pkg/logger"
	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
)

var conn *client.Conn

func Setup() {
	conn, _ = client.Connect(
		fmt.Sprintf("%s:%d", config.RealMysqlConf.Host, config.RealMysqlConf.Port),
		config.RealMysqlConf.User,
		config.RealMysqlConf.Pwd,
		config.RealMysqlConf.Db)
}

type ConsumerHandler struct{}

func (h ConsumerHandler) UseDB(dbName string) error {
	logger.Logger.WithField("cmd", "UseDB").Info(dbName)
	return nil
}
func (h ConsumerHandler) HandleQuery(query string) (*mysql.Result, error) {
	logger.Logger.WithField("cmd", "HandleQuery").Info(query)
	execute, err := conn.Execute(query)
	if err != nil {
		logger.Logger.WithField("cmd", "HandleQuery").Error(err.Error())
	}
	return execute, err
}

func (h ConsumerHandler) HandleFieldList(table string, fieldWildcard string) ([]*mysql.Field, error) {
	logger.Logger.WithField("cmd", "HandleFieldList [table]").Info(table)
	logger.Logger.WithField("cmd", "HandleFieldList [fieldWildcard]").Info(fieldWildcard)
	return nil, nil
}
func (h ConsumerHandler) HandleStmtPrepare(query string) (int, int, interface{}, error) {
	logger.Logger.WithField("cmd", "HandleStmtPrepare").Info(query)
	return 0, 0, nil, nil
}
func (h ConsumerHandler) HandleStmtExecute(context interface{}, query string, args []interface{}) (*mysql.Result, error) {
	logger.Logger.WithField("cmd", "HandleStmtExecute [query]").Info(query)
	logger.Logger.WithField("cmd", "HandleStmtExecute [args]").Info(args)
	return nil, nil
}

func (h ConsumerHandler) HandleStmtClose(context interface{}) error {
	return nil
}

func (h ConsumerHandler) HandleOtherCommand(cmd byte, data []byte) error {
	logger.Logger.WithField("cmd", "HandleOtherCommand [cmd]").Info(cmd)
	logger.Logger.WithField("cmd", "HandleOtherCommand [data]").Info(string(data))
	return nil
}
