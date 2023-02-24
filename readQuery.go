package connection

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"net"
)

// ReadQuery - Получает структуру Query из conn
func ReadQuery(conn net.Conn) (Query, error) {
	var q Query
	b, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return Query{}, errors.New("error read struct from conn")
	}
	b = bytes.TrimSpace(b)
	err = json.Unmarshal(b, &q)
	if err != nil {
		return Query{}, err
	}
	return q, nil
}
