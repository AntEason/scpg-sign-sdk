package sign

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/ProEason/scpg-sign-sdk/encoding"
	"sort"
	"strings"
)
//头部参数
type HeaderData struct {
	appId string
	timestamp string
	orgCode string
}
//加密参数
type Scpg struct {
	headerData HeaderData
	methodType string  //请求类型
	privateKey string
	bodyParamMap map[string]interface{}  //post 请求参数
	pathParamMap map[string]interface{}  // get 请求参数
}

const (
	MethodType_POST string = "POST"
	MethodType_GET string = "GET"
)

func (this *Scpg) SetParamData(appId string,orgCode string,timestamp string,privateKey string,methodType string ,bodyParamMap map[string]interface{} ,pathParamMap map[string]interface{}  ){
		headerData:=HeaderData{appId:appId,orgCode:orgCode,timestamp:timestamp}
		this.headerData=headerData
		this.privateKey=privateKey
		this.methodType=methodType
		if bodyParamMap!=nil{
			this.bodyParamMap=bodyParamMap
		}
		if pathParamMap!=nil{
			this.pathParamMap=pathParamMap
		}
}
func (this *Scpg) GenerateSign() (string,error) {
	paramMap := make(map[string]string)
	headerData:=this.headerData
	paramMap["appId"]=headerData.appId
	paramMap["timestamp"]=headerData.timestamp
	paramMap["orgCode"]=headerData.orgCode
	if strings.EqualFold(this.methodType,MethodType_POST){
		b, _ := json.Marshal(this.bodyParamMap)
		paramMap["bizContent"]=string(b)
	}else if strings.EqualFold(this.methodType,MethodType_GET){
		for k,v:=range this.pathParamMap{
			paramMap[k]=v.(string)
		}
	}
	keys :=mapKeySort(paramMap)
	var buffer bytes.Buffer
    for i,v:=range keys{
		buffer.WriteString(v)
		buffer.WriteString("=")
		val, ok := paramMap[v]
		if(ok){
			buffer.WriteString(val)
		}
		if i != (len(keys) - 1){
			buffer.WriteString("&")
		}
	}

	fmt.Println("加密字符为："+buffer.String())
   pkey,err:= encoding.ParsePKCS8PrivateKey(this.privateKey)
   if pkey==nil{
   	return "",err
   }
	return encoding.Sign(buffer.String(),pkey.(*rsa.PrivateKey))
}

func mapKeySort(paramMap map[string]string)  []string{
	keys := make([]string, 0, len(paramMap))
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys;
}