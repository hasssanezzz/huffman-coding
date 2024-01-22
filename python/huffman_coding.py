import struct
from heapq import heapify, heappop, heappush
from collections import Counter


class Node:
    def __init__(self, chr: str, freq: int, left: 'Node' = None, right: 'Node' = None):
        self.symbol = chr
        self.freq = freq
        self.left = left
        self.right = right

    def __lt__(self, other: 'Node'):
        return self.freq < other.freq


class HuffmanCoding:

    @staticmethod
    def generate_tree(text: str) -> Node:
        if not text:
            return

        char_count = Counter(text)
        nodes = [Node(k, v) for k, v in char_count.items()]
        heapify(nodes)

        while len(nodes) > 1:
            left, right = heappop(nodes), heappop(nodes)
            newfreq = left.freq + right.freq
            heappush(nodes, Node(None, newfreq, left, right))

        return heappop(nodes)

    @staticmethod
    def build_char_table(root: Node, table: dict[str, str] = {}, code: str = ""):
        if not root:
            return

        if root.symbol and root.symbol not in table:
            table[root.symbol] = code

        if root.left:
            HuffmanCoding.build_char_table(root.left, table, code + '0')
        if root.right:
            HuffmanCoding.build_char_table(root.right, table, code + '1')

    @staticmethod
    def encode_data(text: str, char_table: dict[str, str]) -> str:
        return ''.join([char_table[c] for c in text])

    @staticmethod
    def write_char_table(char_table: dict[str, str], file):
        file.write(struct.pack('i', len(char_table)))
        for char in char_table:
            code = char_table[char]
            # write character
            file.write(char.encode())
            # write code size
            file.write(struct.pack('i', len(code)))
            # write code
            file.write(code.encode())

    @staticmethod
    def write_encoded_data(data: str, file):
        binary_string = HuffmanCoding.encode_data(data)
        file.write(struct.pack('i', len(binary_string)))
        file.write(binary_string)

    @staticmethod
    def encode_file(infilepath: str, filepath: str):
        text = open(infilepath, 'r').read()

        char_table = {}
        root = HuffmanCoding.generate_tree(text)
        HuffmanCoding.build_char_table(root, char_table)

        with open(filepath, 'wb') as file:
            HuffmanCoding.write_char_table(char_table, file)

            binary_string = HuffmanCoding.encode_data(text, char_table)
            padding_size = (8 - (len(binary_string) % 8)) % 8
            binary_string += padding_size * '0'
            # write data size
            file.write(struct.pack('i', len(binary_string)))
            # write padding size
            file.write(struct.pack('i', padding_size))
            file.write(int(binary_string, 2).to_bytes(
                (len(binary_string) + 7) // 8, byteorder='big'))

    @staticmethod
    def read_char_table(file) -> dict[str, str]:
        char_table: dict[str, str] = {}
        table_size = struct.unpack('<I', file.read(4))[0]

        for _ in range(table_size):
            # read sybmol
            char = file.read(1).decode('utf-8')
            # read code size
            code_size = struct.unpack('<I', file.read(4))[0]
            # read code
            code = file.read(code_size).decode('utf-8')

            # NOTE: reversed code and character for faster search
            char_table[code] = char

        return char_table

    def read_binary_data(file) -> str:
        data_size = struct.unpack('<I', file.read(4))[0]
        padding_size = struct.unpack('<I', file.read(4))[0]
        binary_data = ''.join(format(byte, '08b')
                              for byte in file.read(data_size))[:-padding_size]

        return binary_data

    def decode_binary_data(binary_data: str, char_table: dict[str, str]):
        curr, result = "", ""

        for c in binary_data:
            curr += c
            if curr in char_table:
                result += char_table[curr]
                curr = ""

        return result

    def decode_file(infilepath: str, filepath: str):
        with open(infilepath, 'rb') as file:
            char_table = HuffmanCoding.read_char_table(file)
            binary_data = HuffmanCoding.read_binary_data(file)
            data = HuffmanCoding.decode_binary_data(binary_data, char_table)

        try:
            open(filepath, 'w').write(data)
        except:
            print(f'Can not write decoded data to file {filepath}')
