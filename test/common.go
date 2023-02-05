package test

import (
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

func newExpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client:   http.DefaultClient,
		BaseURL:  fmt.Sprintf("http://127.0.0.1:8080"), // TODO(vgalaxy): config
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			// httpexpect.NewDebugPrinter(t, false),
			httpexpect.NewCompactPrinter(t),
		},
	})
}

var (
	depth = 1
)

const (
	indentation = 2
	flag        = "*"
)

func describe(text string, fn func()) {
	for i := 0; i < depth; i++ {
		fmt.Print(flag)
	}
	fmt.Printf(" %s\n", text)
	depth += indentation
	fn()
	depth -= indentation
}
