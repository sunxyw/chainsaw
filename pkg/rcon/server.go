package rcon

import (
	"errors"
	"gohub/pkg/logger"

	rconpkg "github.com/gorcon/rcon"
	"go.uber.org/zap"
)

var (
	ErrNotConnected = errors.New("not connected")
)

type server struct {
	Name     string
	Host     string
	password string
	conn     *rconpkg.Conn
}

// newServer 根据配置创建一个新的服务器
func newServer(conf ServerConf) *server {
	return &server{
		Name:     conf.Name,
		Host:     conf.Host,
		password: conf.Password,
	}
}

// Connect 初始化连接
func (s *server) Connect() error {
	conn, err := rconpkg.Dial(s.Host, s.password)
	if err != nil {
		s.logError("Connect", err)
		return err
	}
	s.conn = conn
	return nil
}

// Send 发送命令并返回结果
func (s *server) Send(cmd string) (string, error) {
	if !s.IsConnected() {
		s.logError("Send", ErrNotConnected)
		return "", ErrNotConnected
	}

	resp, err := s.conn.Execute(cmd)
	if err != nil {
		s.logError("Send", err)
		return "", err
	}

	return resp, nil
}

// Close 关闭连接
func (s *server) Close() {
	s.conn.Close()
}

// IsConnected 是否连接
func (s *server) IsConnected() bool {
	return s.conn != nil
}

// logError 打印错误日志
func (s *server) logError(method string, err error) {
	logger.Error("Rcon Server Error", zap.String("server", s.Name), zap.String("method", method), zap.Error(err))
}
