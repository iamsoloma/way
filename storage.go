package way

import (
	"errors"
	"os"
	"time"

	//"github.com/TinajXD/way"
)

type Explorer struct {
	Path string
}



func (e Explorer) CreateBlockChain(genesis string) (error) {
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

	line := []byte{}
	line = append(line, []byte("0")...)
	line = append(line, []byte(" ")...)
	line = append(line, b.Hash...)
	line = append(line, []byte(" ")...)
	line = append(line, []byte(time.Now().UTC().String())...)
	

	file.Write(line)

	return nil
}

// TODO: Sync in memory blockchain and file
func (e Explorer) AddBlock(chain Chain, time_utc time.Time) (id int64, err error) {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(e.Path)
		if err != nil {
			return 0, nil
		}
	} else {
		file, err = os.Open(e.Path)
		if err != nil {
			return 0, nil
		}
	}

	defer file.Close()



	return 0, nil
}

// TODO: Sync in memory blockchain and file
func (e Explorer) SyncBlockChain(chain Chain) (id int64, err error) {
	var file *os.File
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(e.Path)
		if err != nil {
			return 0, nil
		}
	} else {
		file, err = os.Open(e.Path)
		if err != nil {
			return 0, nil
		}
	}

	defer file.Close()



	return 0, nil
}