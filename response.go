package beatles

import "encoding/json"

type response struct {
	c Context
}

func newResponse(c Context) Response {
	return &response{c: c}
}

func (r *response) SetHeader(key string, value string) Response {
	r.c.RequestContext().Response.Header.Set(key, value)
	return r
}

func (r *response) SetStatus(statusCode int) Response {
	r.c.RequestContext().Response.SetStatusCode(statusCode)
	return r
}

func (r *response) SetBody(body []byte) error {
	r.c.RequestContext().Response.SetBody(body)
	return nil
}

func (r *response) JSON(v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	r.SetBody(b)
	r.SetHeader(HeaderContentType, ContentTypeJSON)

	return nil
}
