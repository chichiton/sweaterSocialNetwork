package main

import (
	"github.com/chichiton/sweaterSocialNetwork/controllers"
	docs "github.com/chichiton/sweaterSocialNetwork/docs"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/db_connector"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/db_migrations"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/middleware"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	dbConfig := db_connector.DbConfig{}
	dbConfig.SetDbConfig(gin.Mode())

	dbConnector := db_connector.NewMySqlConnector(dbConfig)

	db_migrations.CreateDatabase(dbConfig)
	db_migrations.Migrate(dbConnector, dbConfig)

	userRepository := repositories.NewUserRepositoryImp(dbConnector)

	friendRepository := repositories.NewFriendRepository(dbConnector)

	authMiddleware, err := middleware.NewAuthMiddleware(userRepository).GetInstance("test_token_secret", "id")

	if err != nil {
		panic("auth middleware error")
	}

	r.Use(gin.Recovery())

	r.Use(gin.ErrorLogger())

	apiPublic := r.Group("/api/v1")
	{
		apiPublic.POST("/auth", authMiddleware.LoginHandler)
		registerController := controllers.NewRegisterController(userRepository)
		apiPublic.POST("/register", registerController.RegisterUser)
	}

	apiSecure := r.Group("/api/v1")
	{
		apiSecure.Use(authMiddleware.MiddlewareFunc())

		userController := controllers.NewUserController(userRepository)
		apiSecure.GET("/user/profile", userController.GetUserProfile)

		friendController := controllers.NewFriendController(friendRepository)
		apiSecure.POST("/user/friend", friendController.AddFriend)
		apiSecure.GET("/user/friend", friendController.GetFriends)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
