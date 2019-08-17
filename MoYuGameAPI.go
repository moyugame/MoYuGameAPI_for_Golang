package MoYuGameAPI

import (
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"unsafe"
)

type ServerInfo struct { //服务器相关信息类型
	ID       int
	status   int
	IP       string
	port     int
	version  string
	versionD string
	motd     string
	MODlist  string
	favicon  string
}
type BanInfo struct { //ban相关信息类型
	IsBan     bool
	BanCause  string
	BanDate   string
	toBanDate string
}

func aPI_get_I(i, r, q string) []byte { //发起API请求
	client := &http.Client{}

	req, err := http.NewRequest(i, "https://api.moyugame.com/hqmc_q/"+r, strings.NewReader(q))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println("body",string(body))
	defer resp.Body.Close()
	return body
}

func QueryserverIP(ip string) (SI ServerInfo) { //查询服务器信息
	body := aPI_get_I("GET", "request", "serverIP="+ip)
	sss, err := jsonparser.GetInt(body, "code")
	if err != nil {

	}
	if sss != 200 {
		return SI
	}
	port, err := jsonparser.GetInt(body, "data", "ID")
	idPointer := (*int)(unsafe.Pointer(&port))
	SI.port = *idPointer
	SI.ID = *idPointer
	SI.IP, err = jsonparser.GetString(body, "data", "IP")
	port, err = jsonparser.GetInt(body, "data", "port")
	idPointer = (*int)(unsafe.Pointer(&port))
	SI.port = *idPointer
	SI.version, err = jsonparser.GetString(body, "data", "version")
	SI.versionD, err = jsonparser.GetString(body, "data", "versionD")
	//fmt.Println(err)
	SI.motd, err = jsonparser.GetString(body, "data", "motd")
	SI.MODlist = "[]"
	SI.favicon, err = jsonparser.GetString(body, "data", "favicon")
	//fmt.Println(jsonparser.GetString(body,"data","motd"))
	//fmt.Println(SI.versionD)
	return SI
}

func QueryUserIsBan(id int, username string) (SI BanInfo) { //查询用户是否被封禁
	//fmt.Println("BAN查询",id,strconv.Itoa(id))
	body := aPI_get_I("GET", "QueryUserIsBan", "ServerId="+strconv.Itoa(id)+"&username="+username)
	sss, err := jsonparser.GetInt(body, "code")
	if err != nil {
		return SI
	}
	if sss != 200 {
		return SI
	}
	//fmt.Println(string(body))
	SI.IsBan, err = jsonparser.GetBoolean(body, "data", "isban")
	SI.BanCause, err = jsonparser.GetString(body, "data", "BanCause")
	SI.BanDate, err = jsonparser.GetString(body, "data", "BanDate")
	SI.toBanDate, err = jsonparser.GetString(body, "data", "toBanDate")
	return SI
}
