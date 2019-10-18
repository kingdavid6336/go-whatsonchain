package whatsonchain

import (
	"fmt"
	"testing"
)

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	if client.Parameters.Network != NetworkMain {
		t.Fatal("expected value to be main")
	}
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient()
	fmt.Println(client.Parameters.Network)
	// Output:main
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient()
	}
}

// TestClient_GetHealth tests the GetHealth()
func TestClient_GetHealth(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var resp string
	resp, err = client.GetHealth()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}
	if resp != "Whats On Chain" {
		t.Fatal("expected value was wrong", resp)
	}

}

// ExampleClient_GetHealth example using GetHealth()
func ExampleClient_GetHealth() {
	client, _ := NewClient()
	resp, _ := client.GetHealth()
	fmt.Println(resp)
	// Output:Whats On Chain
}

// TestClient_GetChainInfo tests the GetChainInfo()
func TestClient_GetChainInfo(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var resp *ChainInfo
	resp, err = client.GetChainInfo()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.BestBlockHash) == 0 {
		t.Fatal("failed to get best block hash")
	}

}

// TestClient_GetBlockByHash tests the GetBlockByHash()
func TestClient_GetBlockByHash(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var resp *BlockInfo
	hash := "0000000000000000025b8506c83450afe84f0318775a52c7b91ee64aad0d5a23"
	resp, err = client.GetBlockByHash(hash)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Hash != "0000000000000000025b8506c83450afe84f0318775a52c7b91ee64aad0d5a23" {
		t.Fatal("failed to get the block hash", resp.Hash)
	}

}

// TestClient_GetBlockByHeight tests the GetBlockByHeight()
func TestClient_GetBlockByHeight(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var resp *BlockInfo
	height := int64(604648)
	resp, err = client.GetBlockByHeight(height)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Hash != "0000000000000000025b8506c83450afe84f0318775a52c7b91ee64aad0d5a23" {
		t.Fatal("failed to get the block hash", resp.Hash)
	}

}

// TestClient_GetBlockPages tests the GetBlockPages()
func TestClient_GetBlockPages(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var resp *BlockPagesInfo
	hash := "000000000000000000885a4d8e9912f085b42288adc58b3ee5830a7da9f4fef4"
	resp, err = client.GetBlockPages(hash, 1)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(*resp) == 0 {
		t.Fatal("no transactions found", resp)
	}

}

// TestClient_GetTxByHash tests the GetTxByHash()
func TestClient_GetTxByHash(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var resp *TxInfo
	hash := "c1d32f28baa27a376ba977f6a8de6ce0a87041157cef0274b20bfda2b0d8df96"
	resp, err = client.GetTxByHash(hash)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Hash != "c1d32f28baa27a376ba977f6a8de6ce0a87041157cef0274b20bfda2b0d8df96" {
		t.Fatal("failed to get the tx hash", resp.Hash)
	}

}

// TestClient_BroadcastTx tests the BroadcastTx()
func TestClient_BroadcastTx(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	txHex := "0100000001d1bda0bde67183817b21af863adaa31fda8cafcf2083ca1eaba3054496cbde10010000006a47304402205fddd6abab6b8e94f36bfec51ba2e1f3a91b5327efa88264b5530d0c86538723022010e51693e3d52347d4d2ff142b85b460d3953e625d1e062a5fa2569623fb0ea94121029df3723daceb1fef64fa0558371bc48cc3a7a8e35d8e05b87137dc129a9d4598ffffffff0115d40000000000001976a91459cc95a8cde59ceda718dbf70e612dba4034552688ac00000000"
	_, err = client.BroadcastTx(txHex)
	if err == nil {
		t.Fatal("error should have occurred")
	}
}
