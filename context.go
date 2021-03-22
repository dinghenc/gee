package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H is a map
type H map[string]interface{}

// Context define request and response context
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// requst info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) Param(key string) string {
	value := c.Params[key]
	return value
}

// PostForm get POST param value
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query get GET param value
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status defines write status code to response header
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader defines set header for response
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String defines the write string method for response
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON defines the write json method for response
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data defines the write binary data method for response
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML defines the write html method for response
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
