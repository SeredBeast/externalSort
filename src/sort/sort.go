package sort

import (
	"externalSort/src/utils"
	"fmt"
	"os"
)

func dropEmptyString(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	if arr[0] == "" || arr[0] == "\n" {
		arr = arr[1:]
	}
	return arr
}

func SortEveryChunk() {
	utils.CreateDir("sorted_chunks")
	filenames := utils.IteratorDirFiles("chunks")

	for filename := range filenames {
		sortChunk(filename)
		utils.DeleteFile("chunks/" + filename)
	}

	fmt.Println("done sorting")

	// entries, _ := os.ReadDir("sorted_chunks")
	// for _, entry := range entries {
	// 	checkName := entry.Name()
	// 	utils.ValidateFile("sorted_chunks/" + checkName)
	// }

	// fmt.Println("done checking")
}

func sortChunk(filename string) {
	chunkData := utils.ReadWholeFile("chunks/" + filename)
	chunkArr := utils.BytesToStrArr(chunkData)

	utils.SortLines(chunkArr)
	chunkArr = dropEmptyString(chunkArr)
	chunkData = utils.StrArrToBytes(chunkArr)

	sortedChunk := createSortedChunkFile(filename)
	sortedChunk.Write(chunkData)
}

func createSortedChunkFile(chunkName string) *os.File {
	chunkPath := fmt.Sprintf("sorted_chunks/%s", chunkName)
	chunkFile, _ := os.Create(chunkPath)
	fmt.Printf("created sorted chunk -- %s\n", chunkName)
	return chunkFile
}
