package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type KnowledgeBase struct {
	Content    string `json:"content"`
	Known      int    `json:"known"`
	Unknown    int    `json:"unknown"`
	Created_at string `json:"created_at"`
	Id         string `json:"id"`
}

func main() {
	aa, err := os.ReadFile("AA.txt")
	url := fmt.Sprintf("https://do-you-know.moromen.com/api/knowledge")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // := (コロンが先)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var knowledge KnowledgeBase
	err = json.Unmarshal(body, &knowledge)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	// fmt.Println("Did you know?:", knowledge.Content)
	// fmt.Println("Known:", knowledge.Known)
	// fmt.Println("Unknown:", knowledge.Unknown)
	// fmt.Println("Created at:", knowledge.Created_at)
	// fmt.Println("ID:", knowledge.Id)

	// カタカタ
	fmt.Println("ID:[" + knowledge.Id + "]")
	for _, char := range knowledge.Content {
		fmt.Print(string(char))
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Println()

	time.Sleep(500 * time.Millisecond)

	if err != nil {
		fmt.Println("AAアート読み込みエラー:", err)
	} else {
		fmt.Print("\033[35m" + string(aa) + "\033[0m")
	}
	fmt.Println()
}
