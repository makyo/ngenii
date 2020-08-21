package network

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/eigenfire/ngenii/commands"
)

func HandleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	writer.WriteString("Welcome to the server!\nPlease type QUIT to disconnect.\n")

	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Printf("Received data: %s", input)
		commands.ParseCommand(conn, strings.Fields(input)[0])
	}
}
