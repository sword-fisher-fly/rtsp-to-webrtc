//go:build windows
// +build windows

package webrtc

import (
	"io"
	"os"
)

type syslog struct {
	inner io.WriteCloser
}

type nopCloser struct {
	io.Writer
}

func (n *nopCloser) Close() error {
	return nil
}

func newSyslog(prefix string) (io.WriteCloser, error) {
	inner := &nopCloser{os.Stdout}

	return &syslog{
		inner: inner,
	}, nil
}

func (ls *syslog) Close() error {
	return ls.inner.Close()
}

func (ls *syslog) Write(p []byte) (int, error) {
	return ls.inner.Write(p)
}
