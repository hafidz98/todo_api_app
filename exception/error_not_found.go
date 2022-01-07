package exception

type NotFound struct {
	Error string
	From  string
}

func NewNotFound(error string, from string) NotFound {
	return NotFound{Error: error, From: from}
}
