package main

import (
	"bufio"
	"fmt"
	forum "forum/src"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// openBrowser opens the specified URL in the default browser if the user confirms.
func openBrowser(url string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(forum.INFO, " Would you like to open the browser? \033[1;33m(y/n)\033[0m: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" {
		var cmd string
		var args []string
		switch runtime.GOOS {
		case "linux":
			cmd = "xdg-open"
		case "windows":
			cmd = "rundll32.exe"
			args = []string{"url.dll,FileProtocolHandler", url}
		case "darwin":
			cmd = "open"
		default:
			fmt.Println(forum.WARNING, "Unsupported platform.")
			return
		}
		if err := exec.Command(cmd, append(args, url)...).Start(); err != nil {
			fmt.Printf("%sFailed to open the browser. Reason: %v\n", forum.ERROR, err)
		} else {
			fmt.Println(forum.SUCCESS, "The browser has been opened successfully.")
		}
	} else if input == "n" {
		fmt.Println(forum.INFO, "Browser opening canceled.")
	} else {
		fmt.Println("Invalid input. ")
		openBrowser(url)
	}
}

// findValidPort returns a valid port that can be used for the web server.
func findValidPort() int {
	for i := 1024; i < 4000; i++ {
		if ln, err := net.Listen("tcp", fmt.Sprintf(":%d", i)); err == nil {
			ln.Close()
			return i
		}
	}
	return 0
}
