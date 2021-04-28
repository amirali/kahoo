package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tarm/serial"
)

var rootCmd = &cobra.Command{
    Use: "",
}

var readCmd = &cobra.Command{
    Use: "read [port] [baud]",
    Short: "Read data from serial",
    Long: "Read data form serial on specific port and baud",
    Run: func(cmd *cobra.Command, args []string) {
        port := args[0]
        baud, err := strconv.Atoi(args[1])
        if err != nil {
            panic(err)
        }

        config := &serial.Config{
            Name: port,
            Baud: baud,
        }

        serialConnection, err := serial.OpenPort(config)
        if err != nil {
            panic(err)
        }

        buffer := make([]byte, 1024)
        for {
            size, err := serialConnection.Read(buffer)
            if err != nil {
                panic(err)
            }

            fmt.Printf("%q", buffer[:size])
        }
    },
}

var writeCmd = &cobra.Command{
    Use: "write [port] [baud]",
    Short: "Write bytes on serial",
    Long: "",
    Run: func(cmd *cobra.Command, args []string) {
        reader := bufio.NewReader(os.Stdin)
        line, err := reader.ReadString('\n')
        if err != nil {
            panic(err)
        }

        bytes := []byte(line)

        port := args[0]
        baud, err := strconv.Atoi(args[1])
        if err != nil {
            panic(err)
        }

        config := &serial.Config{
            Name: port,
            Baud: baud,
        }

        serialConnection, err := serial.OpenPort(config)
        if err != nil {
            panic(err)
        }

        size, err := serialConnection.Write(bytes)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%d bytes were written", size)
    },
}

func main() {
    rootCmd.AddCommand(readCmd, writeCmd)
    err := rootCmd.Execute()
    if err != nil {
        panic(err)
    }
}
