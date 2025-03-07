package v1

import (
	"github.com/mxkacsa/v3-bitget-api-sdk/config"
	"github.com/mxkacsa/v3-bitget-api-sdk/pkg"
	"github.com/mxkacsa/v3-bitget-api-sdk/pkg/common"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init(conf config.Config) *SpotAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init(conf)
	return p
}

func (p *SpotAccountClient) Info() (string, error) {
	params := pkg.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/getInfo", params)
	return resp, err
}

func (p *SpotAccountClient) Assets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/assets-lite", params)
	return resp, err
}

func (p *SpotAccountClient) Bills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/bills", params)
	return resp, err
}

func (p *SpotAccountClient) TransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/transferRecords", params)
	return resp, err
}
