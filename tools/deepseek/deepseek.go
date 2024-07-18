package deepseek

import (
	global "Robit_Deepseek/tools/global"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Getreply(contentText string) (string, error) {
	var err error

	apiKey := global.Conf.Keys.Deepseek_api
	apiUrl := global.Conf.Keys.ApiUrl

	data := map[string]interface{}{
		"model": "deepseek-chat",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": contentText},
		},
		"stream": false,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		logrus.WithError(err).Error("json marshal error")
		return "", errors.Wrapf(err, "Error marshalling JSON")
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.WithError(err).Error("creating request error")
		return "", errors.Wrapf(err, "Error creating request")
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	logrus.WithFields(logrus.Fields{
		"url":     apiUrl,
		"headers": req.Header,
		"body":    string(jsonData),
	}).Debug("Sending API request")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.Wrapf(err, "Error sending request")
	}
	defer resp.Body.Close()

	// 处理响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "Error reading response body")
	}

	logrus.WithFields(logrus.Fields{
		"status_code": resp.StatusCode,
		"body":        string(body),
	}).Debug("Received API response")

	if resp.StatusCode != 200 {
		return "", errors.Errorf("API request failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", errors.Wrapf(err, "Error unmarshalling JSON")
	}

	// 提取content字段
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", errors.New("Invalid API response format: missing choices")
	}

	firstChoice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", errors.New("Invalid API response format: invalid choice format")
	}

	message, ok := firstChoice["message"].(map[string]interface{})
	if !ok {
		return "", errors.New("Invalid API response format: missing message")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", errors.New("Invalid API response format: missing content")
	}

	response := fmt.Sprintf("我是%s，很高兴为你解答\n%s", global.Conf.Keys.Bot_name, content)
	return response, nil
}