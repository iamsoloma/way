package way

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

func lineCounter(fullpath string) (int, error) {
	buf := make([]byte, 1*1024) //32 Kbyte
	count := 0
	lineSep := []byte{'\n'}

	file, err := os.OpenFile(fullpath, os.O_RDONLY, 0600)
	if err != nil {
		return count, errors.New("Error occurred when determining the last line of the file: " + err.Error())
	}

	defer file.Close()

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func GetLineByNum(r io.Reader, lineNum int) (line []byte, lastLine int, err error) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if lastLine == lineNum {
			return sc.Bytes(), lastLine, sc.Err()
		}
		lastLine++
	}
	if lastLine < lineNum {
		return line, lastLine, io.EOF
	} else {
		return line, lastLine, nil
	}
}