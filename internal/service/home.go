package service

import (
	"errors"
	"eth/internal/models"
	"eth/libs/db"
	"eth/libs/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

func HomeStandards() (Standard, error) {
	result := Standard{}
	// 获取最新高度
	blockModel := models.Blocks{}
	blockNum, err := blockModel.GetLastBlock(db.GetDB())
	if err != nil {
		return result, err
	}
	result.LastHeight = blockNum

	err = blockModel.GetBlockByNumber(blockNum, db.GetDB())
	if err != nil {
		return result, err
	}
	result.LastUpdate = blockModel.TimeStamp

	fmt.Println(blockModel.Difficulty)

	diff, ok := new(big.Int).SetString(blockModel.Difficulty[2:], 16)
	if !ok {
		return result, errors.New("difficulty not supported")
	}

	power := utils.FormatHashrate(diff)
	if err != nil {
		return result, err
	}
	result.HashPower = power

	startTime := blockModel.TimeStamp - 86400

	blockB := models.Blocks{}
	err = blockB.GetLastBlockByBeforeTime(startTime, db.GetDB())
	if err != nil {
		return result, err
	}
	powerInc, err := powerIncr(blockB, blockModel)
	if err != nil {
		return result, err
	}

	result.Power24 = powerInc

	h, err := hoursH(blockB, blockModel)
	if err != nil {
		return result, err
	}

	result.Productivity24 = h

	return result, nil
}

func Home24HoursReward() (HomeReward, error) {
	result := HomeReward{}
	// 24小时生成的块数
	now := time.Now().Unix()
	last := now - 84000
	blockModel := models.Blocks{}
	err := blockModel.GetLastBlockByBeforeTime(now, db.GetDB())

	if err != nil {
		return result, err
	}

	beforeBlock := &models.Blocks{}

	err = beforeBlock.GetLastBlockByBeforeTime(last, db.GetDB())

	blocks := blockModel.Number - beforeBlock.Number

	result.Bonus24 = blocks * 2
	return result, nil
}

func HomeEChat() ([]HomeEChatData, error) {
	var result []HomeEChatData
	startTime := time.Now().Add(-13 * time.Hour)
	beforeBloc := &models.Blocks{}
	for i := 0; i < 13; i++ {
		block := &models.Blocks{}

		temp := startTime.Truncate(time.Hour * 24).Unix()

		err := block.GetLastBlockByBeforeTime(temp, db.GetDB())
		if err != nil {
			return result, err
		}
		if i > 0 {
			m := HomeEChatData{
				Hour: startTime.Hour(),
			}
			diff, ok := new(big.Int).SetString(block.Difficulty[2:], 16)
			if !ok {
				return result, errors.New("difficulty not supported")
			}

			power := utils.FormatHashrate(diff)
			if err != nil {
				return result, err
			}
			m.Power = power
			inc, err := powerIncr(*beforeBloc, *block)
			if err != nil {
				return result, err
			}
			m.Incr = inc
			result = append(result, m)
		}
		beforeBloc = block
		startTime = startTime.Add(time.Hour)
	}
	return result, nil
}

// 24小时算力增长
func powerIncr(startBlock, endBlock models.Blocks) (string, error) {

	starDifficulty, ok := new(big.Int).SetString(startBlock.Difficulty[2:], 16)
	if !ok {
		return "", errors.New("Invalid start difficulty block difficulty " + startBlock.Difficulty)
	}

	endDifficulty, ok := new(big.Int).SetString(endBlock.Difficulty[2:], 16)
	if !ok {
		return "", errors.New("Invalid start difficulty block difficulty " + endBlock.Difficulty)
	}

	//startHashrate := new(big.Int).Div(new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil), starDifficulty)
	//
	//fmt.Println("2 > 256 :", new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil))
	//
	//fmt.Println(" start difficulty, ", starDifficulty)
	//
	//fmt.Println("24小时内算力增长: start :", startHashrate)
	//
	//endHashrate := new(big.Int).Div(new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil), endDifficulty)
	//fmt.Println(" end difficulty, ", endDifficulty)
	//
	//fmt.Println("24小时算力增长 end ：", endHashrate)

	hashrateIncrease := new(big.Int).Sub(endDifficulty, starDifficulty)

	//hashrateIncrease := new(big.Int).Sub(endHashrate, startHashrate)

	return utils.FormatHashrate(hashrateIncrease), nil
}

func hoursH(startBlock, endBlock models.Blocks) (string, error) {

	// 获取起始和结束时间内的难度
	startDifficulty := new(big.Int).Set(hexutil.MustDecodeBig(startBlock.Difficulty))
	endDifficulty := new(big.Int).Set(hexutil.MustDecodeBig(endBlock.Difficulty))

	// 计算区块数量和计算数量
	computationCount := new(big.Int).Add(startDifficulty, endDifficulty)

	// 计算挖矿产出效率
	efficiency := new(big.Int).Quo(computationCount, new(big.Int).SetInt64(2))

	fmt.Println("24小时内的挖矿产出效率:", efficiency.String())

	//计算M算力能够产生多少块
	eff, u := utils.FormatHashratePer(efficiency)

	res := new(big.Float).Quo(new(big.Float).SetInt64(1), eff)

	return fmt.Sprintf("%s xx/%s", res.Text('f', 8), u), nil
}
