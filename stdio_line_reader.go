package network

import (
	"io"
	"os"
)

func NextStdioLine() (string, error) {
	s := make([]byte, 0)
	buff := make([]byte, 8)
	for {
		n, err := os.Stdin.Read(buff)
		if err != nil {
			if err == io.EOF {
				s = append(s, buff...)
				break
			}
			return string(s), err
		}
		if n == 0 {
			break
		}
		s = append(s, buff...)
	}
	return string(s), nil
}
