package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/tarm/serial"
)

const defaultPort = "/dev/ttyUSB0"
const defaultBaud = 115200

var rootCmd = &cobra.Command{
	Use: "",
}

var readCmd = &cobra.Command{
	Use:   "read -p [port] -b [baud]",
	Short: "Read data from serial",
	Long:  "Read data form serial on specific port and baud",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		baud, _ := cmd.Flags().GetInt("baud")

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
	Use:   "write -p [port] -b [baud]",
	Short: "Write bytes on serial",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		baud, _ := cmd.Flags().GetInt("baud")

		reader := bufio.NewReader(os.Stdin)

		var message []byte
		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break
			}

			message = append(message, line[:]...)
		}

		config := &serial.Config{
			Name: port,
			Baud: baud,
		}

		serialConnection, err := serial.OpenPort(config)
		if err != nil {
			panic(err)
		}

		size, err := serialConnection.Write(message)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d bytes were written", size)
	},
}

func main() {

	flag.Parse()

	rootCmd.AddCommand(readCmd, writeCmd)

	readCmd.Flags().StringP("port", "p", defaultPort, "Port path")
	readCmd.Flags().IntP("baud", "b", defaultBaud, "Baud rate")

	writeCmd.Flags().StringP("port", "p", defaultPort, "Port path")
	writeCmd.Flags().IntP("baud", "b", defaultBaud, "Baud rate")

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
