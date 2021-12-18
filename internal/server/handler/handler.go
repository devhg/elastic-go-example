package handler

import (
	"github.com/devhg/es/internal/service"
)

var (
	userService *service.UserService
)

// RegisterUserService 服务注册
func RegisterUserService(s *service.UserService) {
	userService = s
}
