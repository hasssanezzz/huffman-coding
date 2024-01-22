# Huffman-coding - CPP

A simple implementation of huffamn-coding written in CPP

## Usage

After you have compiled the program, you can this CLI arguments to compress and decomress file.

To compress a file:
```
$./program <path/to/my/file.txt> <desired/output/path.bin>

# Example:
$./program file.txt compressed_data.bin
```

To compress a file:
```
$./program -d <path/to/my/file.bin> <desired/output/path.text>

# Example:
$./program -d compress_data.bin results.txt
```

## Todos

- [ ] rebuilding the huffman tree when decoding
- [ ] use CLI argument parser