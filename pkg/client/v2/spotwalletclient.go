package v2

import (
	"github.com/mxkacsa/v3-bitget-api-sdk/pkg"
	"github.com/mxkacsa/v3-bitget-api-sdk/pkg/common"
)

type SpotWalletApi struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotWalletApi) Transfer(params map[string]string) (string, error) {
	postBody, jsonErr := pkg.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/transfer", postBody)
	return resp, err
}

func (p *SpotWalletApi) DepositAddress(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-address", params)
	return resp, err
}

func (p *SpotWalletApi) Withdrawal(params map[string]string) (string, error) {
	postBody, jsonErr := pkg.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/withdrawal", postBody)
	return resp, err
}

func (p *SpotWalletApi) WithdrawalRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/withdrawal-records", params)
	return resp, err
}

func (p *SpotWalletApi) DepositRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-records", params)
	return resp, err
}
