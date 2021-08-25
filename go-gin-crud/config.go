package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	DbDriver   string `json:"DB_DRIVER"`
	DbHost     string `json:"DB_HOST"`
	DbPort     string `json:"DB_PORT"`
	DbName     string `json:"DB_NAME"`
	DbUsername string `json:"DB_USERNAME"`
	DbPassword string `json:"DB_PASSWORD"`
}

func initReadFile(path string) map[string]interface{} {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	dict := map[string]interface{}{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		dict[line[0]] = line[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dict
}
func loadConfig(path string, v interface{}) {
	maps := initReadFile(path)
	jsonData, errJ := json.Marshal(maps)
	if errJ != nil {
		log.Fatal("Json map to struct has an error", errJ.Error())
	}
	err := json.Unmarshal(jsonData, &v)
	if err != nil {
		log.Fatal("Can't load config environment", err.Error())
	}
}

func loadConfigFileName(path string, fileName string) Config {
	var config Config
	loadConfig(fmt.Sprintf("%s/%s", path, fileName), &config)
	return config
}
