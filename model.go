package sdk_evm_proxy_wallet

type WalletDTO struct {
	ID        uint   `json:"id"`
	Address   string `json:"address"`
	PublicKey string `json:"public_key,omitempty"`
}

type TxDTO struct {
	ID          uint   `json:"id"`
	NetworkName string `json:"network_name"`
	Hash        string `json:"hash"`
	Type        string `json:"type"`
	From        string `json:"from"`
	To          string `json:"to"`
	Nonce       uint64 `json:"nonce"`
	GasLimit    uint   `json:"gas_limit"`
	GasPrice    uint   `json:"gas_price"`
	GasFee      uint   `json:"gas_fee"`
	Status      string `json:"status"`
}

type CreateWalletResponseDTO struct {
	Address string `json:"address"`
	Secret  string `json:"secret"`
}

type GetBalanceResponseDTO struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type TransactOpt struct {
	Address      string `json:"address" binding:"required"`
	NetworkName  string `json:"network_name" binding:"required"`
	Rpc          string `json:"rpc" binding:"required"`
	GasLimit     uint   `json:"gas_limit" binding:"required"`
	GasPrice     uint   `json:"gas_price" binding:"required"`
	Value        uint   `json:"value"`
	ClientSecret string `json:"client_secret" binding:"required"`
}

type transferRequest struct {
	To          string      `json:"to"`
	TransactOpt TransactOpt `json:"transact_opt"`
}

type encryptDataRequest struct {
	PublicKey string `json:"public_key"`
	Data      string `json:"data"`
}

type decryptDataRequest struct {
	Data         string `json:"data"`
	Address      string `json:"address"`
	ClientSecret string `json:"client_secret"`
}

type deployContractRequest struct {
	Abi         string      `json:"abi" binding:"required"`
	Bin         string      `json:"bin" binding:"required"`
	Args        []string    `json:"args"`
	TransactOpt TransactOpt `json:"transact_opt"`
}

type writeContractRequest struct {
	Abi         string      `json:"abi" binding:"required"`
	Address     string      `json:"address" binding:"required"`
	Method      string      `json:"method" binding:"required"`
	Args        []string    `json:"args"`
	TransactOpt TransactOpt `json:"transact_opt"`
}

type readContractRequest struct {
	Abi     string   `json:"abi" binding:"required"`
	Address string   `json:"address" binding:"required"`
	Method  string   `json:"method" binding:"required"`
	Args    []string `json:"args"`
	Rpc     string   `json:"rpc" binding:"required"`
}
