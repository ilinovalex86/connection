package connection

import (
	"encoding/json"
	"net"
)

// SendResponse - Отправляет структуру QueryResponse в conn
func SendResponse(r Response, conn net.Conn) error {
	data, err := json.Marshal(r)
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
