package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	appId              string
	timestamp          string
	orgCode            string
	sign               string  // 签名
	bodyParamMap       map[string]interface{}  //post 请求参数
	pathParamMap       map[string]interface{}  // get 请求参数
	method             string
	url                string
}

func (this *Client) DoRequest() (err error) {
	client := &http.Client{}
	if strings.EqualFold(this.method,"POST"){
		b, _ := json.Marshal(this.bodyParamMap)
		reqsetBody :=strings.NewReader(string(b))
		req, err := http.NewRequest(this.method, this.url, reqsetBody)
		if err != nil {
			// handle error
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("appId", this.appId)
		req.Header.Set("timestamp", this.timestamp)
		req.Header.Set("sign", this.sign)
		req.Header.Set("orgCode", this.orgCode)
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
		return err
	}else if strings.EqualFold(this.method,"GET") {
		var buffer bytes.Buffer
		for k, v := range this.pathParamMap {
			buffer.WriteString(k)
			buffer.WriteString("=")
			buffer.WriteString(v.(string))
			buffer.WriteString("&")
		}
		query := strings.TrimRight(buffer.String(), "&")
		this.url = this.url + "?" + query
		req, err := http.NewRequest(this.method, this.url, nil)
		if err != nil {
			// handle error
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("appId", this.appId)
		req.Header.Set("timestamp", this.timestamp)
		req.Header.Set("sign", this.sign)
		req.Header.Set("orgCode", this.orgCode)
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
		return err
	}
	return  nil
}


