package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ Usage: gonka-cli \"your prompt\"")
		return
	}

	prompt := os.Args[1]
	data, _ := json.Marshal(map[string]string{"prompt": prompt})

	resp, err := http.Post("http://gonka-node:8080/v1/inference", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("❌ Error: %s\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	fmt.Println("\n🤖 --- AI RESPONSE ---")
	fmt.Println(result["response"])
}
