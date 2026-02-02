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
	fmt.Printf("\nEnter host and ssh port being used. (Ex. x.x.x.x:22 or hostname:22)\n\n\n")
	for i := 1; i <= count; i++ {
		fmt.Print("Enter host: ")
		fmt.Scanf("%s", &host)
		fmt.Println()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("cmds: ")
		cmds, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("You screwed something up: ", err)
		}
		s := strings.Split(cmds, ",")
		fmt.Println()
		hostCmds[host] = s
	}
	return hostCmds
}

func main() {
	user := "arista"
	pass := "arista"
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

	for host, commands := range hostData {
		config := &ssh.ClientConfig{
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			User:            user,
			Auth:            []ssh.AuthMethod{interactiveAuth},
		}
		fmt.Println("++++++++++++++++++++++++++++++++")
		fmt.Println("Connected to: ", host)
		fmt.Println("++++++++++++++++++++++++++++++++")
		conn, err := ssh.Dial("tcp", host, config)
		if err != nil {
			log.Fatal("Failed to dial: ", err)
		}
		defer conn.Close()
		sess, err := conn.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}
		defer sess.Close()
		stdin, err := sess.StdinPipe()
		if err != nil {
			log.Fatal("Failed to connect to remote devices stdin: ", err)
		}
		sess.Stdout = os.Stdout
		sess.Stderr = os.Stderr
		sess.Shell()
		fmt.Fprintf(stdin, "term len 0\n")
		for _, v := range commands {
			fmt.Fprintf(stdin, "%s\n", v)
		}
		stdin.Close()
		sess.Wait()
	}
}


