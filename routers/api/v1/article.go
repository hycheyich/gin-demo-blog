package v1

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		//	 表示校验通过
		if models.ExistArticleByID(id) {
			//	表示存在ID的文章
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		//	表示存在错误
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})

}

// 获取多个文章
func GetArticles(c *gin.Context) {

}

// 新增文章
func AddArticle(c *gin.Context) {

}

// 修改文章
func EditArticle(c *gin.Context) {

}

//删除文章
func DeleteArticle(c *gin.Context) {

}
