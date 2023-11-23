package route

// 路由包

import (
	"frontpkgmanage/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// get Test
	router.GET("/", api.GetIndex)
	// post Test
	router.POST("/test", api.PostTest)
	//pkg upload
	router.POST("/front/single/pkg/upload", api.Upload)
	//multi pkg upload
	router.POST("/front/multi/pkg/upload", api.UploadMultiPkg)
	//pkg recycle
	router.POST("/front/pkg/recycle", api.Recycle)
	router.GET("/downloadFiles", api.DownloadFileService)
	return router
}
