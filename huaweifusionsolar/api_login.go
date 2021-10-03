package huaweifusionsolar

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type loginRequest struct {
	UserName   string `json:"userName"`
	SystemCode string `json:"systemCode"`
}

func (h *HuaweiInverter) Login() error {
	b, _ := json.Marshal(loginRequest{
		UserName:   h.username,
		SystemCode: h.password,
	})
	req, _ := http.NewRequest("POST", h.baseAPI+"/login", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := ApiResponse{}
	err = json.Unmarshal(body, &data)

	if data.FailCode != 0 {
		log.Fatal("Login failure")
	}

	for _, val := range resp.Header.Values("Set-Cookie") {
		if parts := strings.Split(strings.Split(val, ";")[0], "="); parts[0] == "XSRF-TOKEN" {
			h.token = parts[1]
		}
	}

	return nil
}
