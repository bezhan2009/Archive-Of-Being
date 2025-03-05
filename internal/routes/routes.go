package routes

import (
	_ "ArchiveOfBeing/docs"
	"ArchiveOfBeing/internal/controllers"
	"ArchiveOfBeing/internal/controllers/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// auth Маршруты для авторизаций
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
		auth.POST("/refresh", controllers.RefreshToken)
	}

	// Diary маршруты для дневников
	diary := r.Group("/diary", middlewares.CheckUserAuthentication)
	{
		diary.GET("", controllers.GetAllUserDiaries)
		diary.GET("/:id", controllers.GetDiaryByID)
		diary.POST("", controllers.CreateDiary)
		diary.PUT("/:id", controllers.UpdateDiary)
		diary.DELETE("/:id", controllers.DeleteDiary)
	}

	// Characters маршруты для Глав дневников
	character := r.Group("/character", middlewares.CheckUserAuthentication)
	{
		character.GET("diary/:id", controllers.GetCharacterByDiaryID)
		character.GET("/:id", controllers.GetCharacterByID)
		character.POST("", controllers.CreateCharacter)
		character.PUT("/:id", controllers.UpdateCharacter)
		character.DELETE("/:id", controllers.DeleteCharacter)
	}

	return r
}
