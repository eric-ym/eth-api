package models

import (
	"gorm.io/gorm"
)

type Blocks struct {
	Number       int64  // 区块高度
	TimeStamp    int64  // 产生时间
	Hash         string // hash
	ParentHash   string // 父区块哈希
	Difficulty   string // 难度
	GasLimit     string // Gas限制
	GasUsed      string // 已使用Gas
	Coinbase     string //矿工地址
	MessageCount int    // 消息数量
	StorageAt    string // storage at
}

func (b *Blocks) TableName() string {
	return "blocks"
}

// GetLastBlock 获得已经更新的最新区块
func (b *Blocks) GetLastBlock(db *gorm.DB) (int64, error) {
	err := db.Table(b.TableName()).Order("number desc").First(b).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return 0, err
		}
		return 0, nil
	}
	return b.Number, nil
}

func (b *Blocks) GetBlockInfoByHash(hash string, db *gorm.DB) error {
	err := db.Table(b.TableName()).Where("hash = ?", hash).First(b).Error
	return err
}

func (b *Blocks) AddBlockToDb(db *gorm.DB) error {
	return db.Table(b.TableName()).Create(b).Error
}

func (b *Blocks) GetBlockList(h int64, long int, db *gorm.DB) ([]Blocks, error) {
	result := make([]Blocks, long)
	err := db.Table(b.TableName()).Where("number >= ?", h).Limit(long).Find(&result).Error
	return result, err
}

func (b *Blocks) GetBlockByNumber(num int64, db *gorm.DB) error {
	return db.Table(b.TableName()).Where("number=?", num).First(b).Error
}

func (b *Blocks) GetBlockCount(start, end int64, db *gorm.DB) (int64, error) {
	var c int64
	err := db.Table(b.TableName()).Where("time_stamp > ? AND time_stamp < ?", start, end).Count(&c).Error
	return c, err
}

func (b *Blocks) GetHashLastBlockByTimeRange(start, end int64, hash string, db *gorm.DB) (string, error) {
	err := db.Table(b.TableName()).Where("coinbase=? AND time_stamp > ? AND time_stamp < ?", hash, start, end).Order("number desc").First(b).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return "", err
		} else {
			return "", nil
		}
	}
	return b.Hash, nil
}

func (b *Blocks) GetLastBlockByBeforeTime(start int64, db *gorm.DB) error {
	return db.Table(b.TableName()).Where("time_stamp < ?", start).Order("number Desc").First(b).Error
}
