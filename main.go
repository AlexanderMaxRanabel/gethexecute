package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	// Start Geth light client
	gethProcess := exec.Command("geth", "--syncmode=light", "--cache=1024", "--rpc", "--rpcaddr=localhost", "--rpcport=8545", "--rpcapi=eth,net,web3,personal,debug", "--ws", "--wsaddr=localhost", "--wsport=8546", "--wsapi=eth,net,web3,personal,debug", "--wsorigins=*", "--allow-insecure-unlock")

	// Redirect Geth output to stdout
	gethProcess.Stdout = os.Stdout
	gethProcess.Stderr = os.Stderr

	// Start Geth process
	err := gethProcess.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Wait for Geth process to start
	time.Sleep(5 * time.Second)

	// Attach to Geth console
	gethConsole := exec.Command("geth", "attach", "http://localhost:8545")

	// Redirect console output to stdout
	gethConsole.Stdout = os.Stdout
	gethConsole.Stderr = os.Stderr

	// Start Geth console
	err = gethConsole.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Wait for Geth console to finish
	err = gethConsole.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// Wait for Geth process to finish
	err = gethProcess.Wait()
	if err != nil {
		fmt.Println("Geth process exited with error:", err)
	} else {
		fmt.Println("Geth process exited successfully.")
	}
}
