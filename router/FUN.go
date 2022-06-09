package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_code/Green_life/db"
	"go_code/Green_life/eth"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type List struct {
	Code int
	Text string
	Type string
}

type BackData struct {
	Id   int
	Addr string
	Date string
	Fun  string
	Stu  string
}

type BackRep struct {
	Id   int
	Addr string
	Name string
	Date string
}

type RepList struct {
	Name       string
	VotingTime string
}

type BackUser struct {
	Addr    string
	Name    string
	Passwd  string
	FileID  int
	FilePsw string
}

type StuSlice []List

func (ss StuSlice) Len() int {
	return len(ss)
}

func (ss StuSlice) Less(i, j int) bool {
	return ss[i].Type > ss[j].Type
}

func (ss StuSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i] //与上面的交换方法一致
}

func ChangeUser(ip string, b []eth.PerList) []eth.PerList {
	var u []eth.PerList
	for _, v := range b {
		if strings.ToLower(v.Ip.String()) == strings.ToLower(ip) {
			u = append(u, v)
		}
	}
	return u
}

func ChangeBackData(b []eth.PerList) []eth.PerList {
	var us []eth.PerList
	if len(b) >= 5 {
		for i := 0; i < 5; i++ {
			us = append(us, b[i])
		}
	} else {
		for _, v := range b {
			us = append(us, v)
		}
	}

	return us
}

func ChangeRepList(b []eth.VoteList) []RepList {
	var u []RepList
	var us []RepList
	lange := db.LRANGE("data7", 0, -1)
	user := ChangeBackUser(lange)
	for _, v1 := range b {
		for _, v := range user {
			if strings.ToLower(v.Addr) == strings.ToLower(v1.Ip.String()) {
				u = append(u, RepList{Name: v.Name, VotingTime: v1.LimitTime.String()})
			}
		}
	}
	fmt.Println(u)
	if len(u) >= 5 {
		for i := 0; i < 5; i++ {
			us = append(us, u[i])
		}
	} else {
		for _, v := range u {
			us = append(us, v)
		}
	}
	fmt.Println(us)
	return us
}

func ChangeRepListAll(b []eth.VoteList) []RepList {
	var u []RepList
	lange := db.LRANGE("data7", 0, -1)
	user := ChangeBackUser(lange)
	for _, v1 := range b {
		for _, v := range user {
			if strings.ToLower(v.Addr) == strings.ToLower(v1.Ip.String()) {
				u = append(u, RepList{Name: v.Name, VotingTime: v1.LimitTime.String()})
			}
		}
	}
	return u
}

func ChangeList(b []string) []List {
	var u []List
	for _, v := range b {
		var us List
		err := json.Unmarshal([]byte(v), &us)
		if err != nil {
			fmt.Println("ChangeList err:", err)
		}
		u = append(u, us)
	}
	return u
}

func ChangeNUM() StuSlice {
	lange := db.LRANGE("data7", 0, -1)
	user := ChangeBackUser(lange)
	GetNum := eth.GETNUM()
	var list StuSlice
	j := 1
	for _, v := range GetNum {
		s := 0
		for i, v1 := range list {
			for _, v2 := range user {
				if strings.ToLower(v2.Addr) == strings.ToLower(v.Ip.String()) {
					if v1.Text == v2.Name {
						atoi, _ := strconv.Atoi(v1.Type)
						a := atoi + 1
						list[i].Type = strconv.Itoa(a)
						s++
						break
					}
				}
			}
		}
		if s == 0 {
			for _, v2 := range user {
				if strings.ToLower(v2.Addr) == strings.ToLower(v.Ip.String()) {
					list = append(list, List{Code: j, Text: v2.Name, Type: "1"})
					j++
				}
			}
		}
	}
	sort.Sort(list)
	return list
}

func ChangeNUMOFFUN() StuSlice {
	GetNum := eth.GETNUM()
	var lists StuSlice
	j := 1
	for _, v := range GetNum {
		s := 0
		for i, v1 := range lists {
			if v1.Text == v.Fun[0:12] {
				atoi, _ := strconv.Atoi(v1.Type)
				a := atoi + 1
				lists[i].Type = strconv.Itoa(a)
				s++
				break
			}
		}
		if s == 0 {
			lists = append(lists, List{Code: j, Text: v.Fun[0:12], Type: "1"})
			j++
		}
	}
	sort.Sort(lists)
	return lists
}

