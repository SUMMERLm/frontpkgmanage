package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var done bool
var doScale = "0"
var doQuotaScale = "0"

type Hpa_msg struct {
	Namespace  string `json:"namespace"`
	Name       string `json:"name"`
	Hpa_action string `json:"hpaAction"`
}

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
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		//TODO 结合nginx vhost管理配置
		_ = "./" + file.Filename
		// 删除指定目录下指定文件

	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
