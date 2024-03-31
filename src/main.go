package main

import (
	"externalSort/src/dispatch"
	"externalSort/src/merge"
	"externalSort/src/sort"
)

func main() {
	dispatch.DispatchFileToChunks()
	sort.SortEveryChunk()
	merge.MergeChunks()
}
