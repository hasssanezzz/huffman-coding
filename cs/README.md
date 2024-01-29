# Huffman-coding - C#

This repository contains a straightforward implementation of Huffman coding in C#.

## Usage

Once you have successfully compiled the program, utilize the following command-line arguments to compress and decompress files.

### File compression

```
$ ./program <path/to/my/file.txt> <desired/output/path.bin>
```

Example:

```
$ ./program file.txt compressed_data.bin
```

### File decompression

```
$ ./program d <path/to/my/file.bin> <desired/output/path.text>
```

Example:

```
$ ./program d compress_data.bin results.txt
$ ./program decode compress_data.bin results.txt
```
