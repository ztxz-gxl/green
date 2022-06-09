package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/Green_life/db"
	"go_code/Green_life/eth"
	"net/http"
)

func Upload(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	pas := c.PostForm("pas")
	fmt.Println(name, password, pas)
	file, err := c.FormFile("upload")
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	id := db.RPUSH("WJLIST", []byte(fileName))

	fmt.Println("文件名：", fileName)
	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)
	if err := c.SaveUploadedFile(file, "./Green_life/public/"+fileName); err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	CompareName(name, c)
	read := Read(int(id.(int64)))
	addr := "0X" + read[12:52]
	marshal, err := json.Marshal(&BackUser{Name: name, Addr: addr, Passwd: password, FileID: int(id.(int64)), FilePsw: pas})
	if err != nil {
		fmt.Println(err)
	}
	ids := db.RPUSH("data7", marshal)
	//
	lange := db.LRANGE("data7", int(ids.(int64))-1, int(ids.(int64)))
	user := ChangeBackUser(lange)
	reads := Read(user[0].FileID)
	auth := eth.UserInit{Key: reads, Password: user[0].FilePsw}
	eth.InitName(name, auth.Auth())
	c.Redirect(http.StatusFound, "/index7")
}
