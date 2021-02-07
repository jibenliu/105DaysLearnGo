package main

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"os"
)

func main() {
	f, _ := os.OpenFile("./file", os.O_RDWR, 0644)
	defer f.Close()

	maps, _ := mmap.Map(f, mmap.RDWR, 0)
	defer maps.Unmap()
	fmt.Println(string(maps))

	maps[0] = 'x'
	maps.Flush()
}
