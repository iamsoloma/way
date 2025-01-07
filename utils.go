package way

import(
	"fmt"
)


func FullPath (path string, name string, part int) string {
	return path + "/" + name + ".hearing" + "/" + fmt.Sprint(part) + ".prt"
}

func BlockChainPath (path string, name string) string {
	return path + "/" + name + ".hearing"
}
