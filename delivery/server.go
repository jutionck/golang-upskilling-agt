package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-upskilling-agt/config"
	"github.com/jutionck/golang-upskilling-agt/delivery/controller"
	"github.com/jutionck/golang-upskilling-agt/delivery/middleware"
	"github.com/jutionck/golang-upskilling-agt/manager"
	loggerutil "github.com/jutionck/golang-upskilling-agt/utils/logger_util"
)

type Server struct {
	uc            manager.UseCaseManager
	engine        *gin.Engine
	loggerService loggerutil.LoggerUtil
	host          string
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.loggerService).Logger())
	// semua controller di taruh disini
	controller.NewUserController(s.uc.UserUseCase(), s.engine)
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running %s", err.Error()))
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// manager
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		panic(err)
	}
	ur := manager.NewRepoManager(infraManager)
	uUc := manager.NewUseCaseManager(ur)
	engine := gin.Default()
	loggerService := loggerutil.NewLoggerUtil(cfg.FileConfig)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uc:            uUc,
		engine:        engine,
		loggerService: loggerService,
		host:          host,
	}
}
