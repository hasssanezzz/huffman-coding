# Go-huffman

## Usage

1. After cloning the repo, build the executable
    ```
    go build -o main
    ```
1. Compress a file
    ```
    $ ./main e <input_file_path> <output>
    
    Example:
    $ ./main e mybigfile.txt bigfile.bin
    ```
1. Decompress a file a file
    ```
    $ ./main d <binary_file_path> <output>
    
    Example:
    $ ./main d bigfile.bin mybigfile.txt
    ```
### Example

```
$ .\main.exe e D:\Coding\Work\index.html out.bin

Original file size:     20318
Compressed file size:   11366
Compression ratio:      1.788
```