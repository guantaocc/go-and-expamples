package main

import "io"

// 测试方法
type CountWriter struct {
	cw io.Writer
	C  int64
}

func (c *CountWriter) Write(p []byte) (int, error) {
	c.C = int64(len(p))
	n, err := c.cw.Write(p)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &CountWriter{cw: w}
	return c, &c.C
}
