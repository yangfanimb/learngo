package rpcsupport

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"learngo/crawler/helper/log"
)

var logger = log.DLogger()

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	logger.Info("listening on %s", host)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("accept err: %v", err)
			continue
		}
		logger.Info("accept connection")

		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
