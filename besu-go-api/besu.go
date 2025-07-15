package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const besuURL = "http://localhost:8545"

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"` //
	ID      int           `json:"id"`
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func getBlockNumber() (string, error) {
	reqBody := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}

	data, _ := json.Marshal(reqBody)
	resp, err := http.Post(besuURL, "application/json", bytes.NewBuffer(data))

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var rpcResp RPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return "", err
	}

	if rpcResp.Error != nil {
		return "", fmt.Errorf("RPC Err: %s", rpcResp.Error.Message)
	}

	return string(rpcResp.Result), nil

}

func main() {
	http.HandleFunc("/blockNumber", func(w http.ResponseWriter, r *http.Request) {
		{
			blockNum, err := getBlockNumber()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"blockNumber: %s}`, blockNum)
		}
	})

	log.Println("Go BESU API running on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
