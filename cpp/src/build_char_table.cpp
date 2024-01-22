#include "main.hpp"

void build_char_table(const HuffmanNode *node, std::unordered_map<char, std::string> &char_table, std::string code)
{
	if (node)
	{
		if (node->data != '\0')
			char_table[node->data] = code;

		build_char_table(node->left, char_table, code + "0");
		build_char_table(node->right, char_table, code + "1");
	}
}
