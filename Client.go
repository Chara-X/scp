package scp

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Client struct{ *ssh.Client }

func (c *Client) CopyTo(file *os.File, remotePath string, permissions string) {
	var session, _ = c.NewSession()
	defer session.Close()
	var w, _ = session.StdinPipe()
	func() {
		defer w.Close()
		session.Start(fmt.Sprintf("scp -t %q", remotePath))
		var stat, _ = file.Stat()
		fmt.Fprintln(w, "C"+permissions, stat.Size(), path.Base(remotePath))
		io.Copy(w, file)
		w.Write([]byte{0})
	}()
	session.Wait()
}
func (c *Client) CopyFrom(file *os.File, remotePath string) {
	var session, _ = c.NewSession()
	defer session.Close()
	var r, _ = session.StdoutPipe()
	var w, _ = session.StdinPipe()
	session.Start(fmt.Sprintf("scp -f %q", remotePath))
	w.Write([]byte{0})
	var status = make([]byte, 1)
	r.Read(status)
	var header, _ = bufio.NewReader(r).ReadString('\n')
	var size, _ = strconv.Atoi(strings.Split(header, " ")[1])
	w.Write([]byte{0})
	io.CopyN(file, r, int64(size))
	w.Write([]byte{0})
	session.Wait()
}
