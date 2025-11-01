package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigSupportedChainsCoinGetPageReq struct {
	dto.Pagination     `search:"-"`
    HsConfigSupportedChainsCoinOrder
}

type HsConfigSupportedChainsCoinOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_supported_chains_coin"`
    ChainCode string `form:"chainCodeOrder"  search:"type:order;column:chain_code;table:hs_config_supported_chains_coin"`
    ChainName string `form:"chainNameOrder"  search:"type:order;column:chain_name;table:hs_config_supported_chains_coin"`
    NativeSymbol string `form:"nativeSymbolOrder"  search:"type:order;column:native_symbol;table:hs_config_supported_chains_coin"`
    WithdrawCoin string `form:"withdrawCoinOrder"  search:"type:order;column:withdraw_coin;table:hs_config_supported_chains_coin"`
    Network string `form:"networkOrder"  search:"type:order;column:network;table:hs_config_supported_chains_coin"`
    ExplorerUrl string `form:"explorerUrlOrder"  search:"type:order;column:explorer_url;table:hs_config_supported_chains_coin"`
    RpcEndpoint string `form:"rpcEndpointOrder"  search:"type:order;column:rpc_endpoint;table:hs_config_supported_chains_coin"`
    ChainId string `form:"chainIdOrder"  search:"type:order;column:chain_id;table:hs_config_supported_chains_coin"`
    WithdrawEnabled string `form:"withdrawEnabledOrder"  search:"type:order;column:withdraw_enabled;table:hs_config_supported_chains_coin"`
    DepositEnabled string `form:"depositEnabledOrder"  search:"type:order;column:deposit_enabled;table:hs_config_supported_chains_coin"`
    MinWithdrawAmount string `form:"minWithdrawAmountOrder"  search:"type:order;column:min_withdraw_amount;table:hs_config_supported_chains_coin"`
    WithdrawFee string `form:"withdrawFeeOrder"  search:"type:order;column:withdraw_fee;table:hs_config_supported_chains_coin"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_config_supported_chains_coin"`
    SortOrder string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:hs_config_supported_chains_coin"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:hs_config_supported_chains_coin"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_supported_chains_coin"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_supported_chains_coin"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_supported_chains_coin"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_supported_chains_coin"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_supported_chains_coin"`
    
}

func (m *HsConfigSupportedChainsCoinGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfigSupportedChainsCoinInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    ChainCode string `json:"chainCode" comment:"链代码，如 TRC20 / ERC20 / BEP20"`
    ChainName string `json:"chainName" comment:"链名称，如 Tron / Ethereum / BSC"`
    NativeSymbol string `json:"nativeSymbol" comment:"主币，如 TRX / ETH / BNB"`
    WithdrawCoin string `json:"withdrawCoin" comment:"当前链支持提现的币种"`
    Network string `json:"network" comment:"网络类型：mainnet/testnet"`
    ExplorerUrl string `json:"explorerUrl" comment:"区块浏览器URL前缀"`
    RpcEndpoint string `json:"rpcEndpoint" comment:"RPC节点或API Endpoint"`
    ChainId string `json:"chainId" comment:"链ID（EVM链使用）"`
    WithdrawEnabled string `json:"withdrawEnabled" comment:"提现是否启用：1=启用，0=关闭"`
    DepositEnabled string `json:"depositEnabled" comment:"充值是否启用：1=启用，0=关闭"`
    MinWithdrawAmount string `json:"minWithdrawAmount" comment:"最小提现金额"`
    WithdrawFee string `json:"withdrawFee" comment:"固定提现手续费（USDT）"`
    Status string `json:"status" comment:"状态：1=启用，0=停用"`
    SortOrder string `json:"sortOrder" comment:"排序权重"`
    Remark string `json:"remark" comment:"备注说明"`
    common.ControlBy
}

func (s *HsConfigSupportedChainsCoinInsertReq) Generate(model *models.HsConfigSupportedChainsCoin)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ChainCode = s.ChainCode
    model.ChainName = s.ChainName
    model.NativeSymbol = s.NativeSymbol
    model.WithdrawCoin = s.WithdrawCoin
    model.Network = s.Network
    model.ExplorerUrl = s.ExplorerUrl
    model.RpcEndpoint = s.RpcEndpoint
    model.ChainId = s.ChainId
    model.WithdrawEnabled = s.WithdrawEnabled
    model.DepositEnabled = s.DepositEnabled
    model.MinWithdrawAmount = s.MinWithdrawAmount
    model.WithdrawFee = s.WithdrawFee
    model.Status = s.Status
    model.SortOrder = s.SortOrder
    model.Remark = s.Remark
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigSupportedChainsCoinInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigSupportedChainsCoinUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    ChainCode string `json:"chainCode" comment:"链代码，如 TRC20 / ERC20 / BEP20"`
    ChainName string `json:"chainName" comment:"链名称，如 Tron / Ethereum / BSC"`
    NativeSymbol string `json:"nativeSymbol" comment:"主币，如 TRX / ETH / BNB"`
    WithdrawCoin string `json:"withdrawCoin" comment:"当前链支持提现的币种"`
    Network string `json:"network" comment:"网络类型：mainnet/testnet"`
    ExplorerUrl string `json:"explorerUrl" comment:"区块浏览器URL前缀"`
    RpcEndpoint string `json:"rpcEndpoint" comment:"RPC节点或API Endpoint"`
    ChainId string `json:"chainId" comment:"链ID（EVM链使用）"`
    WithdrawEnabled string `json:"withdrawEnabled" comment:"提现是否启用：1=启用，0=关闭"`
    DepositEnabled string `json:"depositEnabled" comment:"充值是否启用：1=启用，0=关闭"`
    MinWithdrawAmount string `json:"minWithdrawAmount" comment:"最小提现金额"`
    WithdrawFee string `json:"withdrawFee" comment:"固定提现手续费（USDT）"`
    Status string `json:"status" comment:"状态：1=启用，0=停用"`
    SortOrder string `json:"sortOrder" comment:"排序权重"`
    Remark string `json:"remark" comment:"备注说明"`
    common.ControlBy
}

func (s *HsConfigSupportedChainsCoinUpdateReq) Generate(model *models.HsConfigSupportedChainsCoin)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ChainCode = s.ChainCode
    model.ChainName = s.ChainName
    model.NativeSymbol = s.NativeSymbol
    model.WithdrawCoin = s.WithdrawCoin
    model.Network = s.Network
    model.ExplorerUrl = s.ExplorerUrl
    model.RpcEndpoint = s.RpcEndpoint
    model.ChainId = s.ChainId
    model.WithdrawEnabled = s.WithdrawEnabled
    model.DepositEnabled = s.DepositEnabled
    model.MinWithdrawAmount = s.MinWithdrawAmount
    model.WithdrawFee = s.WithdrawFee
    model.Status = s.Status
    model.SortOrder = s.SortOrder
    model.Remark = s.Remark
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigSupportedChainsCoinUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigSupportedChainsCoinGetReq 功能获取请求参数
type HsConfigSupportedChainsCoinGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigSupportedChainsCoinGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigSupportedChainsCoinDeleteReq 功能删除请求参数
type HsConfigSupportedChainsCoinDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigSupportedChainsCoinDeleteReq) GetId() interface{} {
	return s.Ids
}
