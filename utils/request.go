//Package httpclient 一个简单的http客户端，用于请求第三方接口，简单实现了GET和POST方法，比较粗糙，一般情况下够用
package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/* type Response struct {
	*http.Response
	request     *http.Request
	requestBody []byte
	cookies     map[string]string
}

//Post domain string "https://aaa.com"
//Post data url.values, eg:
//data := url.Values{}
//data.Add("Name", "Ewan")
//data.Add("Age", 12)
func Post(domain string, data url.Values) (*Response, error) {
	//增加header选项
	r, _ := http.NewRequest("POST", domain, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return response(resp)
}

func Get(domain string, data url.Values) (*Response, error) {
	uri, _ := url.Parse(domain)
	uri.RawQuery = data.Encode()
	//r, _ := http.NewRequest("GET", domain + data.Encode(), nil)
	r, _ := http.NewRequest("GET", uri.String(), nil)
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return response(resp)
}

func response(resp *http.Response) (*Response, error) {
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("http error,http code: %d,body:%s", resp.StatusCode, body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
} */

type Client struct {
	header map[string]string
}

func (c *Client) Get(url string, data interface{}) error {
	return c.DoRequest("GET", url, "", data)
}

func (c *Client) Post(url string, params string, data interface{}) error {
	return c.DoRequest("POST", url, params, data)
}

func (c *Client) DoRequest(method, url, params string, data interface{}) (err error) {
	var req *http.Request
	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	if method == "GET" {
		if req, err = http.NewRequest(method, url, bytes.NewBuffer(nil)); err != nil {
			return err
		}
	} else {
		paramBytes := []byte(params)
		if req, err = http.NewRequest(method, url, bytes.NewReader(paramBytes)); err != nil {
			return err
		}
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, data); err != nil {
		return err
	}
	return nil
}
