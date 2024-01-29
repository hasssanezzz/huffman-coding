# Huffman-coding - Java

This repository contains a straightforward implementation of Huffman coding in Java.

## Usage

### Compile the program

```
$ javac -d target src/main/java/com/huffmancoding/*.java
```

Once you have successfully compiled the program, utilize the following command-line arguments to compress and decompress files.

### File compression

```
$ java -cp target com.huffmancoding.Program <path/to/my/file.txt> <desired/output/path.bin>
```

Example:

```
$ java -cp target com.huffmancoding.Program file.txt compressed_data.bin
```

### File decompression

```
$ java -cp target com.huffmancoding.Program -d <path/to/my/file.bin> <desired/output/path.text>
```

Example:

```
$ java -cp target com.huffmancoding.Program -d compress_data.bin results.txt
```
