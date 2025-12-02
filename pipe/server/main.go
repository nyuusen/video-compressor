package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Filepath string `json:"filepath"`
}

func main() {
	fmt.Println("exec server.go")

	f, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("open file error: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("file read error:%s\n", err)
		os.Exit(1)
	}

	config := &Config{}
	err = json.Unmarshal(b, config)
	if err != nil {
		fmt.Printf("json unmarshal error:%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("filepath: %s\n", config.Filepath)
}
