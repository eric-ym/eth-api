package service

import (
	"encoding/json"
	"errors"
	"eth/internal/models"
	"eth/libs/common"
	"eth/libs/db"
	"eth/libs/utils"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetAddressList(page, limit int) (AddressData, error) {
	result := AddressData{}
	var addrData []AddressInfo
	start := (page - 1) * limit
	modelA := &models.Address{}
	c, err := modelA.GetAddressNums(db.GetDB())
	if err != nil {
		return result, err
	}

	result.Total = int(c)
	if c == 0 || start > int(c) {
		return result, nil
	}

	modelB := &models.Balance{}

	err = modelB.GetNewRecord(db.GetDB())
	if err != nil {
		zap.S().Errorf("get new record failed with error %v", err)
		return result, err
	}

	bList, err := modelB.GetBalanceList(modelB.Year, modelB.Month, modelB.Day, modelB.Hour, page, limit, db.GetDB())
	if err != nil {
		zap.S().Errorf("get balance rank list error")
		return result, err
	}

	if len(bList) > 0 {
		var aList []string
		for _, bl := range bList {
			aList = append(aList, bl.Address)
		}

		addrList, err := modelA.GetAddressListBySlice(aList, db.GetDB())
		if err != nil {
			zap.S().Errorf("get balance address info error %v", err)
			return result, err
		}

		fmt.Println(addrList)

		aMap := map[string]models.Address{}
		for _, addr := range addrList {
			aMap[addr.Hash] = addr
		}

		for k, ba := range bList {
			addrM := AddressInfo{
				Rank:    (page-1)*limit + k + 1,
				Account: ba.Address,
				Balance: ba.Balance,
			}

			a, ok := aMap[ba.Address]
			if ok {
				addrM.LastUpdate = a.UpdatedTime
				addrM.RecentTransfer = a.LastTransTime
			}
			addrData = append(addrData, addrM)
		}

	}
	result.Data = addrData
	return result, nil
}

func GetAddressDetailInfo(hash string) (AddressDetailInfo, error) {
	result := AddressDetailInfo{}
	modelA := &models.Address{}

	err := modelA.GetAddressByHash(hash, db.GetDB())
	if err != nil {
		return result, err
	}
	result.Address = modelA.Hash
	result.AccountID = modelA.Id
	result.LastUpdate = modelA.UpdatedTime

	balance, err := getAddressBalance(hash, "latest")
	if err != nil {
		fmt.Println("balance  ", err.Error())
		return result, err
	}
	result.Balance = balance
	trans := &models.Transactions{}
	c, err := trans.GetMessageCountByHash(hash, db.GetDB())
	if err != nil {
		fmt.Println("count  ", err.Error())
		return result, err
	}
	result.Message = int(c)
	return result, nil

}

func GetAddressMessage(hash string, page, limit int) (AddressMessageData, error) {
	result := AddressMessageData{
		Page:  page,
		Limit: limit,
	}
	var err error
	var messageData []AddressMessageInfo

	trans := &models.Transactions{}
	c, err := trans.GetMessageCountByHash(hash, db.GetDB())
	if err != nil {
		return result, err
	}
	result.Total = int(c)

	messageList, err := trans.GetMessageListByAddressHash(hash, page, limit, db.GetDB())
	if err != nil {
		return result, err
	}
	for _, message := range messageList {
		fmt.Println("Message============")
		m := AddressMessageInfo{
			Id:     message.Hash,
			Height: message.BlockNumber,
			Time:   message.BlockTime,
			From:   message.From,
			To:     message.To,
		}
		v, err := utils.HexToFloat(message.Value)
		if err != nil {
			return result, err
		}
		m.Value = v
		if message.Status == "0x1" {
			m.Status = "OK"
		} else {
			m.Status = "False"
		}
		if message.From == hash {
			m.Method = "Send"
		}
		if message.To == hash {
			m.Method = "Receive"
		}
		messageData = append(messageData, m)
	}
	result.Data = messageData
	return result, nil
}

func GetAddressTrans(hash string, page, limit int) (AddressTransData, error) {
	result := AddressTransData{
		Page:  page,
		Limit: limit,
	}
	var err error
	var messageData []AddressTransInfo

	trans := &models.Transactions{}
	c, err := trans.GetMessageCountByHash(hash, db.GetDB())
	if err != nil {
		return result, err
	}
	result.Total = int(c)
	fmt.Println("page: ", page, " limit: ", limit)
	messageList, err := trans.GetMessageListByAddressHash(hash, page, limit, db.GetDB())
	if err != nil {
		return result, err
	}
	fmt.Println("messageList: ", messageList)

	for _, message := range messageList {
		m := AddressTransInfo{
			Id:   message.Hash,
			Time: message.BlockTime,
			From: message.From,
			To:   message.To,
		}
		v, err := utils.HexToFloat(message.Value)
		if err != nil {
			return result, err
		}
		m.Value = v
		if message.From == hash {
			m.Method = "Send"
		}
		if message.To == hash {
			m.Method = "Receive"
		}
		messageData = append(messageData, m)
	}
	result.Data = messageData
	return result, nil
}

func GetHomeAddressList(limit int) ([]AddressInfo, error) {
	var addrData []AddressInfo
	modelA := &models.Address{}
	modelB := &models.Balance{}
	err := modelB.GetNewRecord(db.GetDB())
	if err != nil {
		zap.S().Errorf("get new record failed with error %v", err)
		return addrData, err
	}
	page := 1
	bList, err := modelB.GetBalanceList(modelB.Year, modelB.Month, modelB.Day, modelB.Hour, page, limit, db.GetDB())
	if err != nil {
		zap.S().Errorf("get balance rank list error")
		return addrData, err
	}
	if len(bList) > 0 {
		var aList []string
		for _, bl := range bList {
			aList = append(aList, bl.Address)
		}
		addrList, err := modelA.GetAddressListBySlice(aList, db.GetDB())
		if err != nil {
			zap.S().Errorf("get balance address info error %v", err)
			return addrData, err
		}
		aMap := map[string]models.Address{}
		for _, addr := range addrList {
			aMap[addr.Hash] = addr
		}
		for k, ba := range bList {
			addrM := AddressInfo{
				Rank:    (page-1)*limit + k + 1,
				Account: ba.Address,
				Balance: ba.Balance,
			}
			a, ok := aMap[ba.Address]
			if ok {
				addrM.LastUpdate = a.UpdatedTime
				addrM.RecentTransfer = a.LastTransTime
			}
			addrData = append(addrData, addrM)
		}

	}
	return addrData, nil
}

// 图标数据

/*
*

	{
	  "xAxis": {
	    "data": ["15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"]
	  },
	  "yAxis": {
	    "type": "value"
	  },
	  "series": [
	    {
	      "data": [320, 332, 301, 334, 390, 415, 652, 789]
	    },
	  ],

}
*/
type AddressEChat struct {
	XAxis  map[string][]string `json:"xAxis"`
	YAxis  map[string]string   `json:"yAxis"`
	Series map[string][]string `json:"series"`
}

func GetAddressTrend(hash string) (AddressEChat, error) {
	result := AddressEChat{
		YAxis: map[string]string{
			"type": "value",
		},
	}
	var xData []string
	var yData []string

	now := time.Now()

	for i := 0; i < 8; i++ {
		hours := now.Hour()
		xData = append(xData, fmt.Sprintf("%d:00", hours))

		b := &models.Balance{}

		Year := now.Year()
		Day := now.Day()
		Hour := now.Hour()

		month := now.Format("01")
		monthInt, _ := strconv.ParseInt(month, 10, 64)
		Month := int(monthInt)

		err := b.GetAddressBalance(Year, Month, Day, Hour, hash, db.GetDB())
		if err != nil {
			zap.S().Errorf("Error getting address balance: %v", err)
			return result, err
		}

		yData = append(yData, fmt.Sprintf("%f", b.Balance))

		now = now.Add(-1 * time.Hour)
	}

	result.XAxis = map[string][]string{
		"data": xData,
	}
	result.Series = map[string][]string{
		"data": yData,
	}
	return result, nil
}

func getAddressBalance(hash string, blockhash string) (string, error) {
	requestBody := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBalance",
		"params": ["%s", "%s"],
		"id": 1
	}`, hash, blockhash)

	resp, err := http.Post(common.EthUrl, "application/json", strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	fmt.Println(data)

	message, ok := data["result"]
	if !ok {
		return "", errors.New("api return error")
	}

	result := message.(string)

	balance, err := utils.HexToFloat(result)
	if err != nil {
		fmt.Println("hex to float ", err.Error())
		return "", err
	}
	return balance, nil
}
