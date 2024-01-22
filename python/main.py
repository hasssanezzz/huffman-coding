import sys
from os.path import getsize
from huffman_coding import HuffmanCoding


def main():
    args = sys.argv

    if len(args) < 3 or (len(args) > 4 and args[1] == '-d') or (len(args) > 3 and args[1] != '-d'):
        print("Invalid arguments")
        return

    if args[1] == '-d':
        HuffmanCoding.decode_file(args[2], args[3])
    else:
        HuffmanCoding.encode_file(args[1], args[2])
        file_size, compressed_file_size = getsize(args[1]), getsize(args[2])
        print("Original file size:", file_size)
        print("compressed file size:", compressed_file_size)
        print("Compression ratio:", (1 - compressed_file_size / file_size) * 100)


if __name__ == "__main__":
    main()
