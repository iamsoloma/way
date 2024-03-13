package way

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Translate struct {}

func (Translate) BlockToLine(block Block) (line []byte) {
	line = []byte{}

	line = append(line, []byte(fmt.Sprint(block.ID))...)    // Block`s ID
	line = append(line, []byte("*/n*")...)                    // Splitter
	line = append(line, []byte(block.Time_UTC.String())...) // The time of the creation of the blockchain.
	line = append(line, []byte("*/n*")...)                    // Splitter
	line = append(line, block.PrevHash...)                  // Hash of Previous Block.
	line = append(line, []byte("*/n*")...)                    // Splitter
	line = append(line, block.Hash...)                      // Hash of Block
	line = append(line, []byte("*/n*")...)                    // Splitter
	line = append(line, block.Data...)                      // Data of Block

	return line
}

func (Translate) LineToBlock(line []byte) (block Block, err error) {
	lineSep := []byte("*/n*")
	time_form := "2006-01-02 15:04:05 Z0700 MST"

	content := bytes.Split(line, lineSep)
	//log.Printf("%x\n", content)
	//id
	block.ID, err = strconv.Atoi(string(content[0]))
	if err != nil {
		return block, errors.New("ID parsing error: " + err.Error())
	}
	//time
	block.Time_UTC, err = time.Parse(time_form, string(content[1]))
	if err != nil {
		return block, errors.New("Time parsing error: " + err.Error())
	}
	//block
	block.PrevHash = content[2]
	block.Hash = content[3]
	block.Data = content[4]

	return block, nil
}

func (Translate) FileToChain(exp *Explorer) (err error) {

	lBlock, err := exp.GetLastBlock()
	if err != nil {
		return errors.New("Error occurred when translating the file to chain: " + err.Error())
	}
	
	for i := 0; i <= lBlock.ID; i++ {
		curblock, err := exp.GetBlockByID(i)
		if err != nil {
			return errors.New("Error occurred when translating the file to chain: " + err.Error())
		}
		exp.Chain.blocks = append(exp.Chain.blocks, curblock)
	}

	return nil
}

func (Translate) ChainToFile(exp *Explorer) (file os.File, err error) {

	lBlock := exp.Chain.GetLastBlock()
	
	for i := 0; i <= lBlock.ID; i++ {
		curblock := exp.Chain.GetBlockByID(i)
		_, err := exp.AddBlock(curblock)
		if err != nil {
			return os.File{}, errors.New("Error occurred when translating the chain to file: " + err.Error())
		}
	}

	return 
}
