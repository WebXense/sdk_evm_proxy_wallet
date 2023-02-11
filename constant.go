package sdk_evm_proxy_wallet

const (
	url_create_wallet    = "/wallet"
	url_get_wallet       = "/wallet/:address"
	url_list_wallet      = "/wallet"
	url_get_balance      = "/wallet/balance"
	url_transfer_balance = "/wallet/transfer"
	url_encrypt_data     = "/wallet/encrypt"
	url_decrypt_data     = "/wallet/decrypt"
	url_get_tx           = "/tx/:hash"
	url_deploy_contract  = "/contract/deploy"
	url_write_contract   = "/contract/write"
	url_read_contract    = "/contract/read"
)

const (
	SUGGESTED_GAS_LIMIT_TRANSFER = 21000
	SUGGESTED_GAS_LIMIT_INTERACT = 1000000
	SUGGESTED_GAS_LIMIT_DEPLOY   = 3000000
)

const (
	ERR_SHAMIR_INVALID_FRAGMENT_COUNT = "e611a97e-443e-4744-b934-70930ab40975"
	ERR_SHAMIR_FRAGMENT_SERVER_ERROR  = "ff034a14-ee70-4b80-bf44-6d5ef8e59e5b"
	ERR_WALLET_CREATE_FAILED          = "f0fcedbf-026f-4b12-9805-7e9aca6bd3b9"
	ERR_WALLET_BALANCE_GET_FAILED     = "31964065-298c-427f-95fa-c9c7584c9e6a"
	ERR_WALLET_TRANSFER_FAILED        = "956b3be0-d68d-451c-8777-f12d7228da19"
	ERR_WALLET_NOT_FOUND              = "948dcaf4-27b7-4458-b818-28b521061d5b"
	ERR_WALLET_LIST_FAILED            = "79fce109-87cb-4290-828e-0396ce5862b8"
	ERR_WALLET_DELETE_FAILED          = "964ef65d-49ec-4cf8-9820-241c803f3208"
	ERR_WALLET_CLIENT_SECRET_INVALID  = "f11dbfbf-5b0c-4104-b13f-5a986ed27243"
	ERR_WALLET_HAS_PENDING_TX         = "68ddb5a8-1afe-40f6-bbf9-b7f85802d998"
	ERR_WALLET_ENCRYPT_DATA_FAILED    = "372d8b5b-5a13-4200-801f-4956fe7e7489"
	ERR_WALLET_DECRYPT_DATA_FAILED    = "d1edec3c-2b3d-419a-b2d3-31362e4e8ad3"
	ERR_WALLET_INSUFFICIENT_BALANCE   = "17e5f76c-255b-46a1-b08b-9fa74390f2df"
	ERR_CONTRACT_INVALID_ARGS         = "f23a99cc-3e93-407c-a304-c06c1619d1f3"
	ERR_CONTRACT_DEPLOY_FAILED        = "d9b14078-dee4-4d75-b654-1625a95824c0"
	ERR_CONTRACT_READ_FAILED          = "aa7dc2a2-bf16-4694-b3d0-3f425481d9cb"
	ERR_CONTRACT_WRITE_FAILED         = "775fe7c5-130f-42aa-9e37-0b17feaef201"
	ERR_TX_NOT_FOUND                  = "05636027-31d8-467f-8f4b-3b9f0eb732c8"
	ERR_TX_CREATE_FAILED              = "206fa05d-6784-4be4-9736-67ed4e267858"
	ERR_TX_UPDATE_FAILED              = "f6aefdfa-775e-49e8-9af7-258c6e598c82"
)
