package main

import (
	"io"

	"golang.org/x/crypto/ssh"
)

type SessionIO struct {
	io.Reader
	io.Writer
}

func NewSessionIO(session *ssh.Session) (r io.Reader, w io.Writer, err error) {
	logger.Debugf("stdin")
	w, err = session.StdinPipe()
	if err != nil {
		return
	}

	logger.Debugf("stdout")
	r, err = session.StdoutPipe()
	if err != nil {
		return
	}

	logger.Debugf("shell")
	err = session.Shell()
	if err != nil {
		return
	}

	logger.Debugf("xterm")
	err = session.RequestPty("xterm", 80, 40, ssh.TerminalModes{})
	if err != nil {
		return
	}

	return
}
