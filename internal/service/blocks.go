package service

import (
	"eth/internal/models"
	"eth/libs/db"
	"eth/libs/utils"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func GetBlocks(start int64, long int) ([]BlocksResponse, error) {
	result := make([]BlocksResponse, 0, long)
	b := &models.Blocks{}
	blocks, err := b.GetBlockList(start, long, db.GetDB())
	if err != nil {
		return nil, err
	}
	var bh []string
	for _, bl := range blocks {
		if bl.MessageCount > 0 {
			bh = append(bh, bl.Hash)
		}
	}
	// transactions list
	var tl []models.Transactions
	var trMap map[string][]models.Transactions
	if len(bh) > 0 {
		tm := &models.Transactions{}
		tl, err = tm.GetTransactionsByBlockHashList(bh, db.GetDB())
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}
		if len(tl) > 0 {
			for _, tr := range tl {
				_, ok := trMap[tr.BlockHash]
				if !ok {
					trMap[tr.BlockHash] = []models.Transactions{}
				}
				trMap[tr.BlockHash] = append(trMap[tr.BlockHash], tr)
			}
		}
	}
	for _, blInfo := range blocks {
		br := BlocksResponse{
			BlockID: blInfo.Hash,
			Height:  blInfo.Number,
			Time:    blInfo.TimeStamp,
		}
		br.Data = []MessageInfo{}
		tList, ok := trMap[blInfo.Hash]
		if ok {
			for _, t := range tList {
				transaction := MessageInfo{
					Cid:     t.Hash,
					Message: blInfo.MessageCount,
				}
				if len(t.GasPrice) > 2 && len(t.GasUsed) > 2 {
					reward, err := utils.GetTransReward(t.GasPrice, t.GasUsed)
					if err != nil {
						return nil, err
					}
					transaction.Reward = reward
				}
				br.Data = append(br.Data, transaction)
			}
		}
		result = append(result, br)
	}
	return result, nil
}

func GetBlockDetail(hash string) (BlockInfoResponse, error) {
	result := BlockInfoResponse{}
	b := &models.Blocks{}
	err := b.GetBlockInfoByHash(hash, db.GetDB())
	if err != nil {
		return result, err
	}
	p := &models.Blocks{}
	err = p.GetBlockInfoByHash(b.ParentHash, db.GetDB())
	result.Cid = b.Hash
	result.Height = b.Number
	result.Time = b.TimeStamp
	result.Message = b.MessageCount
	result.ParentCid = b.ParentHash
	result.ParentWeight = p.Number
	result.StateRoot = fmt.Sprintf("%d", b.Number)
	var re float64
	if b.MessageCount > 0 {
		t := &models.Transactions{}
		tList, err := t.GetTransactionsByBlockHash(hash, db.GetDB())
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return result, err
			}
		}
		if len(tList) > 0 {

			for _, tr := range tList {
				if len(tr.GasPrice) > 2 && len(tr.GasUsed) > 2 {
					rewardString, err := utils.GetTransReward(tr.GasPrice, tr.GasUsed)
					if err != nil {
						return result, err
					}
					reward, err := strconv.ParseFloat(rewardString, 64)
					if err != nil {
						return result, err
					}
					re += reward
				}
			}
		}
	}
	result.Reward = fmt.Sprintf("%f", re)
	return result, nil
}

func GetBlockDetailMessage(hash string, page, limit int) (MessageData, error) {
	result := MessageData{
		Page:  page,
		Limit: limit,
	}
	result.Data = []MessageResponse{}
	b := &models.Blocks{}
	err := b.GetBlockInfoByHash(hash, db.GetDB())
	if err != nil {
		return result, nil
	}
	result.Total = b.MessageCount
	if b.MessageCount > 0 {
		t := &models.Transactions{}
		tl, err := t.GetTransactionsByBlockHashPage(hash, page, limit, db.GetDB())
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return result, err
			}
		}
		if len(tl) > 0 {
			for _, ti := range tl {
				m := MessageResponse{
					Id:     ti.Hash,
					Height: b.Number,
					Time:   b.TimeStamp,
					From:   ti.From,
					To:     ti.To,
				}
				if ti.Status == "0x1" {
					m.Status = "OK"
				} else {
					m.Status = "False"
				}
				if ti.From == hash {
					m.Method = "Send"
				}
				if ti.To == hash {
					m.Method = "Receive"
				}
				v, err := utils.HexToFloat(ti.Value)
				if err != nil {
					return result, err
				}
				m.Value = v
				result.Data = append(result.Data, m)
			}
		}
	}
	return result, nil
}
