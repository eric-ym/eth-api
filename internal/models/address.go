package models

import "gorm.io/gorm"

type Address struct {
	Id            int64  `json:"id"`              // 自增id
	Hash          string `json:"hash"`            // 地址
	LastTransTime int64  `json:"last_trans_time"` // 最后交易时间
	UpdatedTime   int64  `json:"updated_time"`    // 更新时间
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) CreateAddress(db *gorm.DB) error {
	return db.Table(a.TableName()).Create(a).Error
}

func (a *Address) UpdateAddressTime(hash string, m map[string]interface{}, db *gorm.DB) error {
	return db.Table(a.TableName()).Where("hash=?", hash).Updates(m).Error
}

func (a *Address) GetAddressByHash(hash string, db *gorm.DB) error {
	err := db.Table(a.TableName()).Where("hash=?", hash).First(a).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (a *Address) CreateOrUpdate(db *gorm.DB) error {
	ad := &Address{}
	err := ad.GetAddressByHash(a.Hash, db)
	if err != nil {
		return err
	}
	if ad.Id > 0 {
		m := map[string]interface{}{
			"updated_time": a.UpdatedTime,
		}
		if a.LastTransTime == 0 {
			m["last_trans_time"] = ad.LastTransTime
		} else {
			m["last_trans_time"] = a.LastTransTime
		}

		err = a.UpdateAddressTime(a.Hash, m, db)
	} else {
		err = db.Table(a.TableName()).Create(a).Error
	}
	return err
}

func (a *Address) GetAddressNums(db *gorm.DB) (int64, error) {
	var c int64
	err := db.Table(a.TableName()).Count(&c).Error
	return c, err
}

func (a *Address) GetAddressList(page, limit int, db *gorm.DB) ([]Address, error) {
	var start int
	var result []Address
	var err error
	if page > 0 && limit > 0 {
		start = (page - 1) * limit
		err = db.Table(a.TableName()).Offset(start).Limit(limit).Find(&result).Error
	} else {
		err = db.Table(a.TableName()).Find(&result).Error
	}
	return result, err
}

func (a *Address) GetRecentUpdateAddressList(limit int, db *gorm.DB) ([]Address, error) {
	var result []Address
	err := db.Table(a.TableName()).Order("updated_time desc").Limit(limit).Find(&result).Error
	return result, err
}

func (a *Address) GetAddressStringList(db *gorm.DB) ([]string, error) {
	var result []string
	err := db.Table(a.TableName()).Pluck("hash", &result).Error
	return result, err
}

func (a *Address) GetAddressListBySlice(l []string, db *gorm.DB) ([]Address, error) {
	var result []Address
	err := db.Table(a.TableName()).Where("hash in (?)", l).Find(&result).Error
	return result, err
}
