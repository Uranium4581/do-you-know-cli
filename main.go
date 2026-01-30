package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//go:embed AA.txt
var aa string

type KnowledgeBase struct {
	Content   string `json:"content"`
	Known     int    `json:"known"`
	Unknown   int    `json:"unknown"`
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
}

func main() {
	// --- HTTP with timeout ---
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://do-you-know.moromen.com/api/knowledge",
		nil,
	)
	if err != nil {
		fmt.Println("request生成エラー:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("HTTPエラー:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad status:", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body読み込みエラー:", err)
		return
	}

	var knowledge KnowledgeBase
	if err := json.Unmarshal(body, &knowledge); err != nil {
		fmt.Println("JSONデコードエラー:", err)
		return
	}

	// --- 表示 ---
	fmt.Println("ID:[" + knowledge.ID + "]")

	for _, r := range knowledge.Content {
		fmt.Print(string(r))
		time.Sleep(120 * time.Millisecond)
	}
	fmt.Println()

	time.Sleep(400 * time.Millisecond)

	// --- AA ---
	fmt.Print("\033[35m" + aa + "\033[0m")
	fmt.Println()
}
