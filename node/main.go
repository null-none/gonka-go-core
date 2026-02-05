package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Prompt string `json:"prompt"`
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func main() {
	http.HandleFunc("/v1/inference", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Processing prompt: %s\n", req.Prompt)

		ollamaBody, _ := json.Marshal(OllamaRequest{
			Model:  "llama3",
			Prompt: req.Prompt,
			Stream: false,
		})

		resp, err := http.Post("http://host.docker.internal:11434/api/generate", "application/json", bytes.NewBuffer(ollamaBody))
		if err != nil {
			fmt.Println("Error: Ollama is not running on host")
			http.Error(w, "Ollama is offline", 500)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, resp.Body)
	})

	fmt.Println("🚀 Gonka M3 Node running on :8080")
	http.ListenAndServe(":8080", nil)
}
