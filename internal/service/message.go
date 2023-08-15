package service

import (
	"eth/internal/models"
	"eth/libs/db"
	"eth/libs/utils"
	"gorm.io/gorm"
)

func GetMessageList(page, limit int, fType int) (MessageData, error) {
	result := MessageData{
		Page:  page,
		Limit: limit,
	}
	tm := models.Transactions{}

	total, err := tm.GetTransactionsCount(fType, db.GetDB())
	if err != nil {
		return result, err
	}
	result.Total = int(total)
	if result.Total < (page-1)*limit || total == 0 {
		return result, nil
	}

	tList, err := tm.GetTransactionsList(page, limit, fType, db.GetDB())

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return result, err
		}
	}
	if len(tList) > 0 {
		for _, ti := range tList {
			m := MessageResponse{
				Id:     ti.Hash,
				Block:  ti.BlockHash,
				Height: ti.BlockNumber,
				Time:   ti.BlockTime,
				From:   ti.From,
				To:     ti.To,
			}

			v, err := utils.HexToFloat(ti.Value)
			if err != nil {
				return result, err
			}
			m.Value = v
			if ti.Status == "0x1" {
				m.Status = "OK"
			} else {
				m.Status = "False"
			}
			m.Method = "Send"
			result.Data = append(result.Data, m)
		}
	}
	return result, nil
}

func GetMessageDetail(hash string) (MessageDetail, error) {
	result := MessageDetail{
		Version: "2.0",
	}
	transModel := &models.Transactions{}
	err := transModel.GetTransactionByHash(hash, db.GetDB())
	if err != nil {
		return result, err
	}
	result.Nonce = transModel.Nonce
	result.GasUsed = transModel.GasUsed
	gasFee, err := utils.GetTransReward(transModel.GasPrice, transModel.GasUsed)
	if err != nil {
		return result, err
	}

	blockModel := models.Blocks{}
	err = blockModel.GetBlockInfoByHash(transModel.BlockHash, db.GetDB())

	if err != nil {
		return result, err
	}
	result.GasFee = gasFee
	if len(blockModel.GasLimit) > 3 {
		gasLimit, err := utils.HexToFloat(blockModel.GasLimit)
		if err != nil {
			return result, err
		}
		result.GasLimit = gasLimit
	}
	return result, nil
}

func GetMessageDetailInfo(hash string) (MessageResponse, error) {
	result := MessageResponse{}
	tm := &models.Transactions{}

	err := tm.GetTransactionByHash(hash, db.GetDB())
	if err != nil {
		return result, err
	}
	result.Id = tm.Hash
	result.Height = tm.BlockNumber
	result.From = tm.From
	result.Time = tm.BlockTime
	result.To = tm.To
	result.Block = tm.BlockHash
	result.Method = tm.Type

	if tm.Status == "0x1" {
		result.Status = "OK"
	} else {
		result.Status = "False"
	}

	v, err := utils.HexToFloat(tm.Value)
	if err != nil {
		return result, err
	}
	result.Value = v

	return result, nil
}

func GetMessageDetailTransactions(hash string) ([]MessageDetailTransaction, error) {
	var result []MessageDetailTransaction
	tm := &models.Transactions{}

	err := tm.GetTransactionByHash(hash, db.GetDB())
	if err != nil {
		return nil, err
	}

	if tm.Type != "0x0" {
		return result, nil
	}

	re := MessageDetailTransaction{
		From: tm.From,
		To:   tm.To,
		Type: tm.Type,
	}
	v, err := utils.HexToFloat(tm.Value)
	if err != nil {
		return nil, err
	}

	re.Value = v

	result = append(result, re)
	return result, nil
}
