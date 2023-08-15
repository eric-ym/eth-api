package service

import (
	"encoding/json"
	"eth/internal/models"
	"eth/libs/config"
	"eth/libs/db"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const RangeLimit = 20

type Block struct {
	Number       int64         `json:"number"`
	Hash         string        `json:"hash"`
	ParentHash   string        `json:"parentHash"`
	Timestamp    int64         `json:"timestamp"`
	Difficulty   string        `json:"difficulty"`
	GasLimit     string        `json:"gasLimit"`
	GasUsed      string        `json:"gasUsed"`
	Coinbase     string        `json:"coinbase"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionIndex int64  `json:"transaction_index"`
	Type             string `json:"type"`
	BlockHash        string `json:"block_hash"`
	Hash             string `json:"hash"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Nonce            string `json:"nonce"`
	Input            string `json:"input"`
	BlockTime        int64  `json:"block_time"`
	Receipted        int    `json:"receipted"`
	Status           string `json:"status"`
	GasUsed          string `json:"gasUsed"`
}

type Receipt struct {
	Status  string `json:"status"`
	GasUsed string `json:"gasUsed"`
}

func TimeServ() {

	// 以太坊节点的URL
	ethURL := config.Config().Set.Url
	// 定义定时任务时间间隔
	interval := 5 * time.Second

	// 启动定时任务
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			// 获取最新区块号
			latestBlockNumber, err := getLatestBlockNumber()
			if err != nil {
				zap.S().Errorf("get last block num err %e", err)
				continue
			}

			// 获取最新的区块号
			currentBlockNumber, err := getCurrentBlockNumber(ethURL)
			if err != nil {
				zap.S().Errorf("get current block num err %e", err)
				continue
			}
			if latestBlockNumber == currentBlockNumber {
				continue
			}

			// 获取增量区块
			incrementalBlocks, err := getIncrementalBlocks(ethURL, latestBlockNumber, currentBlockNumber)
			if err != nil {
				log.Printf("get incremental blocks err %e", err)
				continue
			}

			// 处理增量区块
			for _, block := range incrementalBlocks {

				err = db.GetDB().Transaction(func(tx *gorm.DB) error {
					// 打印区块信息
					b := &models.Blocks{
						Number:       block.Number,
						TimeStamp:    block.Timestamp,
						Hash:         block.Hash,
						ParentHash:   block.ParentHash,
						Difficulty:   block.Difficulty,
						GasLimit:     block.GasLimit,
						GasUsed:      block.GasUsed,
						Coinbase:     block.Coinbase,
						MessageCount: len(block.Transactions),
					}
					err = b.AddBlockToDb(tx)
					if err != nil {
						log.Printf("add block to db err hash %s, err %v", b.Hash, err)
						return err
					}

					address := &models.Address{
						Hash:        block.Coinbase,
						UpdatedTime: block.Timestamp,
					}
					err = address.CreateOrUpdate(tx)
					if err != nil {
						log.Printf("add block to db err add address hash %s, err %v", b.Coinbase, err)
						return err
					}

					// 处理交易信息
					for _, txa := range block.Transactions {
						tr := &models.Transactions{
							TransactionIndex: txa.TransactionIndex,
							Type:             txa.Type,
							BlockHash:        txa.BlockHash,
							BlockNumber:      block.Number,
							Hash:             txa.Hash,
							From:             txa.From,
							To:               txa.To,
							Value:            txa.Value,
							Gas:              txa.Gas,
							GasPrice:         txa.GasPrice,
							Nonce:            txa.Nonce,
							BlockTime:        block.Timestamp,
							Status:           txa.Status,
							GasUsed:          txa.GasUsed,
						}
						err = tr.AddTransactionToDb(tx)
						if err != nil {
							log.Printf("add block to db err add message hash %s, err %v", tr.Hash, err)
							return err
						}
						if len(tr.From) > 0 {
							aFrom := &models.Address{
								Hash:          tr.From,
								UpdatedTime:   block.Timestamp,
								LastTransTime: block.Timestamp,
							}
							err = aFrom.CreateOrUpdate(tx)
							if err != nil {
								log.Printf("updage to db err add message from hash %s, err %v", tr.From, err)
								return err
							}
						}
						if len(tr.To) > 0 {
							aTo := &models.Address{
								Hash:          tr.To,
								UpdatedTime:   block.Timestamp,
								LastTransTime: block.Timestamp,
							}
							err = aTo.CreateOrUpdate(tx)
							if err != nil {
								log.Printf("updage to db err add message touser hash %s, err %v", tr.To, err)
								return err
							}
						}
					}
					return nil
				})
				if err != nil {
					log.Printf("--------------------insert into db ERROR:%v \n %d \n %s", err, block.Number, block.Hash)
					break
				}
			}
		}
	}
}

