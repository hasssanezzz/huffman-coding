#include "main.hpp"

void write_encoded(const std::string &file_name, const std::string &text, std::unordered_map<char, std::string> &char_table)
{
	std::fstream output_file(file_name, std::ios::binary | std::ios::out);

	if (!output_file.is_open())
	{
		std::cerr << "Error opening file for writing: " << file_name << std::endl;
		exit(1);
	}

	// generate the binary string
	std::string binstr;
	for (auto c : text)
		binstr += char_table[c];

	// get needed padding size and write it
	size_t padding_bits = (8 - binstr.size() % 8) % 8;
	output_file.write(reinterpret_cast<const char *>(&padding_bits), sizeof(padding_bits));
	binstr += std::string(padding_bits, '0');

	// Write character table size
	size_t table_size = char_table.size();
	output_file.write(reinterpret_cast<const char *>(&table_size), sizeof(table_size));

	// Write character table
	for (const auto &entry : char_table)
	{
		output_file.put(entry.first);
		size_t codeSize = entry.second.size();
		output_file.write(reinterpret_cast<const char *>(&codeSize), sizeof(codeSize));
		output_file.write(entry.second.c_str(), codeSize);
	}

	// write binary data
	for (int i = 0; i < binstr.size(); i += 8)
	{
		std::bitset<8> byte(binstr.substr(i, 8));
		char byteChar = static_cast<char>(byte.to_ulong());
		output_file.put(byteChar);
	}

	output_file.close();
}