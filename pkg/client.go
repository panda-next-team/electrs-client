package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type HTTPClient struct {
	Client *resty.Client
}

func NewHTTPClient(hostUrl string, debugMode bool) *HTTPClient {
	restClient := resty.New()
	restClient.SetDebug(debugMode)
	restClient.SetHostURL(hostUrl)
	restClient.SetHeader("Accept", "application/json")
	restClient.SetContentLength(true)
	return &HTTPClient{restClient}
}

func (c *HTTPClient) GetTransaction(txID TxID) (*Transaction, error) {
	uri := fmt.Sprintf("/tx/%s", txID)
	result, err := c.doGet(uri, &Transaction{})
	if err != nil {
		return nil, err
	}
	return result.(*Transaction), nil
}

func (c *HTTPClient) GetTransactionStatus(txID TxID) (*TransactionStatus, error) {
	uri := fmt.Sprintf("/tx/%s/status", txID)
	result, err := c.doGet(uri, &TransactionStatus{})
	if err != nil {
		return nil, err
	}
	return result.(*TransactionStatus), nil
}

func (c *HTTPClient) GetTransactionHex(txID TxID) (TxHex, error) {
	uri := fmt.Sprintf("/tx/%s/hex", txID)
	result, err := c.doGetBody(uri)
	if err != nil {
		return "", err
	}
	return TxHex(result), nil
}

func (c *HTTPClient) GetTransactionMerkleProof(txID TxID) (*TransactionMerkleProof, error) {
	uri := fmt.Sprintf("/tx/%s/merkle-proof", txID)
	result, err := c.doGet(uri, &TransactionMerkleProof{})
	if err != nil {
		return nil, err
	}
	return result.(*TransactionMerkleProof), nil
}

func (c *HTTPClient) GetTransactionOutSpend(txID TxID, vOut int32) (*TransactionOutSpend, error) {
	uri := fmt.Sprintf("/tx/%s/outspend/%d", txID, vOut)
	result, err := c.doGet(uri, &TransactionOutSpend{})
	if err != nil {
		return nil, err
	}
	return result.(*TransactionOutSpend), nil
}

func (c *HTTPClient) GetTransactionOutSpends(txID TxID) ([]*TransactionOutSpend, error) {
	uri := fmt.Sprintf("/tx/%s/outspends", txID)
	txOutSpends := make([]*TransactionOutSpend, 0)
	result, err := c.doGetBody(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &txOutSpends)
	if err != nil {
		return nil, err
	}

	return txOutSpends, nil
}

func (c *HTTPClient) GetAddressInfo(address Address) (*AddressInfo, error) {
	uri := fmt.Sprintf("/address/%s", address)
	result, err := c.doGet(uri, &AddressInfo{})
	if err != nil {
		return nil, err
	}
	return result.(*AddressInfo), nil
}

func (c *HTTPClient) GetScriptHashInfo(hash ScriptHash) (*ScriptHashInfo, error) {
	uri := fmt.Sprintf("/scripthash/%s", hash)
	result, err := c.doGet(uri, &ScriptHashInfo{})
	if err != nil {
		return nil, err
	}
	return result.(*ScriptHashInfo), nil
}

func (c *HTTPClient) GetAddressTransactions(address Address) ([]*Transaction, error) {
	uri := fmt.Sprintf("/address/%s/txs", address)
	result, err := c.doGetBody(uri)

	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, 0)
	err = json.Unmarshal(result, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *HTTPClient) GetScriptHashTransactions(hash ScriptHash) ([]*Transaction, error) {
	uri := fmt.Sprintf("/scripthash/%s/txs", hash)
	result, err := c.doGetBody(uri)

	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, 0)
	err = json.Unmarshal(result, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *HTTPClient) GetAddressTransactionsLatest(address Address, lastTxID TxID) ([]*Transaction, error) {
	uri := fmt.Sprintf("/address/%s/txs/chain/%s", address, lastTxID)
	result, err := c.doGetBody(uri)

	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, 0)
	err = json.Unmarshal(result, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *HTTPClient) GetScriptHashTransactionsLatest(address Address, lastTxID TxID) ([]*Transaction, error) {
	uri := fmt.Sprintf("/scripthash/%s/txs/chain/%s", address, lastTxID)
	result, err := c.doGetBody(uri)

	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, 0)
	err = json.Unmarshal(result, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *HTTPClient) GetAddressTransactionsInMemPool(address Address) ([]*Transaction, error) {
	uri := fmt.Sprintf("/address/%s/txs/mempool", address)
	result, err := c.doGetBody(uri)

	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, 0)
	err = json.Unmarshal(result, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *HTTPClient) GetScriptHashTransactionsInMemPool(hash ScriptHash) ([]*Transaction, error) {
	uri := fmt.Sprintf("/scripthash/%s/txs/mempool", hash)
	result, err := c.doGetBody(uri)

	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, 0)
	err = json.Unmarshal(result, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *HTTPClient) GetAddressUnspentTxOutputs(address Address) ([]*UnspentTransactionOutput, error) {
	uri := fmt.Sprintf("/address/%s/utxo", address)
	result, err := c.doGetBody(uri)
	if err != nil {
		return nil, err
	}

	unspentTxOutputs := make([]*UnspentTransactionOutput, 0)
	err = json.Unmarshal(result, &unspentTxOutputs)
	if err != nil {
		return nil, err
	}
	return unspentTxOutputs, nil
}

