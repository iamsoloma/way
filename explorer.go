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
	File *os.File
}

func (e Explorer) CreateBlockChain(genesis string, time_now_utc time.Time) error {
	/*var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(e.Path)
		if err != nil {
			return nil
		}
	} else {
		return errors.New("BlockChain is Exist! File: " + e.Path)
	}

	defer file.Close()*/

	b, err := Block.InitBlock(Block{}, []byte(genesis), time_now_utc)
	if err != nil {
		return err
	}

	_, err = e.File.Write(Translate.BlockToLine(Translate{}, b))
	if err != nil {
		return errors.New("Error occurred when writing the initialization block to " + e.Path + ": " + err.Error())
	}

	return nil
}

func (e Explorer) GetLastBlock() (lastBlock Block, err error) {
	/*var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return Block{}, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.Open(e.Path)
	if err != nil {
		return Block{}, err
	}

	defer file.Close()*/

	lastNumOfLine, err := lineCounter(e.File)
	if err != nil {
		return Block{}, errors.New("Error occurred when determining the last line of the file: " + err.Error())
	}
	line, _, err := GetLineByNum(e.File, lastNumOfLine)
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
	/*var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return Block{}, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.Open(e.Path)
	if err != nil {
		return Block{}, err
	}

	defer file.Close()*/

	line, _, err := GetLineByNum(e.File, id)
	if err != nil && err != io.EOF {
		return Block{}, err
	} else if err == io.EOF {
		return Block{}, errors.New("Error: the block with this ID does not exist: " + err.Error())
	}
	//fmt.Println(line)

	block, err = Translate.LineToBlock(Translate{}, line) //lineToBlock(line)
	if err != nil {
		return block, errors.New("Error: GetBlockByID: " + err.Error())
	}

	return block, nil
}

func (e Explorer) AddBlock(block Block) (id int, err error) {
	/*var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return 0, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.OpenFile(e.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return 0, err
	}

	defer file.Close()*/

	lastBlock, err := e.GetLastBlock()
	if err != nil {
		return lastBlock.ID+1, errors.New("Error occurred when determining the last Block in the file: " + err.Error())
	}

	block.ID = lastBlock.ID + 1

	line := Translate.BlockToLine(Translate{}, block)
	_, err = e.File.WriteString("\n" + string(line))
	if err != nil {
		return block.ID, errors.New("Error occurred when adding a block to the blockchain: " + err.Error())
	}

	return block.ID, nil
}


func lineCounter(f *os.File) (int, error) {
	buf := make([]byte, 1*1024) //32 Kbyte
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := f.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func GetLineByNum(f *os.File, lineNum int) (line []byte, lastLine int, err error) {
	sc := bufio.NewScanner(f)
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
