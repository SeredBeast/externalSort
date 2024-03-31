package storage

var FILE_PATH_TO_SORT = "source.txt"
var CHUNKS_DIR = "chunks"

// 1 kilobyte = 1000 bytes
// 1 megabyte = 1_000_000 bytes
// 1 gigabyte = 1_000_000_000 bytes
var ONE_KILOBYTE uint64 = 1000
var ONE_MEGABYTE uint64 = ONE_KILOBYTE * 1000
var ONE_GIGABYTE uint64 = ONE_MEGABYTE * 1000

var CHUNK_SIZE_BYTES = ONE_MEGABYTE * 500

// var CHUNK_SIZE_BYTES = ONE_MEGABYTE * 500
