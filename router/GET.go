package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/Green_life/db"
	"go_code/Green_life/eth"
	"net/http"
	"strconv"
)

//WeChat
func GetUser(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	personal := eth.GETNewPersonal(user[0].Addr)
	b := eth.GETNUM()
	change := SortUser(user[0].Addr, b)
	c.JSON(http.StatusOK, gin.H{
		"Addr": user[0].Addr,
		"user": personal,
		"list": change,
	})
}

func All(c *gin.Context) {
	//id := c.Query("id")
	id := GetSession(c)
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	b := eth.GETNUM()
	change := ChangeUser(user[0].Addr, b)
	c.JSON(http.StatusOK, gin.H{
		"list": change,
	})
}

//排行榜=============================================================
func GetData(c *gin.Context) {
	listData1 := ChangeNUM()
	d := db.LRANGE("data3", 0, -1)
	listData2 := ChangeNUMOFFUN()
	listData3 := ChangeListSort(d)
	c.JSON(http.StatusOK, gin.H{
		"listData2": listData2,
		"listData3": listData3,
		"a":         listData1,
	})
}

//===================================================================

//利率
func GetBankData(c *gin.Context) {
	a := db.LRANGE("data4", 0, -1)
	listData4 := ChangeList(a)
	c.JSON(http.StatusOK, gin.H{
		"listData4": listData4,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	pass := c.Query("password")
	id, T := FTLogin(username, pass)
	SETSession(c, id)
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	stu := eth.InitNameStu(user[0].Addr)
	//fmt.Println(T)
	//fmt.Println(stu)
	if T && stu {
		c.JSON(http.StatusOK, gin.H{
			"id": GetSession(c),
		})
	} else {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"id":  GetSession(c),
			"err": "信息有误",
		})
	}
}

func BankSub(c *gin.Context) {
	year, _ := strconv.Atoi(c.Query("year"))
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	amount, _ := strconv.Atoi(c.Query("amount"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	lv := LV(year)
	bank, s := eth.Bank(amount, year, lv, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage":  "成功存入",
		"interest": lv,
		"hash":     bank,
		"timeHash": s,
	})
}

func Load(c *gin.Context) {
	num, _ := strconv.Atoi(c.Query("num"))
	company := c.Query("company")
	dw := c.Query("dw")
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	ss := company + strconv.Itoa(num) + dw
	fmt.Println(ss)
	part := eth.TakePart(ss, num, num, auth.Auth())
	fmt.Println(num, id, company, dw)
	c.JSON(http.StatusOK, gin.H{
		"company": company,
		"num":     num,
		"dw":      dw,
		"EN":      num,
		"massage": part,
	})
}

//nuxt
func Main(c *gin.Context) {
	b := eth.GETNUM()
	b = ChangeBackData(b)
	c.JSON(http.StatusOK, gin.H{
		"all":     eth.GetYearData().CurrentQuota,
		"allUser": eth.GETUserNum(),
		"join":    eth.NUM(),
		"report":  eth.GetRVNUM(),
		"list":    b,
	})
}

//////////////////////////////////////////
func Rep(c *gin.Context) {
	//b := db.LRANGE("rep", 0, -1)
	voting := eth.GetVoting()
	a := ChangeRepList(voting)
	c.JSON(http.StatusOK, gin.H{
		"list": a,
	})
}

/////////////////////////////////////////////////////
func SendRep(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	addr := c.Query("addr")
	reason := c.Query("reason")
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.Report(addr, reason, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "举报成功,后续将会在数据中心进行投票",
	})
}

//////////////////////////////////////////////
func GetRep(c *gin.Context) {
	voting := eth.GetVoting()
	a := ChangeRepList(voting)
	fmt.Println(a)
	c.JSON(http.StatusOK, gin.H{
		"list": a,
	})
}

///////////////////////////////////////////////////////
func GetRepAll(c *gin.Context) {
	voting := eth.GetVoting()
	a := ChangeRepListAll(voting)
	c.JSON(http.StatusOK, gin.H{
		"list": a,
	})
}

