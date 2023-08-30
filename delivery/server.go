package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-upskilling-agt/config"
	"github.com/jutionck/golang-upskilling-agt/delivery/controller"
	"github.com/jutionck/golang-upskilling-agt/delivery/middleware"
	"github.com/jutionck/golang-upskilling-agt/manager"
	"github.com/jutionck/golang-upskilling-agt/usecase"
	loggerutil "github.com/jutionck/golang-upskilling-agt/utils/logger_util"
	"github.com/jutionck/golang-upskilling-agt/utils/service"
)

type Server struct {
	uc            manager.UseCaseManager
	authService   usecase.AuthUseCase
	engine        *gin.Engine
	loggerService loggerutil.LoggerUtil
	jwtService    service.JwtService
	host          string
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.loggerService).Logger())
	// semua controller di taruh disini
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewUserController(s.uc.UserUseCase(), s.engine, authMiddleware)
	controller.NewAuthController(s.engine, s.authService)
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
	jwtService := service.NewJwtService(cfg.JwtConfig)
	authUseCase := usecase.NewAuthUseCase(uUc.UserUseCase(), jwtService)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uc:            uUc,
		engine:        engine,
		loggerService: loggerService,
		jwtService:    jwtService,
		authService:   authUseCase,
		host:          host,
	}
}
