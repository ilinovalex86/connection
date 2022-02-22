package connection

type Query struct {
	Method string
	Query  string
}

type Response struct {
	Response string
	DataLen  int
	Err      string
}

type Event struct {
	Method string
	Event  string
	CorX   int
	CorY   int
	Ctrl   bool
	Shift  bool
}
