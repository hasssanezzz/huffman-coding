#include "main.hpp"

int main(int argc, char *argv[])
{

	std::vector<std::string> args;
	for (int i = 0; i < argc; ++i)
		args.push_back(argv[i]);

	if (argc < 3 || (argc > 4 && args[1] == "-d") || (argc > 3 && args[1] != "-d"))
	{
		std::cout << "CLI args error" << '\n';
		return 0;
	}

	// deocde
	if (args[1] == "-d")
	{
		std::string in_filepath = args[2];
		std::string out_filepath = args[3];

		std::string content = decode(in_filepath);
		write_decoded(out_filepath, content);
		// encode
	}
	else
	{
		std::string in_filepath = args[1];
		std::string out_filepath = args[2];
		std::unordered_map<char, std::string> char_table;

		std::fstream filecontent(in_filepath);
		std::string text((std::istreambuf_iterator<char>(filecontent)), std::istreambuf_iterator<char>());

		HuffmanNode *root = build_huffman_tree(text);
		build_char_table(root, char_table, "");
		write_encoded(out_filepath, text, char_table);

		uintmax_t og_file_size = std::filesystem::file_size(in_filepath);
		uintmax_t compressed_file_size = std::filesystem::file_size(in_filepath);
		double compressionRate = (1 - compressed_file_size / (double)og_file_size) * 100;

		std::cout << "Original file size: " << og_file_size << " bytes" << '\n';
		std::cout << "Encoded file size:  " << compressed_file_size << " bytes" << '\n';
		std::cout << "Compression rate: " << compressionRate << '\n';
	}

	return 0;
}