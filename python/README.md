# Huffman-coding

A simple implementation of Huffamn-coding written in python

## Usage

You can use the following CLI arguments to compress and decomress file.

To compress a file:
```
$python3 -m main <path/to/my/file> <desired/output/path.bin>

# Example:
$python3 -m main file.txt compressed_data.bin
```

To decompress a file:
```
$python3 -m main d <path/to/my/file.bin> <desired/output/path>

# Example:
$python3 -m main d compress_data.bin results.txt
```