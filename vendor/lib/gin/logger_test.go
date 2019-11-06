package gin

import (
	"net/http"
	"testing"
	"time"
)

func initServer() {
	engine := Default()
	engine.GET("/hi", func(c *Context) {})
	go engine.Run(":8091")
	time.Sleep(1 * time.Second)
}

func TestStdoutLogger(t *testing.T) {
	initServer()

	for i := 0; i < 10; i++ {
		http.Get("http://localhost:8091/hi")
	}
}
