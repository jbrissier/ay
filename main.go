package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
)

func getAyFile() string {
	myself, error := user.Current()
	if error != nil {
		panic(error)
	}
	homedir := myself.HomeDir
	desktop := homedir + "/Desktop/ay.txt"
	return desktop
}

func main() {
	args := os.Args[1:]
	ayText := strings.Join(args, " ")
	currentTime := time.Now().Format(time.RFC3339)

	filePath := getAyFile()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(fmt.Sprintf("%s: %s\n", currentTime, ayText)); err != nil {
		panic(err)
	}

	f.Close()
}
