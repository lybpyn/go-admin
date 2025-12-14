package dto

import (
	"strconv"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftcardWriteoffsGetPageReq struct {
	dto.Pagination `search:"-"`
	OrdGiftcardWriteoffsOrder
}

type OrdGiftcardWriteoffsOrder struct {
	Id                 string `form:"idOrder"  search:"type:order;column:id;table:ord_giftcard_writeoffs"`
	UserId             string `form:"userIdOrder"  search:"type:order;column:user_id;table:ord_giftcard_writeoffs"`
	OrderId            string `form:"orderIdOrder"  search:"type:order;column:order_id;table:ord_giftcard_writeoffs"`
	GiftCardId         string `form:"giftCardIdOrder"  search:"type:order;column:gift_card_id;table:ord_giftcard_writeoffs"`
	Status             string `form:"statusOrder"  search:"type:order;column:status;table:ord_giftcard_writeoffs"`
	Remark             string `form:"remarkOrder"  search:"type:order;column:remark;table:ord_giftcard_writeoffs"`
	CreateBy           string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_giftcard_writeoffs"`
	UpdateBy           string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_giftcard_writeoffs"`
	CreatedAt          string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_giftcard_writeoffs"`
	UpdatedAt          string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_giftcard_writeoffs"`
	DeletedAt          string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_giftcard_writeoffs"`
}

func (m *OrdGiftcardWriteoffsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftcardWriteoffsInsertReq struct {
	Id                  int    `json:"-" comment:"主键ID，自增"` // 主键ID，自增
	UserId              string `json:"userId" comment:"用户ID，表示提交/使用礼品卡的用户"`
	OrderId             string `json:"orderId" comment:"订单ID，关联 ord_user_orders.id，用于核销对应的订单"`
	GiftCardId          string `json:"giftCardId" comment:"礼品卡ID，关联礼品卡主表（若有）"`
	Status              int    `json:"status" comment:"核销状态：0=待核销，1=已核销，2=失败"`
	Remark              string `json:"remark" comment:"备注信息，例如失败原因、核销说明"`
	AdminRecognizedCode string `json:"adminRecognizedCode" comment:"管理员识别的兑换码"`
	PlatformSaleRate    string `json:"platformSaleRate" comment:"平台售卡汇率"`
	RecognizedCardValue string `json:"recognizedCardValue" comment:"识别的卡片面值"`
	FailureImageUrl     string `json:"failureImageUrl" comment:"失败时的截图URL"`
	SupplierId          string `json:"supplierId" comment:"收卡品牌商ID"`
	common.ControlBy
}

func (s *OrdGiftcardWriteoffsInsertReq) Generate(model *models.OrdGiftcardWriteoffs) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UserId, _ = strconv.Atoi(s.UserId)
	model.OrderId, _ = strconv.Atoi(s.OrderId)
	model.GiftCardId, _ = strconv.Atoi(s.GiftCardId)
	model.Status = s.Status
	model.Remark = s.Remark
	model.AdminRecognizedCode = s.AdminRecognizedCode
	model.PlatformSaleRate = s.PlatformSaleRate
	model.RecognizedCardValue = s.RecognizedCardValue
	model.FailureImageUrl = s.FailureImageUrl
	model.SupplierId = s.SupplierId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdGiftcardWriteoffsInsertReq) GetId() interface{} {
	return s.Id
}

type OrdGiftcardWriteoffsUpdateReq struct {
	Id                  int    `uri:"id" comment:"主键ID，自增"` // 主键ID，自增
	UserId              string `json:"userId" comment:"用户ID，表示提交/使用礼品卡的用户"`
	OrderId             string `json:"orderId" comment:"订单ID，关联 ord_user_orders.id，用于核销对应的订单"`
	GiftCardId          string `json:"giftCardId" comment:"礼品卡ID，关联礼品卡主表（若有）"`
	Status              int    `json:"status" comment:"核销状态：0=待核销，1=已核销，2=失败"`
	Remark              string `json:"remark" comment:"备注信息，例如失败原因、核销说明"`
	AdminRecognizedCode string `json:"adminRecognizedCode" comment:"管理员识别的兑换码"`
	PlatformSaleRate    string `json:"platformSaleRate" comment:"平台售卡汇率"`
	RecognizedCardValue string `json:"recognizedCardValue" comment:"识别的卡片面值"`
	FailureImageUrl     string `json:"failureImageUrl" comment:"失败时的截图URL"`
	SupplierId          string `json:"supplierId" comment:"收卡品牌商ID"`
	common.ControlBy
}

func (s *OrdGiftcardWriteoffsUpdateReq) Generate(model *models.OrdGiftcardWriteoffs) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UserId, _ = strconv.Atoi(s.UserId)
	model.OrderId, _ = strconv.Atoi(s.OrderId)
	model.GiftCardId, _ = strconv.Atoi(s.GiftCardId)
	model.Status = s.Status
	model.Remark = s.Remark
	model.AdminRecognizedCode = s.AdminRecognizedCode
	model.PlatformSaleRate = s.PlatformSaleRate
	model.RecognizedCardValue = s.RecognizedCardValue
	model.FailureImageUrl = s.FailureImageUrl
	model.SupplierId = s.SupplierId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdGiftcardWriteoffsUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardWriteoffsGetReq 功能获取请求参数
type OrdGiftcardWriteoffsGetReq struct {
	Id int `uri:"id"`
}

func (s *OrdGiftcardWriteoffsGetReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardWriteoffsDeleteReq 功能删除请求参数
type OrdGiftcardWriteoffsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdGiftcardWriteoffsDeleteReq) GetId() interface{} {
	return s.Ids
}

// OrdGiftcardWriteoffsBatchInsertReq 批量核销请求参数
type OrdGiftcardWriteoffsBatchInsertReq struct {
	OrderId      int                             `json:"orderId" binding:"required" comment:"订单ID"`
	WriteoffList []OrdGiftcardWriteoffsBatchItem `json:"writeoffList" binding:"required,min=1" comment:"核销列表"`
	CreateBy     int                             `json:"-" comment:"创建者"` // 不从前端传参，从context获取
	UpdateBy     int                             `json:"-" comment:"更新者"` // 不从前端传参，从context获取
}

// SetCreateBy 设置创建人id
func (e *OrdGiftcardWriteoffsBatchInsertReq) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *OrdGiftcardWriteoffsBatchInsertReq) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}

