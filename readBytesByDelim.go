package connection

import (
	"bufio"
	"errors"
	"net"
)

// ReadByteByDelim - Получает бинарные данные из conn до спецсимвола
func ReadByteByDelim(conn net.Conn) ([]byte, error) {
	b, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return nil, errors.New("error read struct from conn")
	}
	return b, nil
}
