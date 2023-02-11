package sdk_evm_proxy_wallet

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/WebXense/ginger/ginger"
	"github.com/WebXense/http"
)

type evmProxyWalletSdk struct {
	host string
}

func NewEvmProxyWalletSdk(host string) *evmProxyWalletSdk {
	return &evmProxyWalletSdk{
		host: host,
	}
}

func (s *evmProxyWalletSdk) CreateWallet() (*CreateWalletResponseDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_create_wallet, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &CreateWalletResponseDTO{}), resp, nil
}

func (s *evmProxyWalletSdk) GetWallet(address string) (*WalletDTO, *ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](
		s.host+strings.ReplaceAll(url_get_wallet, ":address", address), nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &WalletDTO{}), resp, nil
}

func (s *evmProxyWalletSdk) ListWallet(page, limit int) ([]WalletDTO, *ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](
		s.host+url_list_wallet+"?page="+strconv.Itoa(page)+"&limit="+strconv.Itoa(limit), nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	data := make([]WalletDTO, 0)
	for _, v := range resp.Data.([]any) {
		data = append(data, *mapByJson(v, &WalletDTO{}))
	}
	return data, resp, nil
}

func (s *evmProxyWalletSdk) DeleteWallet(address, secret string) (*ginger.Response, error) {
	status, resp, err := http.Delete[ginger.Response](
		s.host+strings.ReplaceAll(url_get_wallet, ":address", address), nil, nil, map[string]any{
			"client_secret": secret,
		})
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	return resp, nil
}

func (s *evmProxyWalletSdk) GetWalletBalance(address, rpcUrl string) (*GetBalanceResponseDTO, *ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](
		s.host+url_get_balance+"?address="+address+"&rpc="+rpcUrl, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &GetBalanceResponseDTO{}), resp, nil
}

func (s *evmProxyWalletSdk) TransferBalance(toAddress string, opt TransactOpt) (*TxDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](
		s.host+url_transfer_balance, nil, nil, transferRequest{
			To:          toAddress,
			TransactOpt: opt,
		})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &TxDTO{}), resp, nil
}

func (s *evmProxyWalletSdk) EncryptData(publicKey string, data string) (string, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](
		s.host+url_encrypt_data, nil, nil, encryptDataRequest{
			PublicKey: publicKey,
			Data:      data,
		})
	if err != nil {
		return "", nil, err
	}
	if status != 200 {
		return "", nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return "", resp, nil
	}
	return resp.Data.(string), resp, nil
}

func (s *evmProxyWalletSdk) DecryptData(data, address, secret string) (string, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](
		s.host+url_decrypt_data, nil, nil, decryptDataRequest{
			Data:         data,
			Address:      address,
			ClientSecret: secret,
		})
	if err != nil {
		return "", nil, err
	}
	if status != 200 {
		return "", nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return "", resp, nil
	}
	return resp.Data.(string), resp, nil
}

func (s *evmProxyWalletSdk) GetTx(txHash, rpcUrl string) (*TxDTO, *ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](
		s.host+strings.ReplaceAll(url_get_tx, ":hash", txHash)+"?rpc="+rpcUrl, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &TxDTO{}), resp, nil
}

func (s *evmProxyWalletSdk) DeployContract(abi, bin string, args []string, opt TransactOpt) (*TxDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](
		s.host+url_deploy_contract, nil, nil, deployContractRequest{
			Abi:         abi,
			Bin:         bin,
			Args:        args,
			TransactOpt: opt,
		})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &TxDTO{}), resp, nil
}

func (s *evmProxyWalletSdk) WriteContract(address, method, abi string, args []string, opt TransactOpt) (*TxDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](
		s.host+url_write_contract, nil, nil, writeContractRequest{
			Abi:         abi,
			Address:     address,
			Method:      method,
			Args:        args,
			TransactOpt: opt,
		})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &TxDTO{}), resp, nil
}

func (s *evmProxyWalletSdk) ReadContract(address, method, abi string, args []string, rpcUrl string) ([]any, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](
		s.host+url_read_contract, nil, nil, readContractRequest{
			Abi:     abi,
			Address: address,
			Method:  method,
			Args:    args,
			Rpc:     rpcUrl,
		})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return resp.Data.([]any), resp, nil
}

func mapByJson[T any](from interface{}, to *T) *T {
	j, err := json.Marshal(from)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(j, to)
	if err != nil {
		return nil
	}
	return to
}
