package main

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type responseLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
