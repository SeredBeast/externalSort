package test

import (
	"bufio"
	"externalSort/src/storage"
	"fmt"
	"os"
	"strings"
	"testing"
)

// doesn't care how lines ends
func TestReadMegabyte(t *testing.T) {
	// megabyte := make([]byte, storage.ONE_GIGABYTE*3)
	file, _ := os.Open("../source.txt")
	reader := bufio.NewReader(file)
	// recorded, _ := file.Read(megabyte)

	copy, _ := os.Create("../copy.txt")

	// megabyte := make([]byte, storage.ONE_GIGABYTE)
	// record, _ := file.Read(megabyte)
	// megabyte = megabyte[0:record]
	// copy.Write(megabyte)

	// megabyte = make([]byte, storage.ONE_GIGABYTE)
	// record, _ = file.Read(megabyte)
	// megabyte = megabyte[0:record]
	// copy.Write(megabyte)

	// megabyte = make([]byte, storage.ONE_GIGABYTE)
	// record, _ = file.Read(megabyte)
	// megabyte = megabyte[0:record]
	// copy.Write(megabyte)

	// megabyte = make([]byte, storage.ONE_GIGABYTE)
	// record, _ = file.Read(megabyte)
	// megabyte = megabyte[0:record]
	// copy.Write(megabyte)

	megabyte := make([]byte, storage.ONE_GIGABYTE)
	record, _ := reader.Read(megabyte)
	megabyte = megabyte[:record]
	copy.Write(megabyte)

	// megabyte = make([]byte, storage.ONE_GIGABYTE)
	bytes, _ := reader.ReadBytes('\n')
	// megabyte = megabyte[:record]
	copy.Write(bytes)

	megabyte = make([]byte, storage.ONE_GIGABYTE)
	record, _ = reader.Read(megabyte)
	megabyte = megabyte[:record]
	copy.Write(megabyte)

	megabyte = make([]byte, storage.ONE_GIGABYTE)
	record, _ = reader.Read(megabyte)
	megabyte = megabyte[:record]
	copy.Write(megabyte)

	file.Close()
	_ = 0
}

// file.Read() и reader.ReadBytes() скипают себе оффсета причем еще и друг для друга
func TestReadReader(t *testing.T) {
	// megabyte := make([]byte, storage.ONE_MEGABYTE*200)
	file, _ := os.Open("../source.txt")
	reader := bufio.NewReader(file)

	myBytes := make([]byte, 100)
	file.Read(myBytes)
	fmt.Println(myBytes)

	// myBytes = make([]byte, 100)
	// file.Read(myBytes)
	// fmt.Println(myBytes)

	bytes, _ := reader.ReadBytes('\n')
	// reader.Read()
	fmt.Println(bytes)

	// bytes, _ = reader.ReadBytes('\n')
	// fmt.Println(bytes)
	// fmt.Println(len(bytes))

	// offest := len(bytes) + 5
	// file.Seek(int64(offest), 1)
	// myBytes = make([]byte, 77)
	// file.Read(myBytes)
	// fmt.Println(myBytes)

	file.Close()
	_ = 0
}

func TestToString(t *testing.T) {
	file, _ := os.Open("../source.txt")
	reader := bufio.NewReader(file)

	kilobyte := make([]byte, storage.ONE_MEGABYTE*500)
	record, _ := reader.Read(kilobyte)
	kilobyte = kilobyte[:record]
	remain, _ := reader.ReadBytes('\n')
	kilobyte = append(kilobyte, remain...)

	str := string(kilobyte)
	kilobyte = nil
	strArr := strings.Split(str, "\n")
	str = ""
	// runtime.GC()
	fmt.Println(strArr)
}

func TestReadSmallFile(t *testing.T) {
	file, _ := os.Open("../small.txt")
	reader := bufio.NewReader(file)

	kilobyte := make([]byte, storage.ONE_MEGABYTE*500)
	record, err := reader.Read(kilobyte)
	fmt.Println(err)
	fmt.Println(record)
	kilobyte = make([]byte, storage.ONE_MEGABYTE*500)
	record, err = reader.Read(kilobyte)
	fmt.Println(err)
	fmt.Println(record)
}

func TestCut(t *testing.T) {
	bytes := []byte{'1', '1', '1', '0', '0'}
	bytes = bytes[:3]
	fmt.Println(bytes)
}
