package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type Config struct {
	Filepath string `json:"filepath"`
}

func main() {
	fmt.Println("exec server.go")

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

	// ファイル情報取得
	fileInfo, err := os.Stat(fp)
	// ファイルがない場合は作成する
	if err != nil {
		if os.IsNotExist(err) {
			// 親ディレクトリ作成
			dir := filepath.Dir(fp)
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatal("failed to create directory:", err)
			}
		} else {
			log.Fatal("unknown error:", err)
		}
	}
	// 既にファイルが存在する場合はそのファイルを削除
	if fileInfo != nil {
		err = os.Remove(fp)
		if err != nil {
			log.Fatal("the existing file delete error:", err)
		}
		fmt.Println("the existing file deletion is successful")
	}

	// 名前付きパイプ作成
	err = syscall.Mkfifo(fp, 0600)
	if err != nil {
		log.Fatal("create named pipe error:", err)
	}
	log.Printf("create name pipe is successful(filepath: %s)", fp)

	// パイプファイルを開く
	pipe, err := os.OpenFile(fp, os.O_WRONLY, 0)
	if err != nil {
		log.Fatal("open pipe file error:", err)
	}
	defer pipe.Close()

	// 標準入力
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("please enter some words")
		input := scanner.Text()
		fmt.Println("input value is " + input)

		if input == "exit" {
			fmt.Println("program exit")
			break
		}

		// パイプに書き込み
		n, err := pipe.Write([]byte(input + "\n"))
		if err != nil {
			log.Fatal("pipe write error:", err)
		}
		log.Printf("write to pipe completed successfully. written %d bytes\n", n)
	}

}
