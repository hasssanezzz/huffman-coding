#include <iostream>
#include <fstream>
#include <filesystem>
#include <queue>
#include <unordered_map>
#include <vector>
#include <bitset>

struct HuffmanNode
{
	char data;
	int frequency;
	HuffmanNode *left;
	HuffmanNode *right;

	HuffmanNode(char ch, int freq) : data(ch), frequency(freq), left(nullptr), right(nullptr) {}
};

struct CompareNodes
{
	bool operator()(const HuffmanNode *lhs, const HuffmanNode *rhs) const
	{
		return lhs->frequency > rhs->frequency;
	}
};

HuffmanNode *build_huffman_tree(const std::string &text);

void build_char_table(const HuffmanNode *node, std::unordered_map<char, std::string> &char_table, std::string code);

void write_encoded(const std::string &file_name, const std::string &text, std::unordered_map<char, std::string> &char_table);

void write_decoded(const std::string &file_name, const std::string &content);

std::string decode(std::string &file_name);