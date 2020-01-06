package pkg

type TxID string
type TxHex string
type Address string
type ScriptHash string
type BlockHash string
type BlockHeight int32

type Transaction struct {
	ID       TxID              `json:"txid"`
	Version  int32             `json:"version"`
	LockTime int32             `json:"locktime"`
	VIn      []*TransactionIn  `json:"vin"`
	VOut     []*TransactionOut `json:"vout"`
	Size     int32             `json:"size"`
	Weight   int32             `json:"weight"`
	Fee      int32             `json:"fee"`
	Status   TransactionStatus `json:"status"`
}

type TransactionOut struct {
	ScriptPubKey        string `json:"scriptpubkey"`
	ScriptPubKeyAsm     string `json:"scriptpubkey_asm"`
	ScriptPubKeyType    string `json:"scriptpubkey_type"`
	ScriptPubKeyAddress string `json:"scriptpubkey_address"`
	Value               int64  `json:"value"`
}

type TransactionIn struct {
	ID           TxID           `json:"txid"`
	VOut         int64          `json:"vout"`
	PrevOut      TransactionOut `json:"prevout"`
	ScriptSig    string         `json:"scriptsig"`
	ScriptSigAsm string         `json:"scriptsig_asm"`
	Witness      []string       `json:"witness"`
	IsCoinBase   bool           `json:"isCoinBase"`
	Sequence     int64          `json:"sequence"`
}

type TransactionStatus struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight BlockHeight  `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int32  `json:"block_time"`
}

type TransactionMerkleProof struct {
	BlockHeight BlockHeight    `json:"block_height"`
	Merkle      []string `json:"merkle"`
	Pos         int32    `json:"pos"`
}

type TransactionOutSpend struct {
	Spent  bool               `json:"spent"`
	ID     TxID               `json:"txid"`
	VInPos int32              `json:"vin"`
	Status *TransactionStatus `json:"status"`
}

type AddressInfo struct {
	Address    Address    `json:"address"`
	ChainStats ChainStats `json:"chain_stats"`
	MemStats   MemStats   `json:"mem_stats"`
}

type ScriptHashInfo struct {
	ScriptHash ScriptHash `json:"scripthash"`
	ChainStats ChainStats `json:"chain_stats"`
	MemStats   MemStats   `json:"mem_stats"`
}

type ChainStats struct {
	FoundedTxoCount int32   `json:"funded_txo_count"`
	FoundedTxoSum   float64 `json:"founded_txo_sum"`
	SpentTxoCount   int32   `json:"spent_txo_count"`
	SpentTxoSum     float64 `json:"spent_txo_sum"`
	TxCount         int32   `json:"tx_count"`
}

type MemStats ChainStats

type UnspentTransactionOutput struct {
	ID     TxID              `json:"txid"`
	VOut   int32             `json:"vout"`
	Status TransactionStatus `json:"status"`
	Value  int64             `json:"value"`
}

type Block struct {
	ID                BlockHash `json:"id"`
	Height            BlockHeight     `json:"height"`
	Version           int32     `json:"version"`
	Timestamp         int32     `json:"timestamp"`
	TxCount           int32     `json:"tx_count"`
	Size              int32     `json:"size"`
	Weight            int32     `json:"weight"`
	MerkleRoot        string    `json:"merkle_root"`
	PreviousBlockHash BlockHash `json:"previousblockhash"`
	Nonce             int32     `json:"nonce"`
	Bits              int32     `json:"bits"`
}

type BlockStatus struct {
	InBestChain bool      `json:"in_best_chain"`
	Height      BlockHeight     `json:"height"`
	NextBest    BlockHash `json:"next_best"`
}

type MemPoolStatistics struct {
	Count        int32         `json:"count"`
	VSize        int32         `json:"vsize"`
	TotalFee     int64         `json:"totalFee"`
	FeeHistogram []interface{} `json:"fee_histogram"`
}

type MemPoolOverviewData struct {
	ID    TxID  `json:"txid"`
	Fee   int32 `json:"fee"`
	VSize int32 `json:"vsize"`
	Value int64 `json:"value"`
}

type FeeEstimates map[string]float64
