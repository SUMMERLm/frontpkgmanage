package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func GetIndex(c *gin.Context) {
	c.String(200, "Hello, get test")

}

func PostTest(c *gin.Context) {
	c.String(200, "Hello，post test")
}

func Upload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	//TODO 结合nginx vhost管理配置
	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func UploadMultiPkg(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		//TODO 结合nginx vhost管理配置
		dst := "./" + file.Filename
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func Recycle(c *gin.Context) {
	// Multipart form
	fileDir := c.Query("filedir")
	fileName := c.Query("fileName")
	_, errByOpenFile := os.Open(fileDir + "/" + fileName)
	//非空处理
	if errByOpenFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "失败",
			"error":   "资源不存在",
		})
		c.Redirect(http.StatusFound, "/404")
		return
	}
	// 删除指定目录下指定文件
	os.Remove(fileDir + "/" + fileName)
}

func DownloadFileService(c *gin.Context) {
	fileDir := c.Query("filedir")
	fileName := c.Query("fileName")
	_, errByOpenFile := os.Open(fileDir + "/" + fileName)
	//非空处理
	if errByOpenFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "失败",
			"error":   "资源不存在",
		})
		c.Redirect(http.StatusFound, "/404")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(fileDir + "/" + fileName)
	return
}