// OrdGiftcardWriteoffsBatchItem 批量核销项
type OrdGiftcardWriteoffsBatchItem struct {
	GiftCardId               int    `json:"giftCardId" comment:"礼品卡ID，关联 ord_giftcard.id"`
	AdminRecognizedCode      string `json:"adminRecognizedCode" binding:"required" comment:"管理员识别的兑换码"`
	RecognizedCardValue      string `json:"recognizedCardValue" comment:"识别的卡片面值"`
	UserLocalCurrencyAmount  string `json:"userLocalCurrencyAmount" comment:"用户入账的本地货币金额"`
	Status                   int    `json:"status" comment:"核销状态：0=待核销，1=已核销，2=失败"`
	Remark                   string `json:"remark" comment:"备注信息"`
	FailureImageUrl          string `json:"failureImageUrl" comment:"失败时的截图URL"`
	SupplierId               int    `json:"supplierId" comment:"收卡品牌商ID"`
	PlatformSettlementAmount string `json:"platformSettlementAmount" comment:"平台入账货币金额"`
}

// OrdGiftcardWriteoffsCalculateReq 计算用户入账金额请求参数
type OrdGiftcardWriteoffsCalculateReq struct {
	OrderId             int    `json:"orderId" binding:"required" comment:"订单ID"`
	GiftCardId          int    `json:"giftCardId" comment:"礼品卡ID（可选，用于面额校验）"`
	RecognizedCardValue string `json:"recognizedCardValue" binding:"required" comment:"识别的卡片面值"`
	DiscountRate        string `json:"discountRate" comment:"折扣率（可选，传入则使用此值计算，不传则使用礼品卡配置）"`
}

// OrdGiftcardWriteoffsCalculateResp 计算用户入账金额响应
type OrdGiftcardWriteoffsCalculateResp struct {
	UserLocalCurrencyAmount string                                      `json:"userLocalCurrencyAmount" comment:"用户将获得的本地货币金额（卡片面值 × 折扣率 × 汇率）"`
	UserCurrencyCode        string                                      `json:"userCurrencyCode" comment:"用户的货币代码"`
	ConfigRate              string                                      `json:"configRate" comment:"配置的汇率"`
	DiscountRate            string                                      `json:"discountRate" comment:"使用的折扣率"`
	IsCrypto                bool                                        `json:"isCrypto" comment:"是否虚拟币"`
	OrderCurrencyCode       string                                      `json:"orderCurrencyCode" comment:"源货币代码（提供折扣ID时为礼品卡区域货币，否则为订单货币）"`
	DenominationValidation  *OrdGiftcardWriteoffsDenominationValidation `json:"denominationValidation,omitempty" comment:"面额校验信息"`
}

// OrdGiftcardWriteoffsDenominationValidation 面额校验信息
type OrdGiftcardWriteoffsDenominationValidation struct {
	IsValid      bool     `json:"isValid" comment:"是否符合面额规则"`
	ErrorMessage string   `json:"errorMessage,omitempty" comment:"错误信息"`
	AllowedFixed []string `json:"allowedFixed,omitempty" comment:"允许的固定面额"`
	AllowedRange *struct {
		Min string `json:"min" comment:"最小值"`
		Max string `json:"max" comment:"最大值"`
	} `json:"allowedRange,omitempty" comment:"允许的面额区间"`
}
