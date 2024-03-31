package dispatch

import (
	"externalSort/src/storage"
	"externalSort/src/utils"
	"fmt"
	"os"
)

var chunkCounter uint64 = 0

func DispatchFileToChunks() {
	utils.CreateDir("chunks")
	file, _ := os.Open(storage.FILE_PATH_TO_SORT)
	dispatchToChunks(file)
}

func dispatchToChunks(file *os.File) {
	for {
		bytes, err := utils.ReadFromFileWithSize(file)
		if err != nil {
			break
		}
		chunkFile := createChunkFile()
		chunkFile.Write(bytes)
	}
}

func createChunkFile() *os.File {
	chunkCounter += 1
	chunkName := fmt.Sprintf("%d_chunk.txt", chunkCounter)
	chunkPath := fmt.Sprintf("chunks/%s", chunkName)
	chunkFile, _ := os.Create(chunkPath)
	fmt.Printf("created chunk -- %s\n", chunkName)
	return chunkFile
}
