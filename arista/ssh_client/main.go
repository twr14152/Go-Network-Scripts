package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
)

func commands() map[string][]string {
	var count int
	var host string

	fmt.Print("Number of hosts: ")
	fmt.Scanf("%d", &count)

	hostCmds := make(map[string][]string)

	fmt.Printf("\nEnter host and ssh port being used. (Ex. x.x.x.x:22 or hostname:22)\n\n")

	for i := 1; i <= count; i++ {
		fmt.Print("Enter host: ")
		fmt.Scanf("%s", &host)
		fmt.Println()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("cmds: ")
		cmds, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		s := strings.Split(cmds, ",")
		hostCmds[host] = s
	}

	return hostCmds
}

func main() {
	user := "admin"
	pass := "admin"

	hostData := commands()

	interactiveAuth := ssh.KeyboardInteractive(
		func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			answers := make([]string, len(questions))
			for i := range answers {
				answers[i] = pass
			}
			return answers, nil
		},
	)

	for host, cmds := range hostData {
		fmt.Println("++++++++++++++++++++++++++++++++")
		fmt.Println("Connected to:", host)
		fmt.Println("++++++++++++++++++++++++++++++++")

		config := &ssh.ClientConfig{
			User:            user,
			Auth:            []ssh.AuthMethod{interactiveAuth},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		conn, err := ssh.Dial("tcp", host, config)
		if err != nil {
			log.Fatal("Failed to dial:", err)
		}
		defer conn.Close()

		sess, err := conn.NewSession()
		if err != nil {
			log.Fatal("Failed to create session:", err)
		}
		defer sess.Close()
		//
        // Updated section
		// EOS requires a PTY for CLI 
		//
		modes := ssh.TerminalModes{
			ssh.ECHO: 0,
		}

		if err := sess.RequestPty("vt100", 80, 40, modes); err != nil {
			log.Fatal("PTY request failed:", err)
		}

		// Build buffer for commands
		var cmdBuffer strings.Builder
		cmdBuffer.WriteString("term len 0\n")

		for _, v := range cmds {
			cmd := strings.TrimSpace(v)
			if cmd == "" {
				continue
			}
			cmdBuffer.WriteString(cmd + "\n")
		}

		// Run commands and print output
		output, err := sess.CombinedOutput(cmdBuffer.String())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(output))
	}
}

