package utils

import (
	"bufio"
	"externalSort/src/storage"
	"fmt"
	"os"
	"strings"

	"github.com/twotwotwo/sorts/sortutil"
)

func CreateFileLinesIterator(scanner *bufio.Scanner) chan string {
	lines := make(chan string, storage.ONE_MEGABYTE*25)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
				close(lines)
				return
			}

			lines <- line
		}

		fmt.Println("readed")
		close(lines)
	}()

	return lines
}

func OpenFileWithScanner(filename string) *bufio.Scanner {
	file, _ := os.Open(filename)
	reader := bufio.NewScanner(file)
	return reader
}

func CreateDir(dirname string) {
	os.Mkdir(dirname, os.ModePerm)
}

func RemainsOneFile(dirname string) bool {
	entries, _ := os.ReadDir(dirname)
	return len(entries) == 1
}

func IteratorDirFiles(dirname string) chan string {
	entries, _ := os.ReadDir(dirname)
	filenames := make(chan string)

	go func() {
		for _, entry := range entries {
			filenames <- entry.Name()
		}
		close(filenames)
	}()

	return filenames
}

func WriteArrayToFile(filedata []string, chunk *os.File) {
	for _, line := range filedata {
		chunk.WriteString(line + "\n")
	}
}

func FileToArray(lines chan string) []string {
	filedata := []string{}
	for line := range lines {
		filedata = append(filedata, line)
	}
	return filedata
}

func SortLines(filedata []string) {
	sortutil.Strings(filedata)
}

func DeleteFile(filename string) {
	os.Remove(filename)
}

func CreateFileReader(filename string) (*os.File, *bufio.Reader) {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	return file, reader
}

func ReadWholeFile(filename string) []byte {
	bytes, _ := os.ReadFile(filename)
	return bytes
}

func ReadFromFileWithSize(file *os.File) ([]byte, error) {
	bytes := make([]byte, storage.CHUNK_SIZE_BYTES)

	record, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}

	bytes = bytes[:record]
	lastByte := bytes[record-1]

	remains := []byte{}
	for lastByte != '\n' && lastByte != '\x00' {
		single := make([]byte, 1)
		_, err := file.Read(single)
		if err != nil {
			break
		}
		remains = append(remains, single...)
		lastByte = single[0]
	}

	bytes = append(bytes, remains...)

	return bytes, nil
}

func ReadFromFileStrWithSize(file *os.File) ([]string, error) {
	bytes, err := ReadFromFileWithSize(file)
	if err != nil {
		return nil, err
	}
	return BytesToStrArr(bytes), nil
}

func BytesToStrArr(bytes []byte) []string {
	chunkStr := string(bytes)
	chunkArr := strings.Split(chunkStr, "\n")
	return chunkArr
}

func StrArrToBytes(strArr []string) []byte {
	str := strings.Join(strArr, "\n")
	bytes := []byte(str)
	bytes = append(bytes, '\n')
	return bytes
}

func ValidateFile(filename string) {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	prevLine := scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		if prevLine > line {
			fmt.Printf("ERROR: %s\n", filename)
			fmt.Println(prevLine)
			fmt.Println(line)
		}
		prevLine = line
	}
}
