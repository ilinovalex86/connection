package connection

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"net"
)

// ReadQRStruct - Получает структуру QueryResponse из conn
func ReadQRStruct(conn net.Conn) (QueryResponse, error) {
	var qr QueryResponse
	b, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return QueryResponse{}, errors.New("error read struct from conn")
	}
	b = bytes.TrimSpace(b)
	err = json.Unmarshal(b, &qr)
	if err != nil {
		return QueryResponse{}, err
	}
	return qr, nil
}
