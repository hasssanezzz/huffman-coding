#include "main.hpp"

HuffmanNode *build_huffman_tree(const std::string &text)
{
	std::unordered_map<char, int> char_frequency;
	for (char ch : text)
		char_frequency[ch]++;

	std::priority_queue<HuffmanNode *, std::vector<HuffmanNode *>, CompareNodes> pq;

	for (const auto &entry : char_frequency)
	{
		HuffmanNode *node = new HuffmanNode(entry.first, entry.second);
		pq.push(node);
	}

	while (pq.size() > 1)
	{
		HuffmanNode *left = pq.top();
		pq.pop();
		HuffmanNode *right = pq.top();
		pq.pop();

		HuffmanNode *internal_node = new HuffmanNode('\0', left->frequency + right->frequency);
		internal_node->left = left;
		internal_node->right = right;

		pq.push(internal_node);
	}

	HuffmanNode *root = pq.top();
	pq.pop();

	return root;
}
