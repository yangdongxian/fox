package main

import (
	"fmt"
	"fox/config"
	"fox/controller"
	"fox/middleware"
	"fox/repository"
	"fox/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

var (
	db             *gorm.DB                  = config.SetupDatebaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env files")
	}

	fmt.Println("secret:", os.Getenv("SECRET"))
	fmt.Println("GOPATH:", os.Getenv("GOPATH"))

	//

	authRouts := server.Group("api/auth")
	{
		authRouts.POST("/login", authController.Login)
		authRouts.POST("/register", authController.Register)
	}
	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/update", userController.Update)
	}
	server.Run(":9090")
}
