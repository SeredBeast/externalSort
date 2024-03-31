package merge

import (
	"externalSort/src/storage"
	"externalSort/src/utils"
	"fmt"
	"os"
	"runtime"
)

var chunkCounter = 0
var chunkGen = 0

func MergeChunks() {
	for !utils.RemainsOneFile("sorted_chunks") {
		MergeChunksGen()
	}
}

func MergeChunksGen() {
	filenames := utils.IteratorDirFiles("sorted_chunks")
	chunkGen++
	for {
		chunkName1, ok := <-filenames
		if !ok {
			break
		}
		chunkName2, ok := <-filenames
		if !ok {
			break
		}
		sortMerge(chunkName1, chunkName2)
		mergedChunkBytes = []byte{}

		utils.DeleteFile("sorted_chunks/" + chunkName1)
		utils.DeleteFile("sorted_chunks/" + chunkName2)
	}
}

var mergedChunkBytes = []byte{}

func bufferWrite(file *os.File, line string) {
	MAX_SIZE := storage.CHUNK_SIZE_BYTES

	curLen := len(mergedChunkBytes)

	lineBytes := []byte(line)
	lineLen := len(lineBytes)

	if (curLen + lineLen) > int(MAX_SIZE) {
		file.Write(mergedChunkBytes)
		mergedChunkBytes = []byte{}
		mergedChunkBytes = append(mergedChunkBytes, lineBytes...)
		return
	}

	mergedChunkBytes = append(mergedChunkBytes, lineBytes...)
}

func bufferFlush(file *os.File) {
	file.Write(mergedChunkBytes)
	go runtime.GC()
	mergedChunkBytes = []byte{}
}

func sortMerge(chunkName1 string, chunkName2 string) {
	mergedChunk := createChunkFile()

	chunkSc1 := utils.OpenFileWithScanner("sorted_chunks/" + chunkName1)
	chunkSc2 := utils.OpenFileWithScanner("sorted_chunks/" + chunkName2)
	chunkLines1 := utils.CreateFileLinesIterator(chunkSc1)
	chunkLines2 := utils.CreateFileLinesIterator(chunkSc2)

	line1, ok := <-chunkLines1
	if !ok {
		writeAllRemaning(mergedChunk, chunkLines2)
		bufferFlush(mergedChunk)
		return
	}

	line2, ok := <-chunkLines2
	if !ok {
		bufferWrite(mergedChunk, line1+"\n")
		writeAllRemaning(mergedChunk, chunkLines1)
		bufferFlush(mergedChunk)
		return
	}

	for {
		if line1 < line2 {
			bufferWrite(mergedChunk, line1+"\n")
			line1, ok = <-chunkLines1
			if !ok {
				bufferWrite(mergedChunk, line2+"\n")
				writeAllRemaning(mergedChunk, chunkLines2)
				break
			}
			continue
		}

		if line1 > line2 {
			bufferWrite(mergedChunk, line2+"\n")
			line2, ok = <-chunkLines2
			if !ok {
				bufferWrite(mergedChunk, line1+"\n")
				writeAllRemaning(mergedChunk, chunkLines1)
				break
			}
			continue
		}

		bufferWrite(mergedChunk, line1+"\n")
		bufferWrite(mergedChunk, line2+"\n")

		line1, ok = <-chunkLines1
		if !ok {
			// writer <- line2 + "\n")
			writeAllRemaning(mergedChunk, chunkLines2)
			break
		}
		line2, ok = <-chunkLines2
		if !ok {
			bufferWrite(mergedChunk, line1+"\n")
			writeAllRemaning(mergedChunk, chunkLines1)
			break
		}
	}

	bufferFlush(mergedChunk)
}

func writeAllRemaning(chunk *os.File, lines chan string) {
	for line := range lines {
		bufferWrite(chunk, line+"\n")
	}
}

func createChunkFile() *os.File {
	chunkCounter += 1
	chunkName := fmt.Sprintf("%d_%d_chunk.txt", chunkCounter, chunkGen)
	chunkPath := fmt.Sprintf("sorted_chunks/%s", chunkName)
	chunkFile, _ := os.Create(chunkPath)
	fmt.Printf("merged chunk -- %s\n", chunkName)
	return chunkFile
}
