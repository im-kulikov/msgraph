// Taken from https://raw.githubusercontent.com/google/google-api-go-client/43c645d4bcf9251ced36c823a93b6d198764aae4/examples/debug.go
package msgraph

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type logTransport struct {
	rt http.RoundTripper
}

func (t *logTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var buf bytes.Buffer

	fmt.Fprintf(os.Stderr, "\n[request]\n")
	if req.Body != nil {
		req.Body = ioutil.NopCloser(&readButCopy{req.Body, &buf})
	}
	req.Write(os.Stdout)
	if req.Body != nil {
		req.Body = ioutil.NopCloser(&buf)
	}
	fmt.Fprintf(os.Stderr, "\n[/request]\n")

	res, err := t.rt.RoundTrip(req)

	fmt.Fprintf(os.Stderr, "[response]\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	} else {
		body := res.Body
		res.Body = nil
		res.Write(os.Stdout)
		if body != nil {
			res.Body = ioutil.NopCloser(&echoAsRead{body})
		}
	}

	return res, err
}

type echoAsRead struct {
	src io.Reader
}

func (r *echoAsRead) Read(p []byte) (int, error) {
	n, err := r.src.Read(p)
	if n > 0 {
		//os.Stdout.Write(p[:n])
		fmt.Fprintf(os.Stderr, "%s", p[:n])
	}
	if err == io.EOF {
		fmt.Fprintf(os.Stderr, "\n[/response]\n")
	}
	return n, err
}

type readButCopy struct {
	src io.Reader
	dst io.Writer
}

func (r *readButCopy) Read(p []byte) (int, error) {
	n, err := r.src.Read(p)
	if n > 0 {
		r.dst.Write(p[:n])
	}
	return n, err
}
