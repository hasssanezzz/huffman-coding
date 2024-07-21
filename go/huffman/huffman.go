package huffman

import (
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
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
	root, err := h.buildTree()
	if err != nil {
		return fmt.Errorf("encode function can not build tree: %v", err)
	}

	h.buildCharTable(root, "")
	h.writeCharTable()
	return h.writeBinaryCodes()
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

func (h *Huffman) writeCharTable() {
	tableLen := uint(len(h.table))
	var buff bytes.Buffer

	// write table length
	buff.WriteByte(byte(tableLen))

	for b, code := range h.table {
		// write byte and code length
		buff.Write([]byte{b, byte(uint(len(code)))})
		// write code as a string
		buff.Write([]byte(code))
	}

	h.writer.Write(buff.Bytes())
}

func (h *Huffman) writeBinaryCodes() error {
	var buff bytes.Buffer

	var builder strings.Builder
	for _, b := range h.data {
		builder.WriteString(h.table[b])
	}

	// add padding
	paddingSize := (8 - builder.Len()%8) % 8
	for i := 0; i < paddingSize; i++ {
		builder.WriteRune('0')
	}

	buff.WriteByte(byte(uint(paddingSize)))
	codes, length := builder.String(), builder.Len()
	for i := 0; i < length; i += 8 {
		substr := codes[i : i+8]

		b, err := strconv.ParseUint(substr, 2, 8)
		if err != nil {
			return fmt.Errorf("can not parse uint from sub string %q: %v", substr, err)
		}

		buff.WriteByte(byte(uint(b)))
	}

	_, err := h.writer.Write(buff.Bytes())
	if err != nil {
		return fmt.Errorf("can not write binary codes: %v", err)
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

	err = h.readCharTable(reader)
	if err != nil {
		return fmt.Errorf("can read character table: %v", err)
	}

	result, err := h.readBinaryCodes(reader)
	if err != nil {
		return fmt.Errorf("can not read binary codes: %v", err)
	}

	_, err = h.writer.Write(result)
	return fmt.Errorf("can not decode: %v", err)
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
			return nil, fmt.Errorf("err reading a byte from compressed file: %v", err)
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
