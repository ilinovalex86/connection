package connection

import (
	"encoding/json"
	"fmt"
	"net"
)

// SendResponse - Отправляет структуру QueryResponse в conn
func SendResponse(r Response, conn net.Conn) error {
	newR := responseJS{Response: r.Response, DataLen: r.DataLen}
	if r.Err != nil {
		newR = responseJS{Response: r.Response, DataLen: r.DataLen, Err: fmt.Sprint(r.Err)}
	}
	data, err := json.Marshal(newR)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	err = SendBytes(data, conn)
	if err != nil {
		return err
	}
	return nil
}
