package router

import (
	"fox/config"
	"fox/controller"
	"fox/middleware"
	"fox/repository"
	"fox/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatebaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	jwtService        service.JWTService           = service.NewJWTService()
	userService       service.UserService          = service.NewUserService(userRepository)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	rdb               config.IRedis                = config.NewRedisClient()
	wxTokenRepository repository.WxTokenRepository = repository.NewWxTokenConnection(db)
	wxTokenService    service.WxTokenService       = service.NewTokenService(wxTokenRepository, rdb.GetRdb())
	wxTokenController controller.WxTokenController = controller.NewWxTokenController(wxTokenService)
	tickerToken       service.TickerToken          = service.NewTickerTokenService(rdb.GetRdb(), wxTokenService)
	wxUserRepository  repository.WxUserRepository  = repository.NewWxUserRepository(db)
	wxUserService     service.WxUserService        = service.NewWwxUserService(wxUserRepository)
	wxUserController  controller.WxUserController  = controller.NewWxUserController(wxUserService)
)

func InitRouter() {
	defer config.CloseDatabaseConnection(db)
	defer config.CloseRedisClient(rdb.GetRdb())

	server := gin.Default()
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env files")
	}

	authRouts := server.Group("api/auth")
	{
		authRouts.POST("/login", authController.Login)
		authRouts.POST("/register", authController.Register)
		authRouts.GET("/getRedisToken", wxTokenController.GetRedisAccessToken)
		authRouts.GET("/getToken", wxTokenController.GetDbAccessToken)
		authRouts.POST("/wxUser", wxUserController.Insert)
	}
	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/update", userController.Update)
	}
	server.Run(":9090")
}
