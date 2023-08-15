package models

import "gorm.io/gorm"

type Balance struct {
	Id      int     `json:"id"`
	Address string  `json:"address"`
	Year    int     `json:"year"`
	Month   int     `json:"month"`
	Day     int     `json:"day"`
	Hour    int     `json:"hour"`
	Balance float64 `json:"balance"`
}

func (b *Balance) TableName() string {
	return "balance"
}

func (b *Balance) Create(db *gorm.DB) error {
	return db.Table(b.TableName()).Create(b).Error
}

func (b *Balance) GetHoursCount(y, m, d, h int, db *gorm.DB) (int64, error) {
	var c int64
	err := db.Table(b.TableName()).Where("year=? AND month=? AND day=? AND hour=?", y, m, d, h).Count(&c).Error
	return c, err
}

func (b *Balance) GetNewRecord(db *gorm.DB) error {
	return db.Table(b.TableName()).Order("id DESC").First(&b).Error
}

func (b *Balance) GetBalanceList(y, m, d, h int, page, limit int, db *gorm.DB) ([]Balance, error) {
	var result []Balance
	dbT := db.Table(b.TableName()).Where("year=? AND month=? AND day=? AND hour=?", y, m, d, h).Order("balance DESC")
	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		dbT = dbT.Offset(start).Limit(limit)
	}
	err := dbT.Find(&result).Error
	return result, err
}

func (b *Balance) GetAddressBalance(y, m, d, h int, hash string, db *gorm.DB) error {
	err := db.Table(b.TableName()).Where("year=? AND month=? AND day=? AND hour =? AND address=? ", y, m, d, h, hash).First(b).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
