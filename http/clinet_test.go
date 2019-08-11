package http

import (
	"fmt"
	"github.com/ProEason/scpg-sign-sdk/sign"
	"strconv"
	"testing"
	"time"
)

const (
	PRIVATE_KEY  string= "MIICdAIBADANBgkqhkiG9w0BAQEFAASCAl4wggJaAgEAAoGBAJAIPomOrxlXeZmLZaupHE7yLm6nCPPErLgZmP5h9kh7+GML8hNVeOqPT1nYqH5Nrv0lrvNAFi+W07IxSyXimULy2n/LGBn80LRysV87colAFGXiPZhK7ic2E5kCQX++4C+6xde7qogmnlG3j00p7Y5b/xZh0t9WnCYxDbziymLLAgMBAAECfxXxK5XJJuO1b9D0a4i7OpTMUEaLK9H3JFjnjWczhVGFkVGFgFtxqfoR2wTU/NZA/2eRVhW4raiSUa3T1J7w5mgR5PnR5uszlskZDRRL11GMWZfYzsQtX9Ha+16mzs/fm20WZI8ca2pFLgAIKPP5/xrBO5sISLyNkcAYTJPh3RECQQDQHsuduWXnxI2BzESXn2vUJklbJHGGYpo4YrSQb3x+wiPRB8ERa14D5iY45NhXR8PW1fWWTkl37Xq7b2wDCVmjAkEAsSr+Qd0i/gWsR0z3ihdI0JFJLyJBH4UBoxhwskuwAkoHWoVU/7ht0Au6hC3WY2OeNC1TkSo0ejaGR1hCjha0uQJAZqbZ3ajld6S9+0iKiJLMd66OvhLCn/sggDFHSHqE2Gzfh6mgc/wWOiwZrDRL9PjjwqJz6LgiirGquKqjhvfkqwJBAIc5HmldBIlIAS9GkqVGkmAGlAoypewN06sNS276j/OfVy+DFtY7iS1C1nBun/SqIyUF94OT6Avjs+eeeOYEdJkCQEGE/Seoz0AIFl5QOXwBatuTtZuUcwcyLTqFLXtZxwuloAEvObYarkhW3BtWoHx/2m4k4tPF/LJtyNzsx1zcGlA="
	APP_ID string= "f82b8253a7ff79de994a0c2fe505fac6"
	ORG_CODE string = "G001Z003C0002"
)

func TestClient_DoRequest(t *testing.T) {
	//post
	var param =new (sign.Scpg)
	var time =strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
	bodyParamMap:=make(map[string]interface{})
	bodyParamMap["transId"]=time
	bodyParamMap["remark"]=""
	bodyParamMap["terminal"]=""
	bodyParamMap["event"]="CONSUME"
	bodyParamMap["point"]=100
	bodyParamMap["memberId"]="134680196342081156"
	param.SetParamData(APP_ID,ORG_CODE,time,PRIVATE_KEY,sign.MethodType_POST,bodyParamMap,nil)
	text,_:=param.GenerateSign()
	client:=Client{sign:text,appId:APP_ID,url:"http://scrm-uat.scpgroup.com.cn/openapi-uat/api/v1/point/reduce",method:sign.MethodType_POST,timestamp:time,bodyParamMap:bodyParamMap,orgCode:ORG_CODE}
	err:=client.DoRequest()
	fmt.Println(err)
	//get
	//pathParamMap:=make(map[string]interface{})
	//pathParamMap["accountType"]="1"
	//pathParamMap["accountValue"]="66511"
	//param.SetParamData(APP_ID,ORG_CODE,time,PRIVATE_KEY,sign.MethodType_GET,nil,pathParamMap)
	//text,_:=param.GenerateSign()
	//client:=Client{sign:text,appId:APP_ID,url:"http://scrm-uat.scpgroup.com.cn/openapi-uat/api/v1/member",method:sign.MethodType_GET,timestamp:time,pathParamMap:pathParamMap,orgCode:ORG_CODE}
	//err:=client.DoRequest()
	//fmt.Println(err)
	}