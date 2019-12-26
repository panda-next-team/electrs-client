package pkg

import (
	"fmt"
	"testing"
)

const (
	hostUrl = "http://localhost:3000"
)

var (
	client = NewHTTPClient(hostUrl, false)
)

func TestHTTPClient_GetTransaction(t *testing.T) {
	transaction, err := client.GetTransaction("6b1d869856d857484ab0ac53575ac88f9a123616e22725deff7542537c827899")
	if err != nil {
		t.Fatal(err.Error())
	}

	if transaction.ID == "" {
		t.Errorf("invalid transaction id")
		t.Fail()
	}
}

func TestHTTPClient_GetTransactionStatus(t *testing.T) {
	txStatus, err := client.GetTransactionStatus("6b1d869856d857484ab0ac53575ac88f9a123616e22725deff7542537c827899")
	if err != nil {
		t.Fatal(err.Error())
	}
	if txStatus.BlockHeight <= 0 {
		t.Errorf("invalid transaction status")
		t.Fail()
	}
}

func TestHTTPClient_GetTransactionHex(t *testing.T) {
	txHex, err := client.GetTransactionHex("6b1d869856d857484ab0ac53575ac88f9a123616e22725deff7542537c827899")
	if err != nil {
		t.Fatal(err.Error())
	}

	if txHex == "" {
		t.Errorf("invalid transaction hex")
		t.Fail()
	}
}

func TestHTTPClient_GetTransactionMerkleProof(t *testing.T) {
	txMerkleProof, err := client.GetTransactionMerkleProof("6b1d869856d857484ab0ac53575ac88f9a123616e22725deff7542537c827899")
	if err != nil {
		t.Fatal(err.Error())
	}

	if txMerkleProof.BlockHeight <= 0 {
		t.Errorf("invalid transaction merkle proof")
		t.Fail()
	}
}

func TestHTTPClient_GetTransactionOutSpend(t *testing.T) {
	txOutSpend, err := client.GetTransactionOutSpend("9a7ba47d71b2526b9f9a4376ea83c7afe4bf13cb0957a17148ce4adbb8eb47b0", 1)
	if err != nil {
		t.Fatal(err.Error())
	}

	if txOutSpend.ID == "" {
		t.Errorf("invalid transaction out spend")
		t.Fail()
	}
}

func TestHTTPClient_GetTransactionOutSpends(t *testing.T) {
	txOutSpends, err := client.GetTransactionOutSpends("9a7ba47d71b2526b9f9a4376ea83c7afe4bf13cb0957a17148ce4adbb8eb47b0")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(txOutSpends) <= 0 {
		t.Errorf("invalid transaction out spend")
		t.Fail()
	}
}

func TestHTTPClient_GetAddressInfo(t *testing.T) {
	address := Address("152f1muMCNa7goXYhYAQC61hxEgGacmncB")
	addressInfo, err := client.GetAddressInfo(address)
	if err != nil {
		t.Fatal(err.Error())
	}

	if addressInfo.Address != address {
		t.Errorf("invalid address info")
		t.Fail()
	}
}

func TestHTTPClient_GetScriptHashInfo(t *testing.T) {
	scriptHash := ScriptHash("55c3e0412df763244b0fe23a5129cda6f606be45")
	scriptHashInfo, err := client.GetScriptHashInfo(scriptHash)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(scriptHashInfo)

	if scriptHashInfo.ScriptHash != scriptHash {
		t.Errorf("invalid scripthash info")
		t.Fail()
	}
}

//3c9018e8d5615c306d72397f8f5eef44308c98fb576a88e030c25456b4f3a7ac

func TestHTTPClient_GetAddressTransactions(t *testing.T) {
	transactions, err := client.GetAddressTransactions("152f1muMCNa7goXYhYAQC61hxEgGacmncB")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(transactions) == 0 {
		t.Errorf("invalid transactions")
		t.Fail()
	}
}

func TestHTTPClient_GetScriptHashTransactions(t *testing.T) {
	transactions, err := client.GetScriptHashTransactions("6fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(transactions) == 0 {
		t.Errorf("invalid transactions")
		t.Fail()
	}
}

func TestHTTPClient_GetAddressTransactionsLatest(t *testing.T) {
	transactions, err := client.GetAddressTransactionsLatest("152f1muMCNa7goXYhYAQC61hxEgGacmncB", "ae6077ec4bb4f3938eb204972d6dfaa6da74543af1ebaca666cce0453a0fe9a6")
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(transactions) == 0 {
		t.Errorf("invalid transactions")
		t.Fail()
	}
}

