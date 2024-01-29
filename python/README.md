# Huffman-coding - Python3

This repository contains a straightforward implementation of Huffman coding in Python3.

## Usage

### File compression

```
$ python3 main.py <path/to/my/file.txt> <desired/output/path.bin>
```

Example:

```
$ python3 main.py file.txt compressed_data.bin
```

### File decompression

```
$ python3 main.py -d <path/to/my/file.bin> <desired/output/path.text>
```

Example:

```
$ python3 main.py -d compress_data.bin results.txt
```
