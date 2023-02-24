package connection

import (
	"encoding/json"
	"net"
)

// SendQuery - Отправляет структуру Query в conn
func SendQuery(q Query, conn net.Conn) error {
	data, err := json.Marshal(q)
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
