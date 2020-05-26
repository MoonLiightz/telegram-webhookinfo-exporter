package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/moonliightz/telegram-webhookinfo-exporter/internal/config"
	"github.com/moonliightz/telegram-webhookinfo-exporter/internal/model"
)

// LoadWebhookInfo calls the telegram bot api getWebhookInfo
func LoadWebhookInfo(config *config.Config) (*model.WebhookInfo, error) {
	res, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getWebhookInfo", config.Telegram.Token))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var winfo model.WebhookInfo
	err = json.Unmarshal(body, &winfo)
	if err != nil {
		fmt.Println(string(body))
		return nil, err
	}

	return &winfo, nil
}
