package core

type Response struct {
	code   RespCode
	result interface{}
}

func (resp *Response) Ok() bool {
	if resp.code == HIT {
		return true
	}
	return false
}

func (resp *Response) Content() interface{} {
	return resp.result
}
