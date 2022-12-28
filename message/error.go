package message

type errorcode uint

type MsgError struct {
	Code errorcode `json:"code"`
	Err  error     `json:"msg"`
}

func Error(code errorcode, err error) *MsgError {
	e := &MsgError{
		Code: code,
		Err:  err,
	}

	return e
}

const (
	ERROR_DB_FIND errorcode = 10000 + iota
	ERROR_DB_INSERT
)