func TestHTTPClient_GetAddressTransactionsInMemPool(t *testing.T) {
	_, err := client.GetAddressTransactionsInMemPool("152f1muMCNa7goXYhYAQC61hxEgGacmncB")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHTTPClient_GetScriptHashTransactionsInMemPool(t *testing.T) {
	_, err := client.GetScriptHashTransactions("6fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHTTPClient_GetAddressUnspentTxOutputs(t *testing.T) {
	outputs, err := client.GetAddressUnspentTxOutputs("152f1muMCNa7goXYhYAQC61hxEgGacmncB")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(outputs) == 0 {
		t.Errorf("invalid unspent transaction outputs")
		t.Fail()
	}
}

func TestHTTPClient_GetScriptHashTransactionsLatest(t *testing.T) {
	_, err := client.GetScriptHashTransactions("6fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHTTPClient_GetBlock(t *testing.T) {
	hash := BlockHash("00000000000000000003efa46ef30fe654bca88081953f65b6ae217bceaec20a")
	block, err := client.GetBlock(hash)

	if err != nil {
		t.Fatal(err.Error())
	}
	if block.ID != hash {
		t.Errorf("invalid block")
		t.Fail()
	}
}

func TestHTTPClient_GetBlockStatus(t *testing.T) {
	hash := BlockHash("00000000000000000003efa46ef30fe654bca88081953f65b6ae217bceaec20a")
	blockHash, err := client.GetBlockStatus(hash)

	if err != nil {
		t.Fatal(err.Error())
	}

	if blockHash.Height <= 0 {
		t.Errorf("invalid block status")
		t.Fail()
	}
}

func TestHTTPClient_GetBlockTransactions(t *testing.T) {
	hash := BlockHash("00000000000000000003efa46ef30fe654bca88081953f65b6ae217bceaec20a")
	transactions, err := client.GetBlockTransactions(hash, 50)

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(transactions) == 0 {
		t.Errorf("invalid transactions")
		t.Fail()
	}
}

func TestHTTPClient_GetBlockTxIDs(t *testing.T) {
	hash := BlockHash("00000000000000000003efa46ef30fe654bca88081953f65b6ae217bceaec20a")
	txIds, err := client.GetBlockTxIDs(hash)

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(txIds) == 0 {
		t.Errorf("invalid txIds")
		t.Fail()
	}
}

func TestHTTPClient_GetBlockTxID(t *testing.T) {
	hash := BlockHash("00000000000000000003efa46ef30fe654bca88081953f65b6ae217bceaec20a")
	txID, err := client.GetBlockTxID(hash, 3)

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(txID) == 0 {
		t.Errorf("invalid txId")
		t.Fail()
	}
}

func TestHTTPClient_GetBlockHash(t *testing.T) {
	height := BlockHeight(49999)

	blockHash, err := client.GetBlockHash(height)

	if err != nil {
		t.Fatal(err.Error())
	}

	if blockHash != BlockHash("000000000845517b31c6820d83f25cff46429bf136a7515fe504116427e60f8e") {
		t.Errorf("invalid block hash")
		t.Fail()
	}
}

func TestHTTPClient_GetBlocks(t *testing.T) {
	height := BlockHeight(49999)
	blocks, err := client.GetBlocks(height)

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(blocks) == 0 {
		t.Errorf("invalid blocks")
		t.Fail()
	}
}

func TestHTTPClient_GetLastBlockHeight(t *testing.T) {
	height, err := client.GetLastBlockHeight()
	if err != nil {
		t.Fatal(err.Error())
	}

	if height == 0 {
		t.Errorf("invalid block height")
		t.Fail()
	}
}

func TestHTTPClient_GetLastBlockHash(t *testing.T) {
	hash, err := client.GetLastBlockHash()

	if err != nil {
		t.Fatal(err.Error())
	}

	if hash == "" {
		t.Errorf("invalid block hash")
		t.Fail()
	}
}

func TestHTTPClient_GetMemPoolStatistics(t *testing.T) {
	statistics, err := client.GetMemPoolStatistics()

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(statistics.FeeHistogram) == 0 {
		t.Errorf("invalid mempool statistics")
		t.Fail()
	}
}

func TestHTTPClient_GetMemPoolTxIDs(t *testing.T) {
	txIds, err := client.GetMemPoolTxIDs()

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(txIds) == 0 {
		t.Errorf("invalid mempool txids")
		t.Fail()
	}
}

func TestHTTPClient_GetMemPoolRecentOverviews(t *testing.T) {
	overviews, err := client.GetMemPoolRecentOverviews()
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(overviews) == 0 {
		t.Errorf("invalid mempool overviews")
		t.Fail()
	}
}

func TestHTTPClient_GetFeeEstimates(t *testing.T) {
	fees, err := client.GetFeeEstimates()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(*fees) == 0 {
		t.Errorf("invalid mempool overviews")
		t.Fail()
	}
}




