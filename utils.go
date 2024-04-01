package way

import(
	"fmt"
)


func FullPath (path string, name string, part int) string {
	return path + "/" + name + ".bc" + "/" + fmt.Sprint(part) + ".prt"
}

func BlockChainPath (path string, name string) string {
	return path + "/" + name + ".bc"
}
