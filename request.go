package beatles

type request struct {
	c Context
}

func newRequest(c Context) Request {
	return &request{c: c}
}

func (r *request) GetHeader(key string) string {
	return string(r.c.RequestContext().Request.Header.Peek(key))
}

func (r *request) Method() string {
	return string(r.c.RequestContext().Request.Header.Method())
}

func (r *request) Path() string {
	return string(r.c.RequestContext().URI().Path())
}

func (r *request) PathVariable(key string) string {
	parameters := r.c.Parameters()

	value, exists := parameters[key]
	if !exists {
		return ""
	}

	return value
}
