package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hasssanezzz/huffman-coding/go/huffman"
)

func main() {

	mode := os.Args[1]
	input := os.Args[2]
	output := os.Args[3]

	reader, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	writer, err := os.OpenFile(output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	huff := huffman.NewHuffman(writer, reader)

	if mode == "e" {
		err := huff.Encode()
		if err != nil {
			log.Fatalf("fatal in main: %v\n", err)
		}

		readerInfo, err := reader.Stat()
		if err != nil {
			log.Fatalf("can not read input file info: %v", err)
		}

		writeInfo, err := os.Stat(output)
		if err != nil {
			log.Fatalf("can not read output file info: %v", err)
		}

		ratio := (float32(readerInfo.Size()) / float32(writeInfo.Size()))
		fmt.Printf("Original file size:\t%d\nCompressed file size:\t%d\nCompression ratio:\t%.3f", readerInfo.Size(), writeInfo.Size(), ratio)
	} else {
		err := huff.Decode()
		if err != nil {
			log.Fatalf("fatal in main: %v\n", err)
		}
	}
}
