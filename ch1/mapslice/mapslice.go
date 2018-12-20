package main


import (
	"fmt"
)

func main() {

mapOfSlices  := make(map[string][]string)

attributeFile("first", "file1", mapOfSlices)
attributeFile("second", "file1", mapOfSlices)
attributeFile("second", "file2", mapOfSlices)
attributeFile("first", "file3", mapOfSlices)
attributeFile("third", "file4", mapOfSlices)


fmt.Println(mapOfSlices)

}

func attributeFile(line string, fname string, sliceMap map[string][]string) {
	sliceMap[line] = append(sliceMap[line], fname)
}
