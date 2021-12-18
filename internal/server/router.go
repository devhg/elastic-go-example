package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/devhg/es/internal/conf"
	"github.com/devhg/es/internal/server/handler"
	"github.com/devhg/es/internal/service"
)

func NewHTTPRouter() *gin.Engine {
	r := gin.Default()

	gin.ForceConsoleColor()

	// http 接口注册
	setAPIRouter(r)

	return r
}

func setAPIRouter(r *gin.Engine) {
	// 接口注册
	api := r.Group("/api")

	userApi := api.Group("/user")
	{
		userApi.Handle(http.MethodPost, "/create", handler.Create)
		userApi.Handle(http.MethodPut, "/update", handler.Update)
		userApi.Handle(http.MethodDelete, "/delete", handler.Delete)
		userApi.Handle(http.MethodGet, "/info", handler.MGet)
		userApi.Handle(http.MethodPost, "/search", handler.Search)
	}
}

func NewHTTPServer(conf *conf.Config, userService *service.UserService) (*http.Server, func()) {
	r := NewHTTPRouter()
	server := &http.Server{
		Addr:           conf.Server.Addr,
		Handler:        r,
		ReadTimeout:    1000 * time.Millisecond,
		WriteTimeout:   1000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	// 手动注入依赖。todo 也许有更好的方式 比如 grpc 的自动生成，类似的生成http（proto-gen-http）
	handler.RegisterUserService(userService)

	return server, func() {
		err := server.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}
}
