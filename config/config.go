package config

import "github.com/mxkacsa/v3-bitget-api-sdk/constants"

const (
	baseUrl = "https://api.bitget.com"
	wsUrl   = "wss://ws.bitget.com/mix/v1/stream"

	timeoutSecond = 30
	signType      = constants.SHA256
)

type Config struct {
	BaseUrl       string
	WsUrl         string
	ApiKey        string
	SecretKey     string
	PASSPHRASE    string
	TimeoutSecond int
	SignType      string
}

func DefaultConfig() Config {
	return Config{
		BaseUrl:       baseUrl,
		WsUrl:         wsUrl,
		ApiKey:        "",
		SecretKey:     "",
		PASSPHRASE:    "",
		TimeoutSecond: timeoutSecond,
		SignType:      signType,
	}
}
