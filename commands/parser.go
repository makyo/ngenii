package commands

import "net"

func ParseCommand(conn net.Conn, input string) {
	if input == "QUIT" {
		conn.Close()
	}
}
