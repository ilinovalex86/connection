package connection

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"net"
)

// ReadEvent - Получает структуру Event из conn
func ReadEvent(conn net.Conn) (Event, error) {
	var e Event
	b, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return Event{}, errors.New("error read struct from conn")
	}
	b = bytes.TrimSpace(b)
	err = json.Unmarshal(b, &e)
	if err != nil {
		return Event{}, err
	}
	return e, nil
}