func ChangeListSort(b []string) []List {
	var u []List
	for _, v := range b {
		var us List
		err := json.Unmarshal([]byte(v), &us)
		if err != nil {
			fmt.Println("ChangeListSort err:", err)
		}
		u = append(u, us)
	}
	for i := 0; i < len(u)-1; i++ {
		for j := 0; j < len(u)-1-i; j++ {
			t1, _ := strconv.Atoi(u[j].Type)
			t2, _ := strconv.Atoi(u[j+1].Type)
			if t1 > t2 {
				u[j], u[j+1] = u[j+1], u[j]
			}
		}
	}
	return u
}

func SortUser(ip string, b []eth.PerList) []eth.PerList {
	var u []eth.PerList
	var us []eth.PerList
	for _, v := range b {
		if strings.ToLower(v.Ip.String()) == strings.ToLower(ip) {
			u = append(u, v)
		}
	}
	if len(u) >= 5 {
		for i := 0; i < 5; i++ {
			us = append(us, u[i])
		}
	} else {
		for _, v := range u {
			us = append(us, v)
		}
	}

	return us
}

func ChangeBackUser(b []string) []BackUser {
	var u []BackUser
	for _, v := range b {
		var us BackUser
		err := json.Unmarshal([]byte(v), &us)
		if err != nil {
			fmt.Println("ChangeBackUser err:", err)
		}
		u = append(u, us)
	}
	return u
}

func LV(year int) int {
	a := db.LRANGE("data4", 0, -1)
	listData4 := ChangeList(a)
	if year < 1 {
		floats, _ := strconv.ParseFloat(listData4[0].Text, 32)
		return int(floats * 100)
	} else if year >= 1 && year < 3 {
		floats, _ := strconv.ParseFloat(listData4[0].Text, 32)
		return int(floats * 100)
	} else if year >= 3 && year < 5 {
		floats, _ := strconv.ParseFloat(listData4[1].Text, 32)
		return int(floats * 100)
	} else if year >= 5 && year < 10 {
		floats, _ := strconv.ParseFloat(listData4[2].Text, 32)
		return int(floats * 100)
	} else if year >= 10 && year < 20 {
		floats, _ := strconv.ParseFloat(listData4[3].Text, 32)
		return int(floats * 100)
	} else {
		floats, _ := strconv.ParseFloat(listData4[4].Text, 32)
		return int(floats * 100)
	}
}

func FTLogin(name string, passwd string) (int, bool) {
	b := db.LRANGE("data7", 0, -1)
	var u []BackUser
	for _, v := range b {
		var us BackUser
		err := json.Unmarshal([]byte(v), &us)
		if err != nil {
			fmt.Println("ChangeListSort err:", err)
		}
		u = append(u, us)
	}
	for i, v := range u {
		if v.Name == name && v.Passwd == passwd {
			return i + 1, true
		}
	}
	return 0, false
}

func Read(id int) string {
	lrange := db.LRANGE("WJLIST", id-1, id)
	file, err := os.Open("Green_life\\public\\" + lrange[0])
	if err != nil {
		fmt.Println("file", err)
		return ""
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("fileinfo", err)
		return ""
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("bytesread", err)
		return ""
	}
	return string(buffer)
}

func SETSession(c *gin.Context, id int) {
	session := sessions.Default(c)
	session.Set("id", id)
	err := session.Save()
	if err != nil {
		fmt.Println("SETSession", err)
	}
}

func GetSession(c *gin.Context) int {
	session := sessions.Default(c)
	get := session.Get("id")
	fmt.Println(get)
	if get == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "登录信息有误",
		})
	}
	return get.(int)
}

func CompareName(name string, c *gin.Context) {
	lange := db.LRANGE("data7", 0, -1)
	user := ChangeBackUser(lange)
	for _, v := range user {
		if v.Name == name {
			c.JSON(http.StatusBadRequest, gin.H{
				"massage": "用户名被重复使用",
			})
		}
	}
}
