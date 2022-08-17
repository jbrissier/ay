package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
)

func getDesktopPath() string {
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
	ay_text := strings.Join(args, " ")
	current_time := time.Now().Format(time.RFC3339)

	filePath := getDesktopPath()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(fmt.Sprintf("%s: %s\n", current_time, ay_text)); err != nil {
		panic(err)
	}

	f.Close()
}
