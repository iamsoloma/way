package way

import (
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Explorer struct {
	Path      string
	Name      string
	Part      int
	Partition int
	Chain     Chain
}

func (e Explorer) CreateBlockChain(genesis string, time_now_utc time.Time) error {
	var file *os.File
	e.Part = 0
	if _, err := os.Stat(FullPath(e.Path, e.Name, e.Part)); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(BlockChainPath(e.Path, e.Name), 0764)
		if err != nil {
			return errors.New("Can`t create a workspace! Can`t create a path: " + e.Path + "\n" + err.Error())
		}
		file, err = os.Create(FullPath(e.Path, e.Name, e.Part))
		if err != nil {
			return errors.New("Can`t create blockchain! Can`t create a file: " + FullPath(e.Path, e.Name, e.Part))
		}
	} else if os.IsExist(err) {
		return errors.New("BlockChain is Exist! File: " + e.Path)
	}

	defer file.Close()

	b := Block{}
	err := b.InitBlock([]byte(genesis), time_now_utc)
	if err != nil {
		return err
	}

	file.Write(Translate.BlockToLine(Translate{}, b))

	return nil
}

func (e Explorer) DeleteBlockChain() (found bool, err error) {
	fp := BlockChainPath(e.Path, e.Name)

	if _, err := os.Stat(fp); err == nil {
		err = os.RemoveAll(fp)
		if err != nil {
			return true, errors.New("Can`t remove blockchain: " + err.Error())
		}
	} else if os.IsNotExist(err) {
		return false, errors.New("Blockchain is not found:" + err.Error())
	}
	return true, nil
}

func (e Explorer) GetListOfParts() (nums []int, err error) {
	var dir *os.File
	bcp := BlockChainPath(e.Path, e.Name)
	if _, err := os.Stat(bcp); errors.Is(err, os.ErrNotExist) {
		return nums, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	if dir, err = os.Open(bcp); err != nil {
		return nums, errors.New("Can`t open blockchain directory: " + err.Error())
	}

	parts, err := dir.Readdirnames(-1)
	if err != nil {
		return nums, errors.New("Can`t Read blockchain directory: " + err.Error())
	}

	for i := len(parts) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(strings.ReplaceAll(parts[i], ".prt", ""))
		if err != nil {
			return nums, errors.New("Can`t translate part name to the int: " + err.Error())
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)

	return nums, nil
}

func (e Explorer) GetLastBlock() (lastBlock Block, err error) {
	var file *os.File
	if _, err := os.Stat(BlockChainPath(e.Path, e.Name)); errors.Is(err, os.ErrNotExist) {
		return Block{}, errors.New("BlockChain is NOT Exist! A file is required: " + e.Path)
	}

	parts, err := e.GetListOfParts()
	if err != nil {
		return Block{}, err
	}
	e.Part = parts[len(parts)-1]

	file, err = os.Open(FullPath(e.Path, e.Name, e.Part))
	if err != nil {
		return Block{}, err
	}

	defer file.Close()

	lastNumOfLine, err := LineCounter(FullPath(e.Path, e.Name, e.Part))
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
	mod := id / e.Partition
	e.Part = mod

	//println(e.Part, id)
	lineNum := id - (e.Partition * mod)
	//lineNum := id
	if _, err := os.Stat(FullPath(e.Path, e.Name, e.Part)); errors.Is(err, os.ErrNotExist) {
		return Block{}, errors.New("BlockChainPart is NOT Exist! A file is required: " + FullPath(e.Path, e.Name, e.Part))
	}

	file, err = os.Open(FullPath(e.Path, e.Name, e.Part))
	if err != nil {
		return Block{}, err
	}

	defer file.Close()

	line, _, err := GetLineByNum(file, lineNum)
	if err != nil && err != io.EOF {
		return Block{}, err
	} else if err == io.EOF {
		return Block{}, errors.New("Error: the block with this ID does not exist: " + err.Error())
	}

	block, err = Translate.LineToBlock(Translate{}, line) //lineToBlock(line)
	if err != nil {
		return block, errors.New("Error: Can`t translate file line to block:" + err.Error())
	}

	return block, nil
}

func (e Explorer) AddBlock(data []byte, time_now_utc time.Time) (num int, id int, err error) {
	if _, err := os.Stat(BlockChainPath(e.Path, e.Name)); os.IsNotExist(err) {
		return 0, 0, errors.New("Add Block: BlockChain is NOT Exist! A file is required: " + BlockChainPath(e.Path, e.Name))
	}
	nums, err := e.GetListOfParts()
	if err != nil {
		return 0, 0, errors.New("Add Block: " + err.Error())
	}
	e.Part = nums[len(nums)-1]

	lastBlock, err := e.GetLastBlock()
	if err != nil {
		return e.Part, lastBlock.ID, errors.New("Error occurred when determining the last Block in the file: " + err.Error())
	}

	//println(lastBlock.ID+1, (lastBlock.ID+1)%e.Partition == 0)
	if (lastBlock.ID+1)%e.Partition == 0 {
		e.Part += 1
		e.Part, id, err = e.addBlockInNewPart(data, time_now_utc, lastBlock)
	} else {
		e.Part, id, err = e.addBlockInCurrentPart(data, time_now_utc, lastBlock)
	}
	if err != nil {
		return e.Part, lastBlock.ID, err
	}

	return e.Part, id, nil
}

func (e Explorer) addBlockInNewPart(data []byte, time_now_utc time.Time, lastblock Block) (num int, id int, err error) {
	var file *os.File
	if _, err := os.Stat(FullPath(e.Path, e.Name, e.Part)); os.IsNotExist(err) {
		file, err = os.Create(FullPath(e.Path, e.Name, e.Part))
		if err != nil {
			return e.Part, lastblock.ID, errors.New("Can`t create blockchain Part! Can`t create a file: " + FullPath(e.Path, e.Name, e.Part))
		}
	} else if os.IsExist(err) {
		return e.Part, lastblock.ID, errors.New("BlockChain Part is Exist! File: " + e.Path)
	}

	defer file.Close()

	b := Block{}
	b.NewBlock(data, lastblock, time_now_utc)
	if err != nil {
		return e.Part, b.ID, err
	}

	line := Translate.BlockToLine(Translate{}, b)
	_, err = file.Write(line)
	if err != nil {
		return e.Part, b.ID, errors.New("Error occurred when adding a block to the new part of blockchain: " + err.Error())
	}

	return e.Part, b.ID, nil
}

func (e Explorer) addBlockInCurrentPart(data []byte, time_utc time.Time, lastblock Block) (num int, id int, err error) {
	var file *os.File
	if _, err := os.Stat(FullPath(e.Path, e.Name, e.Part)); errors.Is(err, os.ErrNotExist) {
		return 0, 0, errors.New("BlockChainPart is NOT Exist! A file is required: " + FullPath(e.Path, e.Name, e.Part))
	}
	file, err = os.OpenFile(FullPath(e.Path, e.Name, e.Part), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return e.Part, 0, errors.New("Add Block: Can`t open file:" + err.Error())
	}

	defer file.Close()

	nBlock := Block{}
	nBlock.NewBlock(data, lastblock, time_utc)

	line := Translate.BlockToLine(Translate{}, nBlock)
	_, err = file.WriteString("\n" + string(line))
	if err != nil {
		return e.Part, nBlock.ID, errors.New("Error occurred when adding a block to the current part of blockchain: " + err.Error())
	}
	return e.Part, nBlock.ID, nil
}