func (c *HTTPClient) GetScriptHashUnspentTxOutputs(hash ScriptHash) ([]*UnspentTransactionOutput, error) {
	uri := fmt.Sprintf("/scripthash/%s/utxo", hash)
	result, err := c.doGetBody(uri)
	if err != nil {
		return nil, err
	}

	unspentTxOutputs := make([]*UnspentTransactionOutput, 0)
	err = json.Unmarshal(result, &unspentTxOutputs)
	if err != nil {
		return nil, err
	}
	return unspentTxOutputs, nil
}

func (c *HTTPClient) GetBlock(hash BlockHash) (*Block, error) {
	uri := fmt.Sprintf("/block/%s", hash)
	result, err := c.doGet(uri, &Block{})
	if err != nil {
		return nil, err
	}
	return result.(*Block), nil
}

func (c *HTTPClient) GetBlockStatus(hash BlockHash) (*BlockStatus, error) {
	uri := fmt.Sprintf("/block/%s/status", hash)
	result, err := c.doGet(uri, &BlockStatus{})
	if err != nil {
		return nil, err
	}
	return result.(*BlockStatus), nil
}

func (c *HTTPClient) GetBlockTransactions(hash BlockHash, startIndex int32) ([]*Transaction, error) {
	uri := fmt.Sprintf("/block/%s/txs/%d", hash, startIndex)
	result, err := c.doGetBody(uri)
	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, 0)
	err = json.Unmarshal(result, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *HTTPClient) GetBlockTxIDs(hash BlockHash) ([]TxID, error) {
	uri := fmt.Sprintf("/block/%s/txids", hash)
	result, err := c.doGetBody(uri)
	if err != nil {
		return nil, err
	}

	txIds := make([]TxID, 0)
	err = json.Unmarshal(result, &txIds)
	if err != nil {
		return nil, err
	}
	return txIds, nil
}

func (c *HTTPClient) GetBlockTxID(hash BlockHash, index int32) (TxID, error) {
	uri := fmt.Sprintf("/block/%s/txid/%d", hash, index)
	result, err := c.doGetBody(uri)
	if err != nil {
		return "", err
	}
	return TxID(result), nil
}

func (c *HTTPClient) GetBlockHash(height BlockHeight) (BlockHash, error) {
	uri := fmt.Sprintf("/block-height/%d", height)
	result, err := c.doGetBody(uri)
	if err != nil {
		return "", err
	}
	return BlockHash(result), nil
}

func (c *HTTPClient) GetBlocks(height BlockHeight) ([]*Block, error) {
	uri := fmt.Sprintf("/blocks/%d", height)
	result, err := c.doGetBody(uri)

	if err != nil {
		return nil, err
	}

	blocks := make([]*Block, 0)
	err = json.Unmarshal(result, &blocks)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func (c *HTTPClient) GetLastBlockHeight() (BlockHeight, error) {
	uri := fmt.Sprintf("/blocks/tip/height")
	result, err := c.doGetBody(uri)
	if err != nil {
		return BlockHeight(0), err
	}
	blockHeightInt, err := strconv.Atoi(string(result))
	if err != nil {
		return BlockHeight(0), err
	}
	return BlockHeight(blockHeightInt), nil
}

func (c *HTTPClient) GetLastBlockHash() (BlockHash, error) {
	uri := fmt.Sprintf("/blocks/tip/hash")
	result, err := c.doGetBody(uri)
	if err != nil {
		return BlockHash(""), err
	}
	return BlockHash(result), nil
}

func (c *HTTPClient) GetMemPoolStatistics() (*MemPoolStatistics, error) {
	uri := fmt.Sprintf("/mempool")
	result, err := c.doGet(uri, &MemPoolStatistics{})
	if err != nil {
		return nil, err
	}
	return result.(*MemPoolStatistics), nil
}

func (c *HTTPClient) GetMemPoolTxIDs() ([]TxID, error) {
	uri := fmt.Sprintf("/mempool/txids")
	result, err := c.doGetBody(uri)
	if err != nil {
		return nil, err
	}

	txIds := make([]TxID, 0)
	err = json.Unmarshal(result, &txIds)
	if err != nil {
		return nil, err
	}
	return txIds, nil
}

func (c *HTTPClient) GetMemPoolRecentOverviews() ([]*MemPoolOverviewData, error) {
	uri := fmt.Sprintf("/mempool/recent")
	result, err := c.doGetBody(uri)
	if err != nil {
		return nil, err
	}

	overviews := make([]*MemPoolOverviewData, 0)
	err = json.Unmarshal(result, &overviews)
	if err != nil {
		return nil, err
	}
	return overviews, nil
}

func (c *HTTPClient) GetFeeEstimates() (*FeeEstimates, error) {
	uri := fmt.Sprintf("/fee-estimates")
	result, err := c.doGet(uri, &FeeEstimates{})
	if err != nil {
		return nil, err
	}
	return result.(*FeeEstimates), nil
}



func (c *HTTPClient) doGet(uri string, entity interface{}) (interface{}, error) {
	resp, err := c.Client.R().SetResult(entity).
		Get(uri)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("conn err: %s", err.Error()))
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("request err: %s", string(resp.Body())))
	}

	return resp.Result(), nil
}

func (c *HTTPClient) doGetBody(uri string) ([]byte, error) {
	resp, err := c.Client.R().Get(uri)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("conn err: %s", err.Error()))
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("request err: %s", string(resp.Body())))
	}
	return resp.Body(), nil
}
