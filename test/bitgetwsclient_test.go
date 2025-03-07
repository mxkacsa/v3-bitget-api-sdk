package test

import (
	"fmt"
	"github.com/mxkacsa/v3-bitget-api-sdk/config"
	"github.com/mxkacsa/v3-bitget-api-sdk/pkg/client/ws"
	"github.com/mxkacsa/v3-bitget-api-sdk/pkg/model"
	"testing"
)

func TestBitgetWsClient_New(t *testing.T) {

	client := new(ws.BitgetWsClient).Init(config.DefaultConfig(), true, func(message string) {
		fmt.Println("default error:" + message)
	}, func(message string) {
		fmt.Println("default error:" + message)
	})

	var channelsDef []model.SubscribeReq
	subReqDef1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channelsDef = append(channelsDef, subReqDef1)
	client.SubscribeDef(channelsDef)

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		fmt.Println("appoint:" + message)
	})
	client.Connect()

}
