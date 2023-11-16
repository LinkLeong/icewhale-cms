package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/model"
)

type VersionService struct{}

func (e *VersionService) GetVersion() (string, error) {
	url := global.GlobalConfig.Url + "/v1/sys/versions?user=icewhale"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsInBhc3N3b3JkIjoiIiwiZXhwIjoxNjYwNzMzNjg4LCJpc3MiOiJnaW4tYmxvZyJ9.AYzSmJg1fMhGVL3uG6T-zJ7rPeIVodtBAw_cHbbargs")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
func (e *VersionService) DeleteVersion(id string) (string, error) {
	url := global.GlobalConfig.Url + "/v1/sys/version/" + id + "?user=icewhale"
	method := "DELETE"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsInBhc3N3b3JkIjoiIiwiZXhwIjoxNjYwNzMzNjg4LCJpc3MiOiJnaW4tYmxvZyJ9.AYzSmJg1fMhGVL3uG6T-zJ7rPeIVodtBAw_cHbbargs")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
func (e *VersionService) AddVersion(version model.Version) (string, error) {
	url := global.GlobalConfig.Url + "/v1/sys/version?user=icewhale"
	method := "POST"
	configdata, _ := json.Marshal(version)

	body := bytes.NewBuffer([]byte(configdata))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsInBhc3N3b3JkIjoiIiwiZXhwIjoxNjYwNzMzNjg4LCJpc3MiOiJnaW4tYmxvZyJ9.AYzSmJg1fMhGVL3uG6T-zJ7rPeIVodtBAw_cHbbargs")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	bodyRes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(bodyRes), nil
}
