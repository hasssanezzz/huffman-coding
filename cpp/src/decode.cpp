#include "main.hpp"

std::string decode(std::string &file_name)
{
    std::ifstream input_file(file_name, std::ios::binary | std::ios::in);

    if (!input_file.is_open())
    {
        std::cerr << "Error opening file for writing: " << file_name << std::endl;
        exit(1);
    }

    // read padding size
    size_t padding_bits;
    input_file.read(reinterpret_cast<char *>(&padding_bits), sizeof(padding_bits));

    // read table size
    size_t table_size;
    input_file.read(reinterpret_cast<char *>(&table_size), sizeof(table_size));

    // read and build char table
    std::unordered_map<char, std::string> char_table;
    for (size_t i = 0; i < table_size; ++i)
    {
        char character;
        input_file.get(character);

        size_t code_size;
        input_file.read(reinterpret_cast<char *>(&code_size), sizeof(code_size));

        std::string code(code_size, '\0');
        input_file.read(&code[0], code_size);

        char_table[character] = code;
    }

    // read and create the binary string
    std::string encoded_data, decoded, current;
    char byte;
    while (input_file.get(byte))
    {
        for (int i = 7; i >= 0; --i)
        {
            char bit = ((byte >> i) & 1) ? '1' : '0';
            encoded_data += bit;
        }
    }

    input_file.close();

    // remove padding bits
    if (padding_bits > 0)
        encoded_data.resize(encoded_data.size() - padding_bits);

    // decode bits
    for (auto bit : encoded_data)
    {
        current += bit;

        for (const auto &entry : char_table)
        {
            if (entry.second == current)
            {
                decoded += entry.first;
                current.clear();
                break;
            }
        }
    }

    // TODO do something here
    return decoded;
}