/////////////////////////////////////////////////////////////////
func Voting(c *gin.Context) {
	id := GetSession(c)
	Id, _ := strconv.Atoi(c.Query("Id"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.Voting(Id, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "投票成功",
	})
}

func GetInterest(c *gin.Context) {
	id := GetSession(c)
	hash := c.Query("hash")
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.GetInterest(hash, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "取出成功",
	})
}

func GetChart(c *gin.Context) {
	s := eth.GETYear().String()
	data := eth.GetYearData()
	c.JSON(http.StatusOK, gin.H{
		"bDate": []string{"2003", "2004", "2005", "2006", "2007", "2008", "2009", "2010", "2011", "2012", "2013", s},
		"a":     []int{250, 150, 120, 320, 480, 350, 150},
		"b":     []int{20, 25, 40, 30, 45, 40, 55, 40, 48, 40, 42, int(data.PeopleNUM.Int64())},
	})
}

func GetReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"bDate": []string{"2003", "2004", "2005", "2006", "2007", "2008", "2009", "2010", "2011", "2012", "2013", "2014"},
		"a":     []int{2, 0, 1, 0, 0, 0, 3},
		"b":     []int{2, 2, 4, 3, 4, 4, 5, 4, 4, 4, 4, 5},
	})
}

func FY(c *gin.Context) {
	b := eth.GETNUM()
	c.JSON(http.StatusOK, gin.H{
		"list": b,
	})
}

//geth
func SQ(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	addr := c.Query("addr")
	amount, _ := strconv.Atoi(c.Query("amount"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.Approve(addr, amount, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "授权成功",
	})
}

func ZJSQ(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	addr := c.Query("addr")
	amount, _ := strconv.Atoi(c.Query("amount"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.INApprove(addr, amount, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "增加授权成功",
	})
}

func JSSQ(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	addr := c.Query("addr")
	amount, _ := strconv.Atoi(c.Query("amount"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.DEApprove(addr, amount, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "减少授权成功",
	})
}

func DH(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	amount, _ := strconv.Atoi(c.Query("amount"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.Exchange(amount, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "兑换成功",
	})
}

func SQZZ(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	amount, _ := strconv.Atoi(c.Query("amount"))
	addr := c.Query("addr")
	spender := c.Query("spender")
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.TransferForm(spender, addr, amount, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "授权转账成功",
	})
}

func ZZ(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	amount, _ := strconv.Atoi(c.Query("amount"))
	addr := c.Query("addr")
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.Transfer(addr, amount, auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "转账成功",
	})
}

func HBInfo(c *gin.Context) {
	info := eth.ERCInfo()
	c.JSON(http.StatusOK, gin.H{
		"massage": "货币信息:" + info,
	})
}

func HBLogo(c *gin.Context) {
	logo := eth.ERCLogo()
	c.JSON(http.StatusOK, gin.H{
		"massage": "货币标志:" + logo,
	})
}

func HBAmount(c *gin.Context) {
	amount := eth.ERCAmount()
	c.JSON(http.StatusOK, gin.H{
		"massage": "发币数量:" + amount.String(),
	})
}

func CXSQ(c *gin.Context) {
	SQZ := c.Query("SQZ")
	BSQ := c.Query("BSQ")
	approve := eth.SelectApprove(SQZ, BSQ)
	c.JSON(http.StatusOK, gin.H{
		"massage": approve,
	})
}

//back
func QCBank(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	amount, _ := strconv.Atoi(c.Query("num"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.TakeOut(amount, "取出", auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "取出成功",
	})
}

func CRBank(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	id := GetSession(c)
	amount, _ := strconv.Atoi(c.Query("num"))
	lange := db.LRANGE("data7", id-1, id)
	user := ChangeBackUser(lange)
	read := Read(user[0].FileID)
	auth := eth.UserInit{Key: read, Password: user[0].FilePsw}
	eth.Deposit(amount, "存入", auth.Auth())
	c.JSON(http.StatusOK, gin.H{
		"massage": "存入成功",
	})
}