// 获取记录最大区块号
func getLatestBlockNumber() (string, error) {
	b := &models.Blocks{}
	blockNumber, err := b.GetLastBlock(db.GetDB())
	fmt.Println("----------------------------------------------------", blockNumber, err)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%x", b.Number), nil
}

// 获得最新区块号
func getCurrentBlockNumber(url string) (string, error) {
	requestBody := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_blockNumber",
		"params": [],
		"id": 1
	}`)

	resp, err := http.Post(url, "application/json", strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	type data struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      int64  `json:"id"`
		Result  string `json:"result"`
	}
	d := &data{}

	err = json.NewDecoder(resp.Body).Decode(d)
	if err != nil {
		return "", err
	}
	return d.Result, nil
}

// 最新区块高度
func getIncrementalBlocks(url string, startBlock, endBlock string) ([]Block, error) {
	fmt.Println("================================", startBlock, endBlock, "================================")
	startInt, err := strconv.ParseInt(startBlock, 0, 64)
	if err != nil {
		return nil, err
	}
	endInt, err := strconv.ParseInt(endBlock, 0, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var blockList []string
	for i := startInt + 1; i <= endInt; i++ {
		blockList = append(blockList, fmt.Sprintf("%x", i))
		if len(blockList) > RangeLimit {
			break
		}
	}

	var blocks []Block

	for _, blockHash := range blockList {
		block, err := getBlockDetails(url, blockHash)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

// 获取区块详细信息
func getBlockDetails(url, blockNum string) (Block, error) {
	requestBody := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": ["0x%s", true],
		"id": 1
	}`, blockNum)

	resp, err := http.Post(url, "application/json", strings.NewReader(requestBody))
	if err != nil {
		return Block{}, err
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Block{}, err
	}

	blockData := data["result"].(map[string]interface{})
	number, err := strconv.ParseInt(blockData["number"].(string), 0, 64)
	if err != nil {
		return Block{}, err
	}
	blockTime, err := strconv.ParseInt(blockData["timestamp"].(string), 0, 64)
	if err != nil {
		blockTime = 0
	}

	block := Block{
		Number:     number,
		Hash:       blockData["hash"].(string),
		ParentHash: blockData["parentHash"].(string),
		Timestamp:  blockTime,
		Difficulty: blockData["difficulty"].(string),
		GasLimit:   blockData["gasLimit"].(string),
		GasUsed:    blockData["gasUsed"].(string),
		Coinbase:   blockData["miner"].(string),
	}

	transactionsData := blockData["transactions"].([]interface{})
	transactions := make([]Transaction, len(transactionsData))

	for i, txData := range transactionsData {
		tx := txData.(map[string]interface{})
		ti := tx["transactionIndex"].(string)
		tid, err := strconv.ParseInt(ti, 0, 64)
		if err != nil {
			log.Println(err)
		}
		transaction := Transaction{
			TransactionIndex: tid,
			Type:             tx["type"].(string),
			Hash:             tx["hash"].(string),
			BlockHash:        tx["blockHash"].(string),
			From:             tx["from"].(string),
			Value:            tx["value"].(string),
			Gas:              tx["gas"].(string),
			GasPrice:         tx["gasPrice"].(string),
			Nonce:            tx["nonce"].(string),
			BlockTime:        blockTime,
		}
		if tx["to"] != nil {
			transaction.To = tx["to"].(string)
		}

		re, err := getTransactionReceipt(url, transaction.Hash)
		if err != nil {
			return Block{}, err
		}

		transaction.Status = re.Status
		transaction.GasUsed = re.GasUsed

		transactions[i] = transaction
	}

	block.Transactions = transactions

	return block, nil
}

func getTransactionReceipt(url, hash string) (Receipt, error) {
	requestBody := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getTransactionReceipt",
		"params": ["%s"],
		"id": 1
	}`, hash)

	resp, err := http.Post(url, "application/json", strings.NewReader(requestBody))
	if err != nil {
		return Receipt{}, err
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Receipt{}, err
	}

	reData := data["result"].(map[string]interface{})

	fmt.Println(reData)

	re := Receipt{
		Status:  reData["status"].(string),
		GasUsed: reData["gasUsed"].(string),
	}
	return re, nil
}
