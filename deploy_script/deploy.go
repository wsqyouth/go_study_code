package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: dpdeploy <ip> <username> <appname>")
		os.Exit(1)
	}

	ip := os.Args[1]
	username := os.Args[2]
	appName := os.Args[3] // Assuming the app name is 'app'
	buildApp(appName)
	deployApp(appName, username, ip)
}

func buildApp(appName string) {
	fmt.Println("Building app...")

	// Set environment variables for cross-compilation
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")

	_, filename, _, _ := runtime.Caller(0)
	rootPath := filepath.Dir(filename)
	appPath := filepath.Join(rootPath, appName+".go")

	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("go build -o %s %s", appName, appPath))
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error building app: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("App built successfully.")
}

/*
	func buildApp(appName string) {
		fmt.Println("Building app for Linux/AMD64...")

		// Set environment variables for cross-compilation
		os.Setenv("GOOS", "linux")
		os.Setenv("GOARCH", "amd64")

		// Build the app
		cmd := exec.Command("go", "build", "-o", fmt.Sprintf("./gobin/%s", appName))
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)
		cmd.Stdout = mw
		cmd.Stderr = mw

		if err := cmd.Run(); err != nil {
			fmt.Printf("Error building app: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("App built successfully.")
	}
*/
func deployApp(appName, username, ip string) {
	fmt.Println("Deploying app...")

	scpCmd := fmt.Sprintf("scp -P 36000 %s %s@%s:/tmp", appName, username, ip)
	execCommand(scpCmd)

	sshCmd := fmt.Sprintf("ssh -p 36000 %s@%s 'chmod 777 /tmp/%s && mv /tmp/%s ~/app && ~/app'", username, ip, appName, appName)
	execCommand(sshCmd)

	fmt.Println("App deployed and running.")
}

func execCommand(command string) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command '%s': %s\n", command, err)
		os.Exit(1)
	}
}
