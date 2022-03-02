package connection

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

// ReadResponse - Получает структуру Response из conn
func ReadResponse(conn net.Conn) (Response, error) {
	var r responseJS
	b, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return Response{}, errors.New("error read struct from conn")
	}
	b = bytes.TrimSpace(b)
	err = json.Unmarshal(b, &r)
	if err != nil {
		return Response{}, err
	}
	if r.Err != "" {
		fmt.Println("yes err")
		return Response{Response: r.Response, DataLen: r.DataLen, Err: errors.New(r.Err)}, nil
	}
	fmt.Println("no err")
	return Response{Response: r.Response, DataLen: r.DataLen, Err: nil}, nil
}
