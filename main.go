package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
}


// AdjustedTime uint `json:"adjusted_time"`
// 	AltBlocksCount uint `json:"alt_blocks_count"`
// 	BlockSizeLimit uint `json:"block_size_limit"`
// 	BlockSizeMedian uint `json:"block_size_median"`
// 	BlockWeightLimit uint `json:"block_weight_limit"`
// 	BlockWeightMedian uint `json:"block_weight_median"`
// 	BootstrapDaemonAddress string `json:"bootstrap_daemon_address"`
// 	BusySyncing bool `json:"busy_syncing"`
// 	Credits uint `json:"credits"`
// 	CumulativeDifficulty uint `json:"cumulative_difficulty"`
// 	CumulativeDifficultyTop64 uint `json:"cumulative_difficulty_top64"`
// 	DatabaseSize uint `json:"database_size"`
// 	Difficulty uint `json:"difficulty"`
// 	DifficultyTop64 uint `json:"difficulty"`
// 	FreeSpace uint `json:"free_space"`
// 	GreyPeerlistSize uint `json:"grey_peerlist_size"`
// 	Height uint `json:"height"`
// 	HeightWithoutBootstrap uint `json:"height_without_bootstrap"`
// 	IncomingConnectionsCount uint `json:"incoming_connections_count"`
// 	MainNet bool `json:"mainnet"`
// 	NetType string `json:"nettype"`		
// 	Offline bool `json:"offline"`
// 	OutgoingConnectionsCount uint `json:"outgoing_connections_count"`
// 	RpcConnectionsCount uint `json:"rpc_connections_count"`
// 	Stagenet bool `json:"stagenet"`
// 	StartTime uint `json:"start_time"`
// 	Status string `json:"status"`
// 	Synchronized bool `json:"synchronized"`
// 	Target uint `json:"target"`
// 	TargetHeight uint `json:"target_height"`
// 	Testnet bool `json:"testnet"`
// 	TopBlockHash string `json:"top_block_hash"`
// 	TopHash string `json:"top_hash"`
// 	TxCount uint `json:"tx_count"`
// 	TxPoolSize uint `json:"tx_pool_size"`
// 	Untrusted bool `json:"untrusted"`
// 	UpdateAvailable bool `json:"update_available"`
// 	Version string `json:"version"`
// 	WasBootstrapEverUsed bool `json:"was_bootstrap_ever_used"`
// 	WhitePeerlistSize uint `json:"white_peerlist_size"`
// 	WideCumulativeDifficulty string `json:"wide_cumulative_difficulty"`
// 	WideDifficulty string `json:"wide_difficulty"`

type RPCResponse struct {
	ID     string `json:"id"`
	Result struct {
		Height uint `json:"height"`
		Status string `json:"status"`
		NetType string `json:"nettype"`	
		Synchronized bool `json:"synchronized"`
		FreeSpace uint `json:"free_space"`
		OutgoingConnectionsCount uint `json:"outgoing_connections_count"`
		TxCount uint `json:"tx_count"`
		
	} `json:"result"`
	JSONRPC string `json:"jsonrpc"`
}

func main() {
	rpcRequest := RPCRequest{
		JSONRPC: "2.0",
		ID:      "0",
		Method:  "get_info",
	}

	requestBody, err := json.Marshal(rpcRequest)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Send POST request to monerod
	url := "http://127.0.0.1:18081/json_rpc"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	var rpcResponse RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResponse); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	fmt.Println("Height:", rpcResponse.Result.Height)
	fmt.Println("Status:", rpcResponse.Result.Status)
	fmt.Println("Network:", rpcResponse.Result.NetType)
	fmt.Println("Syncronized:", rpcResponse.Result.Synchronized)

	GB := rpcResponse.Result.FreeSpace / 1073741824

	fmt.Println("Free space:", rpcResponse.Result.FreeSpace,"|", GB , "GB")
	fmt.Println("Outgoing Connections:", rpcResponse.Result.OutgoingConnectionsCount)
	fmt.Println("TxCount:", rpcResponse.Result.TxCount)
}
