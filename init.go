package connection

import "fmt"

type Query struct {
	Method  string
	Query   string
	DataLen int
}

type Response struct {
	Response string
	DataLen  int
	Err      error
}

type responseJS struct {
	Response string
	DataLen  int
	Err      string
}

func secToMinSec(sec int) (int, int) {
	m := sec / 60
	s := sec % 60
	return m, s
}

func speedCount(s float32) string {
	if s > 1024 {
		return fmt.Sprintf("%3.1f Mb/s", s/1024)
	}
	return fmt.Sprintf("%4.f Kb/s", s)
}
