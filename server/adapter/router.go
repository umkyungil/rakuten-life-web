package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rakuten-life-web/server/adapter/controllers"
	"rakuten-life-web/server/entity/util"
	"time"
)

func Init() {
	r := Router(false)
	r.Run()
}

func Router(gae bool) *gin.Engine {
	//CORS対応
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{util.CORS_KENMEI_RELEASE, util.CORS_BPMG_API_RELEASE},
		AllowOrigins:     []string{util.CORS_ALL},
		AllowMethods:     []string{"PUT", "DELETE", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := r.Group("v1")
	{
		ctrl := controllers.Controllers{}

		// AutoMigration
		//v1.GET("/autoMigration", ctrl.AutoMigration)
		// 顧客BP情報管理システム
		//v1.GET("/purchase/authorization", ctrl.GetByIdController)
		//v1.GET("/purchase/search/:id", ctrl.GetByIdController)
		//v1.POST("/purchase/search", ctrl.GetByConditionsController)
		//v1.POST("/purchase/insert", ctrl.CreateModelController)
		//v1.PUT("/purchase/update/:id", ctrl.UpdateByIdController)
		//v1.DELETE("/purchase/delete/:id", ctrl.DeleteByIdController)
		v1.POST("/purchase/csv", ctrl.CsvController)
		// 件名システム
		//v1.GET("/purchase/vendors", ctrl.GetByAllOfVendorsController)
		//v1.GET("/purchase/vendor", ctrl.GetByIdOfVendorController)
	}

	return r
}
