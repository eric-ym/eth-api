package utils

import (
	"errors"
	"fmt"
	"math/big"
)

// func StrToInt64(str string) int64 {
//
// }

// HexToFloat value数据转化成浮点数
func HexToFloat(hexString string) (string, error) {
	valueInt, success := new(big.Int).SetString(hexString[2:], 16) // Removing the '0x' prefix from the hex string
	if !success {
		return "", fmt.Errorf("invalid hex string: %s", hexString)
	}
	//valueFloat, _ := new(big.Float).SetInt(valueInt).Float64()

	etherPerWei := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	result := new(big.Float).Quo(new(big.Float).SetInt(valueInt), new(big.Float).SetInt(etherPerWei))

	resultString := result.Text('f', 10)

	return resultString, nil
}

func GetTransReward(price, used string) (string, error) {
	gasPrice, ok := new(big.Int).SetString(price[2:], 16)
	if !ok {
		return "", errors.New("invalid price")
	}

	gasUsed, ok := new(big.Int).SetString(used[2:], 16)
	if !ok {
		return "", errors.New("invalid used")
	}

	fee := new(big.Int).Mul(gasPrice, gasUsed)

	etherPerWei := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	result := new(big.Float).Quo(new(big.Float).SetInt(fee), new(big.Float).SetInt(etherPerWei))

	resultString := result.Text('f', 10)

	return resultString, nil
}

func FormatHashrate(hashrate *big.Int) string {
	unit := ""

	value := new(big.Float).SetInt(hashrate)
	switch {
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)) >= 0:
		unit = "P"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil)) >= 0:
		unit = "T"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)) >= 0:
		unit = "G"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)) >= 0:
		unit = "M"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), nil)) >= 0:
		unit = "K"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), nil)))
	}
	return fmt.Sprintf("%s%s", value.Text('f', 5), unit)
}

func FormatHashratePer(hashrate *big.Int) (*big.Float, string) {
	unit := ""

	value := new(big.Float).SetInt(hashrate)
	switch {
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)) >= 0:
		unit = "P"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil)) >= 0:
		unit = "T"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)) >= 0:
		unit = "G"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)) >= 0:
		unit = "M"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)))
	case hashrate.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), nil)) >= 0:
		unit = "K"
		value.Quo(value, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), nil)))
	}
	return value, unit
}
