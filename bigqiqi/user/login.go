package user

import (
	"bigqiqi/utils"
	"bigqiqi/utils/errorcode"

	//"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	/*name := c.Query("name")
	password := c.Query("password")
	name1 := c.PostForm("name")
	password1 := c.PostForm("password")
	body, _ := ioutil.ReadAll(c.Request.Body)*/

	rsp := utils.Response(errorcode.SUCCESS, "Success", "")

	c.JSON(200, utils.FormatRes(&rsp))
}
