package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type RPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
}

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

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	rpcRequest := RPCRequest{
		JSONRPC: "2.0",
		ID:      "0",
		Method:  "get_info",
	}

	requestBody, _ := json.Marshal(rpcRequest)

	url := "http://127.0.0.1:18081/json_rpc"
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	defer resp.Body.Close()

	var rpcResponse RPCResponse
	_ = json.NewDecoder(resp.Body).Decode(&rpcResponse)

	tmpl, _ := template.ParseFiles("index.gohtml")
	_ = tmpl.Execute(w, rpcResponse.Result)
}

func main() {
	http.HandleFunc("/", getInfoHandler)
	fmt.Println("Starting server on :8080...")
	_ = http.ListenAndServe(":8080", nil)
}
