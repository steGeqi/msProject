package common

type BusinessCode int

type Result struct {
	Code BusinessCode
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

func (r *Result) Fail(code BusinessCode, msg string) *Result {
	r.Msg = msg
	r.Code = code
	return r
}
