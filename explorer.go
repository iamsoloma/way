package way

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"time"
)

type Explorer struct {
	Path  string
	Chain Chain
}

func (e Explorer) CreateBlockChain(genesis string, time_now_utc time.Time) error {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(e.Path)
		if err != nil {
			return errors.New("Can`t create blockchain! Can`t create a file: " + e.Path)
		}
	} else {
		return errors.New("BlockChain is Exist! File: " + e.Path)
	}

	defer file.Close()

	b, err := Block.InitBlock(Block{}, []byte(genesis), time_now_utc)
	if err != nil {
		return err
	}

	file.Write(Translate.BlockToLine(Translate{}, b))

	return nil
}

func (e Explorer) GetLastBlock() (lastBlock Block, err error) {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return Block{}, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.Open(e.Path)
	if err != nil {
		return Block{}, err
	}

	defer file.Close()

	lastNumOfLine, err := lineCounter(e.Path)
	if err != nil {
		return Block{}, errors.New("Error occurred when determining the last line of the file: " + err.Error())
	}
	line, _, err := GetLineByNum(file, lastNumOfLine)
	if err != nil {
		return Block{}, errors.New("Error occurred when getting the last line of the file: " + err.Error())
	}
	lastBlock, err = Translate.LineToBlock(Translate{}, line)
	if err != nil {
		return Block{}, errors.New("Error occurred when translating the last line of the file: " + err.Error())
	}


	return lastBlock, nil
}

func (e Explorer) GetBlockByID(id int) (block Block, err error) {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return Block{}, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.Open(e.Path)
	if err != nil {
		return Block{}, err
	}

	defer file.Close()

	line, _, err := GetLineByNum(file, id)
	if err != nil && err != io.EOF {
		return Block{}, err
	} else if err == io.EOF {
		return Block{}, errors.New("Error: the block with this ID does not exist: " + err.Error())
	}

	block, err = Translate.LineToBlock(Translate{}, line) //lineToBlock(line)
	if err != nil {
		return block, errors.New("Error: GetBlockByID: " + err.Error())
	}

	return block, nil
}

func (e Explorer) AddBlock(block Block) (id int, err error) {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return 0, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.OpenFile(e.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	lastBlock, err := e.GetLastBlock()
	if err != nil {
		return lastBlock.ID+1, errors.New("Error occurred when determining the last Block in the file: " + err.Error())
	}

	block.ID = lastBlock.ID + 1

	line := Translate.BlockToLine(Translate{}, block)
	_, err = file.WriteString("\n" + string(line))
	if err != nil {
		return block.ID, errors.New("Error occurred when adding a block to the blockchain: " + err.Error())
	}

	return block.ID, nil
}


func lineCounter(path string /*, r io.Reader*/) (int, error) {
	buf := make([]byte, 1*1024) //32 Kbyte
	count := 0
	lineSep := []byte{'\n'}

	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return count, errors.New("Error occurred when determining the last line of the file: " + err.Error())
	}

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
