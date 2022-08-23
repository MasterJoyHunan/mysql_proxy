package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/MasterJoyHunan/mysql_proxy/config"
	"github.com/MasterJoyHunan/mysql_proxy/handle"
	"github.com/MasterJoyHunan/mysql_proxy/pkg/logger"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/server"
	"github.com/pkg/errors"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "etc/config.yaml", "default etc/config.yaml")
}

func main() {
	flag.Parse()
	config.Setup(configPath)
	handle.Setup()
	logger.Setup()
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.ProxyMysqlConf.Host, config.ProxyMysqlConf.Port))
	if err != nil {
		logger.Fatalf("%+v", errors.WithStack(err))
	}

	for {
		c, err := l.Accept()
		if err != nil {
			logger.Errorf("%+v", errors.WithStack(err))
			continue
		}

		conn, err := server.NewConn(c, config.ProxyMysqlConf.User, config.ProxyMysqlConf.Pwd, handle.ConsumerHandler{})
		if err != nil {
			if err != mysql.ErrBadConn {
				logger.Errorf("%+v", errors.WithStack(err))
			}
			continue
		}
		for {
			if err = conn.HandleCommand(); err != nil {
				logger.Errorf("%+v", errors.WithStack(err))
				break
			}
		}
	}
}
