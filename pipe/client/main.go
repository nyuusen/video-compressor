package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	Filepath string `json:"filepath"`
}

func main() {
	fmt.Println("exec client.go")

	// ファイルOpen処理(FD割り当て)
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal("open file error:", err)
	}
	// リソースリークを防止するためCloseする
	defer f.Close()

	// ファイル内容を読み取る
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("file read error:", err)
	}

	// ファイル内容を構造体に詰める
	config := &Config{}
	err = json.Unmarshal(b, config)
	if err != nil {
		log.Fatal("json unmarshal error:", err)
	}
	fp := config.Filepath
	log.Printf("filepath: %s\n", fp)

	// 名前付きファイルを取得
	pipe, err := os.OpenFile(fp, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("open named pipe file error:", err)
	}
	defer pipe.Close()

	scanner := bufio.NewScanner(pipe)
	for {
		if scanner.Scan() {
			data := scanner.Text()
			fmt.Println("receiving data from pipe:", data)
		} else {
			if scnErr := scanner.Err(); scnErr != nil {
				fmt.Println("pipe was closed:", scnErr)
				break
			}
		}
	}

}
