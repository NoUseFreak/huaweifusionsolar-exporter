package huaweifusionsolar

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *HuaweiInverter) queryURL(path string, query interface{}, data interface{}) error {
	b, err := json.Marshal(query)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"json": string(b),
	}).Debug("performing request: " + path)
	req, err := http.NewRequest("POST", h.baseAPI+path, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("XSRF-TOKEN", h.token)

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debug("ApiResponse: " + string(body))
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"json": string(body),
		}).Warn("Failed to parse response - " + err.Error())
	}
	return err

}
