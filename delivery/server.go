package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-upskilling-agt/config"
	"github.com/jutionck/golang-upskilling-agt/delivery/controller"
	"github.com/jutionck/golang-upskilling-agt/delivery/middleware"
	"github.com/jutionck/golang-upskilling-agt/docs"
	"github.com/jutionck/golang-upskilling-agt/manager"
	"github.com/jutionck/golang-upskilling-agt/usecase"
	loggerutil "github.com/jutionck/golang-upskilling-agt/utils/logger_util"
	"github.com/jutionck/golang-upskilling-agt/utils/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	controller.NewUserController(s.uc.UserUseCase(), s.engine, authMiddleware).Route()
	controller.NewAuthController(s.engine, s.authService).Route()
}

func (s *Server) swageDocs() {
	docs.SwaggerInfo.Title = "UPSKILLING GOLANG"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.Host = "localhost:8888"
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Server) Run() {
	s.setupControllers()
	s.swageDocs()
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
