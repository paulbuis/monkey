package context

import (
	"io"
)

type Context struct {
	in  io.Reader
	out io.Writer
}

func New(in io.Reader, out io.Writer) *Context {
	return &Context{in: in, out: out}
}

func (c *Context) In() io.Reader {
	return c.in
}

func (c *Context) Out() io.Writer {
	return c.out
}
