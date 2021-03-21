package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
	"gitlab.com/sdk-go/api"
	"gitlab.com/sdk-go/config"
	"gitlab.com/sdk-go/library/logger"
	"gitlab.com/sdk-go/library/validator"
)

type SDKServer struct {
	g *gin.Engine
}

func newEngine() *gin.Engine {
	g := gin.Default()
	if err := validator.InitTrans("zh", "label"); err != nil {
		panic(err)
	}
	return g
}

func NewSDKServer() *SDKServer {
	g := newEngine()
	conf := config.GetAppConfig()
	gin.SetMode(gin.ReleaseMode)
	if gin.DebugMode == conf.RunMode {
		gin.SetMode(gin.DebugMode)
	}

	s := &SDKServer{g: g}
	//basic auth
	group := s.g.Group(conf.ApiPrefix, gin.BasicAuth(gin.Accounts{
		"admin": "admin@123",
	}))
	s.g.RouterGroup = *group

	//add middleware
	s.defaultGinMiddleware()

	//add pprof
	pprof.Register(s.g)

	//add router
	api.Register(s.g)

	return s
}

func (s *SDKServer) UseMiddleware(middleware ...gin.HandlerFunc) {
	s.g.Use(middleware...)
}

func (s *SDKServer) defaultGinMiddleware() {
	s.g.Use(gin.Recovery())
}

func (s *SDKServer) Run() {
	conf := config.GetAppConfig()
	logger.Setup("logs", config.ProjectName)
	logger.Info(fmt.Sprintf("Start to listening the incoming requests on http address: %s", conf.Addr))
	server := &http.Server{
		Addr:    conf.Addr,
		Handler: s.g,
	}
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(fmt.Printf("listen: %s", err.Error()))
		}
	}()
	gracefulStop(server)
}

// gracefulStop 优雅退出
// 等待中断信号以超时 5 秒正常关闭服务器
// 官方说明：https://github.com/gin-gonic/gin#graceful-shutdown-or-restart
func gracefulStop(server *http.Server) {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}
	logger.Info("Server exiting")
}
