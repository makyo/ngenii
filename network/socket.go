package network

import (
	"crypto/tls"
	"net"
)

func CreateTLSListener(bindAddr string, certPath string, keyPath string) (net.Listener, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	listener, err := tls.Listen("tcp", bindAddr, &config)
	if err != nil {
		return nil, err
	}

	return listener, nil
}
