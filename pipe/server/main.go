package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
)

type Config struct {
	Filepath string `json:"filepath"`
}

func main() {
	fmt.Println("exec server.go")

	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal("open file error:", err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("file read error:", err)
	}

	config := &Config{}
	err = json.Unmarshal(b, config)
	if err != nil {
		log.Fatal("json unmarshal error:", err)
	}

	fmt.Printf("filepath: %s\n", config.Filepath)

	err = syscall.Mkfifo(config.Filepath, 0600)
	if err != nil {
		log.Fatal("create named pipe error:", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if input == "exit" {
			break
		}

	}

}
