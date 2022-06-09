package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go_code/Green_life/router"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mySession", store))

	r.GET("/", router.Main)
	r.GET("/getUser", router.GetUser)
	r.GET("/all", router.All)
	r.GET("/getData", router.GetData)
	r.GET("/getBankData", router.GetBankData)
	r.GET("/login", router.Login)
	r.GET("/bankSub", router.BankSub)
	r.GET("/getChart", router.GetChart)
	r.GET("/getReport", router.GetReport)
	r.GET("/load", router.Load)
	r.GET("/fy", router.FY)
	r.GET("/rep", router.Rep)
	r.GET("/getRep", router.GetRep)
	r.GET("/getRepAll", router.GetRepAll)
	r.GET("/sendRep", router.SendRep)
	r.GET("/voting", router.Voting)
	r.GET("/getInterest", router.GetInterest)

	//geth
	r.GET("/SQ", router.SQ)
	r.GET("/ZJSQ", router.ZJSQ)
	r.GET("/JSSQ", router.JSSQ)
	r.GET("/DH", router.DH)
	r.GET("/SQZZ", router.SQZZ)
	r.GET("/ZZ", router.ZZ)
	r.GET("/HBInfo", router.HBInfo)
	r.GET("/HBLogo", router.HBLogo)
	r.GET("/HBAmount", router.HBAmount)
	r.GET("/CXSQ", router.CXSQ)

	//back
	r.GET("/QCBank", router.QCBank)
	r.GET("/CRBank", router.CRBank)

	r.POST("/upload", router.Upload)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("run is err", err)
	}
}
