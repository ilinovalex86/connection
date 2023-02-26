package connection

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

func timeCount(sec int) (int, int) {
	m := sec / 60
	s := sec % 60
	return m, s
}
