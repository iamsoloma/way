package way

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"strconv"
	"time"
	//"github.com/TinajXD/way"
)

type Explorer struct {
	Path string
}

func (e Explorer) CreateBlockChain(genesis string, time_now_utc time.Time) error {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(e.Path)
		if err != nil {
			return nil
		}
	} else {
		return errors.New("BlockChain is Exist! File: " + e.Path)
	}

	defer file.Close()

	b, err := Block.InitBlock(Block{}, []byte(genesis))
	if err != nil {
		return err
	}

	file.Write(blockToLine(0, b, time_now_utc))

	return nil
}

func (e Explorer) GetBlockByID(id int64) (block Block, lastID int64, time_utc time.Time, err error) {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return Block{}, 0, time_utc, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.Open(e.Path)
	if err != nil {
		return Block{}, 0, time_utc, err
	}

	defer file.Close()

	line, lastID, err := ReadLine(file, id)
	if err != nil && err != io.EOF {
		return Block{}, lastID, time_utc, err
	} else if err == io.EOF {
		return Block{}, lastID, time_utc, errors.New("Error: the block with this ID does not exist: " + err.Error())
	}

	lastID, block, time_utc, err = lineToBlock(line)
	if err != nil {
		return block, lastID, time_utc, errors.New("Error: GetBlockByID: " + err.Error())
	}

	return block, lastID, time_utc, nil
}

// TODO: Sync in memory blockchain and file
func (e Explorer) AddBlock(block Block, time_utc time.Time) (id int64, err error) {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		return 0, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	file, err = os.OpenFile(e.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	lastID, err := lineCounter(e.Path)
	if err != nil {
		return lastID, errors.New("Error occurred when determining the last line of the file: " + err.Error())
	}
	//log.Println("LastID: " + fmt.Sprint(lastID))

	line := blockToLine(lastID+1, block, time_utc)
	_, err = file.WriteString("\n" + string(line))
	if err != nil {
		return lastID, errors.New("Error occurred when adding a block to the blockchain: " + err.Error())
	}

	return 0, nil
}

// TODO: Sync in memory blockchain and file

func blockToLine(id int64, block Block, time_utc time.Time) (line []byte) {
	line = []byte{}

	line = append(line, []byte(strconv.FormatInt(id, 10))...) // Block`s ID
	line = append(line, []byte("/n")...)                      // Splitter
	line = append(line, []byte(time_utc.String())...)         // The time of the creation of the blockchain.
	line = append(line, []byte("/n")...)                      // Splitter
	line = append(line, block.PrevHash...)                    // Hash of Previous Block.
	line = append(line, []byte("/n")...)                      // Splitter
	line = append(line, block.Hash...)                        // Hash of Block
	line = append(line, []byte("/n")...)                      // Splitter
	line = append(line, block.Data...)                        // Data of Block

	return line
}

func lineToBlock(line []byte) (id int64, block Block, time_utc time.Time, err error) {
	lineSep := []byte("/n")
	time_form := "2006-01-02 15:04:05 Z0700 MST"

	content := bytes.Split(line, lineSep)
	//log.Printf("%x\n", content)
	//id
	buf := bytes.NewBuffer(content[0])
	binary.Read(buf, binary.BigEndian, id)
	//time
	time_utc, err = time.Parse(time_form, string(content[1]))
	if err != nil {
		return id, block, time_utc, errors.New("Time parsing error: " + err.Error())
	}
	//block
	block.PrevHash = content[2]
	block.Hash = content[3]
	block.Data = content[4]

	return id, block, time_utc, nil
}

func lineCounter(path string /*, r io.Reader*/) (int64, error) {
	buf := make([]byte, 32*1024) //32 Kbyte
	var count int64 = 0
	lineSep := []byte{'\n'}

	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return count, errors.New("Error occurred when determining the last line of the file: " + err.Error())
	}

	for {
		c, err := file.Read(buf)
		count += int64(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func ReadLine(r io.Reader, lineNum int64) (line []byte, lastLine int64, err error) {
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
