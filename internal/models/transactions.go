package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

/**
+-------------------+--------------+------+-----+---------+-------+
| Field             | Type         | Null | Key | Default | Extra |
+-------------------+--------------+------+-----+---------+-------+
| transaction_index | bigint       | NO   | PRI | 0       |       |
| type              | varchar(128) | NO   |     |         |       |
| block_hash        | varchar(128) | NO   |     |         |       |
| hash              | varchar(128) | NO   | UNI |         |       |
| from              | varchar(128) | NO   | MUL |         |       |
| to                | varchar(128) | NO   | MUL |         |       |
| value             | varchar(128) | NO   |     |         |       |
| gas               | varchar(128) | NO   |     |         |       |
| gas_price         | varchar(128) | NO   |     |         |       |
| nonce             | varchar(128) | NO   |     |         |       |
+-------------------+--------------+------+-----+---------+-------+
*/

type Transactions struct {
	TransactionIndex int64  `json:"transactions_index"`
	Type             string `json:"type"`
	BlockHash        string `json:"block_hash"`
	BlockNumber      int64  `json:"block_number"`
	Hash             string `json:"hash"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gas_price"`
	Nonce            string `json:"nonce"`
	Input            string `json:"input"`
	Status           string `json:"status"`
	BlockTime        int64  `json:"block_time"`
	Receipted        int    `json:"receipted"`
	GasUsed          string `json:"gas_used"`
}

func (t *Transactions) TableName() string {
	return "transactions"
}

func (t *Transactions) AddTransactionToDb(db *gorm.DB) error {
	return db.Table(t.TableName()).Create(t).Error
}

func (t *Transactions) GetTransactionsByBlockHash(blockHash string, db *gorm.DB) ([]Transactions, error) {
	var result []Transactions
	err := db.Table(t.TableName()).Where("block_hash = ?", blockHash).Find(&result).Error
	return result, err
}

func (t *Transactions) GetTransactionsByBlockHashPage(blockHash string, page, limit int, db *gorm.DB) ([]Transactions, error) {
	var result []Transactions
	start := (page - 1) * limit
	err := db.Table(t.TableName()).Where("block_hash = ?", blockHash).Offset(start).Limit(limit).Find(&result).Error
	return result, err
}

func (t *Transactions) GetTransactionsByBlockHashList(hashList []string, db *gorm.DB) ([]Transactions, error) {
	var result []Transactions
	if len(hashList) == 0 {
		return result, nil
	}
	err := db.Table(t.TableName()).Where("block_hash in (?)", hashList).Find(&result).Error
	return result, err
}

func (t *Transactions) GetTransactionsList(page, limit, fType int, db *gorm.DB) ([]Transactions, error) {
	var result []Transactions
	var err error
	start := (page - 1) * limit

	handle := db.Table(t.TableName())
	if fType == 0 {
		handle = handle.Where("`type`=?", "0x0")
	}

	if start > 0 && limit > 0 {
		handle = handle.Offset(start).Limit(limit)
	}
	err = handle.Find(&result).Error
	return result, err
}

func (t *Transactions) GetTransactionsCount(fType int, db *gorm.DB) (int64, error) {
	var count int64
	var err error
	switch fType {
	case 0:
		err = db.Table(t.TableName()).Where("type=?", "0x0").Count(&count).Error
	case 1:
		err = db.Table(t.TableName()).Count(&count).Error
	}
	return count, err
}

func (t *Transactions) GetTransactionByHash(hash string, db *gorm.DB) error {
	return db.Table(t.TableName()).Where("hash = ?", hash).First(t).Error
}

func (t *Transactions) GetMessageCountByHash(hash string, db *gorm.DB) (int64, error) {
	var count int64
	err := db.Table(t.TableName()).Where("`from` = ? or `to` = ?", hash, hash).Count(&count).Error
	return count, err
}

func (t *Transactions) GetMessageListByAddressHash(hash string, page, limit int, db *gorm.DB) ([]Transactions, error) {
	var result []Transactions
	fmt.Printf("MOdel page %d, limit %d \n", page, limit)
	start := (page - 1) * limit
	if start < 0 || limit < 1 {
		return result, errors.New("page or limit must be greater than 1")
	}
	err := db.Table(t.TableName()).Where("`from` = ? or `to` = ?", hash, hash).Offset(start).Limit(limit).Find(&result).Error

	fmt.Println(result)
	fmt.Println(err)
	return result, err
}
