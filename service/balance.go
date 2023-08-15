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
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HourBalance() {
	for {

		now := time.Now()

		Year := now.Year()
		Day := now.Day()
		Hour := now.Hour()

		month := now.Format("01")
		monthInt, err := strconv.ParseInt(month, 10, 64)
		Month := int(monthInt)

		bal := &models.Balance{}

		c, err := bal.GetHoursCount(Year, Month, Day, Hour, db.GetDB())
		if err != nil && err != gorm.ErrRecordNotFound {
			zap.S().Errorf("get balance count error: %v", err)
			continue
		}

		if c > 0 {
			zap.S().Infof("%d-%d-%d %d is already exsist", Year, Month, Day, Hour)
			time.Sleep(time.Minute)
			continue
		}

		addressList, err := AddressList()
		if err != nil {
			zap.S().Errorf("get address list error %v", err)
			continue
		}

		if len(addressList) > 0 {
			for _, address := range addressList {
				b, err := getAddressBalance(address)
				if err != nil {
					zap.S().Errorf("get address balance error %v", err)
					continue
				}

				balance, err := strconv.ParseFloat(b, 64)
				bl := &models.Balance{
					Address: address,
					Year:    now.Year(),
					Day:     now.Day(),
					Hour:    now.Hour(),
					Balance: balance,
				}
				month := now.Format("01")
				monthInt, err := strconv.ParseInt(month, 10, 64)
				bl.Month = int(monthInt)

				err = bl.Create(db.GetDB())
				if err != nil {
					zap.S().Errorf("insert address balance error %v", err)
					continue
				}
			}

		}

		time.Sleep(time.Minute)
	}
}

func AddressList() ([]string, error) {
	address := &models.Address{}
	return address.GetAddressStringList(db.GetDB())
}

func getAddressBalance(hash string) (string, error) {
	requestBody := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBalance",
		"params": ["%s", "latest"],
		"id": 1
	}`, hash)

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
