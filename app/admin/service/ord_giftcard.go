package service

import (
	"errors"
	"fmt"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type OrdGiftcard struct {
	service.Service
}

// GetPage 获取OrdGiftcard列表
func (e *OrdGiftcard) GetPage(c *dto.OrdGiftcardGetPageReq, p *actions.DataPermission, list *[]models.OrdGiftcard, count *int64) error {
	var err error
	var data models.OrdGiftcard

	err = e.Orm.Model(&data).
		Select("ord_giftcard.*, ord_giftcard_region.category_id, ord_giftcard_category.name as category_name, ord_giftcard_category.logo as category_logo, ord_giftcard_region.region_code, ord_giftcard_region.currency_code").
		Joins("LEFT JOIN ord_giftcard_region ON ord_giftcard.region_id = ord_giftcard_region.id").
		Joins("LEFT JOIN ord_giftcard_category ON ord_giftcard_region.category_id = ord_giftcard_category.id").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdGiftcard对象
func (e *OrdGiftcard) Get(d *dto.OrdGiftcardGetReq, p *actions.DataPermission, model *models.OrdGiftcard) error {
	var data models.OrdGiftcard

	err := e.Orm.Model(&data).
		Select("ord_giftcard.*, ord_giftcard_region.category_id, ord_giftcard_category.name as category_name, ord_giftcard_category.logo as category_logo, ord_giftcard_region.region_code, ord_giftcard_region.currency_code").
		Joins("LEFT JOIN ord_giftcard_region ON ord_giftcard.region_id = ord_giftcard_region.id").
		Joins("LEFT JOIN ord_giftcard_category ON ord_giftcard_region.category_id = ord_giftcard_category.id").
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdGiftcard error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdGiftcard对象
func (e *OrdGiftcard) Insert(c *dto.OrdGiftcardInsertReq) error {
	var err error
	var data models.OrdGiftcard
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardService Insert error:%s \r\n", err)
		return err
	}
	if err := e.syncRegionDiscountRate(e.Orm, data.RegionId); err != nil {
		return err
	}
	return nil
}

// Update 修改OrdGiftcard对象
func (e *OrdGiftcard) Update(c *dto.OrdGiftcardUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.OrdGiftcard{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	oldRegionId := data.RegionId
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("OrdGiftcardService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	if err := e.syncRegionDiscountRate(e.Orm, data.RegionId); err != nil {
		return err
	}
	if oldRegionId != "" && oldRegionId != data.RegionId {
		if err := e.syncRegionDiscountRate(e.Orm, oldRegionId); err != nil {
			return err
		}
	}
	return nil
}

// Remove 删除OrdGiftcard
func (e *OrdGiftcard) Remove(d *dto.OrdGiftcardDeleteReq, p *actions.DataPermission) error {
	var data models.OrdGiftcard

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveOrdGiftcard error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// BatchSet 批量设置礼品卡（根据ID判断新增或更新）
func (e *OrdGiftcard) BatchSet(c *dto.OrdGiftcardBatchSetReq, p *actions.DataPermission) error {
	if len(c.Items) == 0 {
		return errors.New("礼品卡列表不能为空")
	}

	// 使用事务确保数据一致性
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		regionIdsToSync := make(map[string]struct{})
		for _, item := range c.Items {
			if item.Id > 0 {
				// 有ID，执行更新操作
				var existingData models.OrdGiftcard

				// 先查询是否存在，并检查权限
				err := tx.Scopes(
					actions.Permission(existingData.TableName(), p),
				).First(&existingData, item.Id).Error

				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						e.Log.Errorf("OrdGiftcard BatchSet: record not found or no permission, id=%d", item.Id)
						return fmt.Errorf("礼品卡ID=%d不存在或无权更新", item.Id)
					}
					e.Log.Errorf("OrdGiftcard BatchSet query error:%s", err)
					return err
				}

				oldRegionId := existingData.RegionId
				// 更新字段
				existingData.RegionId = item.RegionId
				existingData.Name = item.Name
				existingData.ValuesConfig = item.ValuesConfig
				existingData.CardType = item.CardType
				existingData.DiscountRate = item.DiscountRate
				existingData.Status = item.Status
				existingData.Extra = item.Extra
				existingData.UpdateBy = c.UpdateBy

				// 保存更新
				if err := tx.Save(&existingData).Error; err != nil {
					e.Log.Errorf("OrdGiftcard BatchSet update error:%s, id=%d", err, item.Id)
					return fmt.Errorf("更新礼品卡ID=%d失败: %s", item.Id, err.Error())
				}

				if existingData.RegionId != "" {
					regionIdsToSync[existingData.RegionId] = struct{}{}
				}
				if oldRegionId != "" && oldRegionId != existingData.RegionId {
					regionIdsToSync[oldRegionId] = struct{}{}
				}

				e.Log.Infof("Updated giftcard id=%d", item.Id)
			} else {
				// 无ID，执行新增操作
				newData := models.OrdGiftcard{
					RegionId:     item.RegionId,
					Name:         item.Name,
					ValuesConfig: item.ValuesConfig,
					CardType:     item.CardType,
					DiscountRate: item.DiscountRate,
					Status:       item.Status,
					Extra:        item.Extra,
				}
				newData.CreateBy = c.CreateBy

				if err := tx.Create(&newData).Error; err != nil {
					e.Log.Errorf("OrdGiftcard BatchSet create error:%s", err)
					return fmt.Errorf("创建礼品卡失败: %s", err.Error())
				}

				if newData.RegionId != "" {
					regionIdsToSync[newData.RegionId] = struct{}{}
				}

				e.Log.Infof("Created new giftcard id=%d", newData.Id)
			}
		}

		for regionId := range regionIdsToSync {
			if err := e.syncRegionDiscountRate(tx, regionId); err != nil {
				return err
			}
		}

		return nil
	})
}

func (e *OrdGiftcard) syncRegionDiscountRate(tx *gorm.DB, regionId string) error {
	if regionId == "" || regionId == "0" {
		return nil
	}

	var maxRate string
	if err := tx.Model(&models.OrdGiftcard{}).
		Select("COALESCE(MAX(discount_rate), 0)").
		Where("region_id = ?", regionId).
		Scan(&maxRate).Error; err != nil {
		e.Log.Errorf("OrdGiftcard syncRegionDiscountRate query error: regionId=%s, err=%s", regionId, err)
		return err
	}

	if maxRate == "" {
		maxRate = "0"
	}

	if err := tx.Model(&models.OrdGiftcardRegion{}).
		Where("id = ?", regionId).
		Update("discount_rate", maxRate).Error; err != nil {
		e.Log.Errorf("OrdGiftcard syncRegionDiscountRate update error: regionId=%s, err=%s", regionId, err)
		return err
	}

	return nil
}
