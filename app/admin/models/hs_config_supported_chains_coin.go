package models

import (

	"go-admin/common/models"

)

type HsConfigSupportedChainsCoin struct {
    models.Model
    
    ChainCode string `json:"chainCode" gorm:"type:varchar(32);comment:链代码，如 TRC20 / ERC20 / BEP20"` 
    ChainName string `json:"chainName" gorm:"type:varchar(64);comment:链名称，如 Tron / Ethereum / BSC"` 
    NativeSymbol string `json:"nativeSymbol" gorm:"type:varchar(20);comment:主币，如 TRX / ETH / BNB"` 
    WithdrawCoin string `json:"withdrawCoin" gorm:"type:varchar(20);comment:当前链支持提现的币种"` 
    Network string `json:"network" gorm:"type:varchar(32);comment:网络类型：mainnet/testnet"` 
    ExplorerUrl string `json:"explorerUrl" gorm:"type:varchar(255);comment:区块浏览器URL前缀"` 
    RpcEndpoint string `json:"rpcEndpoint" gorm:"type:varchar(255);comment:RPC节点或API Endpoint"` 
    ChainId string `json:"chainId" gorm:"type:varchar(64);comment:链ID（EVM链使用）"` 
    WithdrawEnabled string `json:"withdrawEnabled" gorm:"type:tinyint(1);comment:提现是否启用：1=启用，0=关闭"` 
    DepositEnabled string `json:"depositEnabled" gorm:"type:tinyint(1);comment:充值是否启用：1=启用，0=关闭"` 
    MinWithdrawAmount string `json:"minWithdrawAmount" gorm:"type:decimal(20,8);comment:最小提现金额"` 
    WithdrawFee string `json:"withdrawFee" gorm:"type:decimal(20,8);comment:固定提现手续费（USDT）"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态：1=启用，0=停用"` 
    SortOrder string `json:"sortOrder" gorm:"type:int(11);comment:排序权重"` 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:备注说明"` 
    models.ModelTime
    models.ControlBy
}

func (HsConfigSupportedChainsCoin) TableName() string {
    return "hs_config_supported_chains_coin"
}

func (e *HsConfigSupportedChainsCoin) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigSupportedChainsCoin) GetId() interface{} {
	return e.Id
}