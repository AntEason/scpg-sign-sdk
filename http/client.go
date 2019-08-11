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
	AppId        string
	Timestamp    string
	OrgCode      string
	Sign         string                 // 签名
	BodyParamMap map[string]interface{} //post 请求参数
	PathParamMap map[string]interface{} // get 请求参数
	Method       string
	Url          string
}

func (this *Client) DoRequest() (err error) {
	client := &http.Client{}
	if strings.EqualFold(this.Method, "POST") {
		b, _ := json.Marshal(this.BodyParamMap)
		reqsetBody := strings.NewReader(string(b))
		req, err := http.NewRequest(this.Method, this.Url, reqsetBody)
		if err != nil {
			// handle error
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("appId", this.AppId)
		req.Header.Set("timestamp", this.Timestamp)
		req.Header.Set("sign", this.Sign)
		req.Header.Set("orgCode", this.OrgCode)
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
		return err
	} else if strings.EqualFold(this.Method, "GET") {
		var buffer bytes.Buffer
		for k, v := range this.PathParamMap {
			buffer.WriteString(k)
			buffer.WriteString("=")
			buffer.WriteString(v.(string))
			buffer.WriteString("&")
		}
		query := strings.TrimRight(buffer.String(), "&")
		this.Url = this.Url + "?" + query
		req, err := http.NewRequest(this.Method, this.Url, nil)
		if err != nil {
			// handle error
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("appId", this.AppId)
		req.Header.Set("timestamp", this.Timestamp)
		req.Header.Set("sign", this.Sign)
		req.Header.Set("orgCode", this.OrgCode)
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
		return err
	}
	return nil
}
