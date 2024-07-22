package huffman

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Huffman struct {
	writer   io.Writer
	reader   io.Reader
	table    map[byte]string
	revTable map[string]byte
	data     []byte
}

func NewHuffman(writer io.Writer, reader io.Reader) *Huffman {
	return &Huffman{
		writer: writer,
		reader: reader,
	}
}

func (h *Huffman) Encode() error {
	h.table = map[byte]string{}

	start := time.Now()
	root, err := h.buildTree()
	if err != nil {
		return fmt.Errorf("encode function can not build tree: %v", err)
	}
	fmt.Printf("tree built in:\t%d\n", time.Since(start).Milliseconds())

	start = time.Now()
	h.buildCharTable(root, "")
	fmt.Printf("table built in:\t%d\n", time.Since(start).Milliseconds())

	start = time.Now()
	err = h.writeCharTable()
	if err != nil {
		return fmt.Errorf("can not write char table: %v", err)
	}
	fmt.Printf("table written:\t%d\n", time.Since(start).Milliseconds())

	start = time.Now()
	err = h.writeBinaryCodes()
	if err != nil {
		return fmt.Errorf("can not write binary codes: %v", err)
	}
	fmt.Printf("code written:\t%d\n", time.Since(start).Milliseconds())

	return nil
}

func (h *Huffman) buildTree() (*Node, error) {
	freq := map[byte]int{}

	// read all bytes
	bytes, err := ioutil.ReadAll(h.reader)
	if err != nil {
		return nil, err
	}
	h.data = bytes

	// count bytes
	for _, b := range bytes {
		freq[b]++
	}

	nodes := make([]Node, len(freq))
	pq, i := make(PriorityQueue, 0), 0

	// create and push initial nodes
	for b, freq := range freq {
		nodes[i] = Node{
			Freq:  freq,
			Byte:  b,
			Left:  nil,
			Right: nil,
		}

		heap.Push(&pq, &nodes[i])
		i++
	}

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		newNode := &Node{
			Byte:       0,
			Freq:       left.Freq + right.Freq,
			Left:       left,
			Right:      right,
			isInternal: true,
		}

		heap.Push(&pq, newNode)
	}

	return heap.Pop(&pq).(*Node), nil
}

func (h *Huffman) buildCharTable(root *Node, code string) {
	if root == nil {
		return
	}

	_, exists := h.table[root.Byte]
	if !exists && !root.isInternal {
		h.table[root.Byte] = code
	}

	if root.Left != nil {
		h.buildCharTable(root.Left, code+"1")
	}

	if root.Right != nil {
		h.buildCharTable(root.Right, code+"0")
	}
}

func (h *Huffman) writeCharTable() error {
	tableLen := uint(len(h.table))
	writer := bufio.NewWriter(h.writer)

	// write table length
	err := writer.WriteByte(byte(tableLen))
	if err != nil {
		return fmt.Errorf("can not write table length %d: %v", tableLen, err)
	}

	for b, code := range h.table {
		// write byte and code length
		_, err = writer.Write([]byte{b, byte(uint(len(code)))})
		if err != nil {
			return fmt.Errorf("can not write byte and code length: %v", err)
		}

		// write code as a string
		_, err = writer.Write([]byte(code))
		if err != nil {
			return fmt.Errorf("can not write code %s as string: %v", code, err)
		}
	}

	return nil
}

func (h *Huffman) writeBinaryCodes() error {
	writer := bufio.NewWriter(h.writer)

	var builder strings.Builder
	for _, b := range h.data {
		builder.WriteString(h.table[b])
	}

	// add padding
	paddingSize := (8 - builder.Len()%8) % 8
	for i := 0; i < paddingSize; i++ {
		builder.WriteRune('0')
	}

	err := writer.WriteByte(byte(uint(paddingSize)))
	if err != nil {
		return fmt.Errorf("can not write padding size as byte (%d): %v", paddingSize, err)
	}

	codes, length := builder.String(), builder.Len()
	for i := 0; i < length; i += 8 {
		bitsString := codes[i : i+8]

		b, err := strconv.ParseUint(bitsString, 2, 8)
		if err != nil {
			return fmt.Errorf("can not parse uint from sub string %q: %v", bitsString, err)
		}

		err = writer.WriteByte(byte(uint(b)))
		if err != nil {
			return fmt.Errorf("can not write byte from sub string (%q): %v", bitsString, err)
		}
	}

	return nil
}

func (h *Huffman) Decode() error {
	h.table = map[byte]string{}
	h.revTable = map[string]byte{}

	data, err := ioutil.ReadAll(h.reader)
	if err != nil {
		return fmt.Errorf("can not read all data from reader: %v", err)
	}
	reader := bytes.NewReader(data)

	// read table
	start := time.Now()
	err = h.readCharTable(reader)
	if err != nil {
		return fmt.Errorf("can read character table: %v", err)
	}
	fmt.Printf("char table read in:\t%d\n", time.Since(start).Milliseconds())

	// read binary codes
	start = time.Now()
	result, err := h.readBinaryCodes(reader)
	if err != nil {
		return fmt.Errorf("can not read binary codes: %v", err)
	}
	fmt.Printf("binary codes read in:\t%d\n", time.Since(start).Milliseconds())

	// write decompressed data
	start = time.Now()
	_, err = h.writer.Write(result)
	if err != nil {
		return fmt.Errorf("can not decode: %v", err)
	}
	fmt.Printf("results written in:\t%d\n", time.Since(start).Milliseconds())

	return nil
}

func (h *Huffman) readCharTable(reader *bytes.Reader) error {
	tableLengthAsByte, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("can not read table length: %v", err)
	}
	tableLength := uint(tableLengthAsByte)

	for i := 0; i < int(tableLength); i++ {
		currByte, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("can not read a byte: %v", err)
		}

		codeSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("can not read code size: %v", err)
		}

		codeStrBuffer := make([]byte, int(codeSize))
		_, err = reader.Read(codeStrBuffer)
		if err != nil {
			return fmt.Errorf("can not read code string: %v", err)
		}

		code := string(codeStrBuffer)
		h.table[currByte] = code
		h.revTable[code] = currByte
	}

	return nil
}

func (h *Huffman) readBinaryCodes(reader *bytes.Reader) ([]byte, error) {
	paddingSizeAsByte, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	var codes strings.Builder
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("can not read a byte from compressed file: %v", err)
		}

		codes.WriteString(fmt.Sprintf("%08b", b))
	}

	var result bytes.Buffer
	var currCode strings.Builder

	codesStr, length := codes.String(), codes.Len()-int(paddingSizeAsByte)
	for i := 0; i < length; i++ {
		currCode.WriteByte(codesStr[i])
		if b, ok := h.revTable[currCode.String()]; ok {
			result.WriteByte(b)
			currCode.Reset()
		}
	}

	return result.Bytes(), nil
}
