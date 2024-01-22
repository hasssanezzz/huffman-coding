using System.Text;

namespace HuffmanTree
{
    public class HuffmanEncoder
    {
        public Dictionary<char, string> CharTable { get; set; }

        public HuffmanEncoder(string input, string filePath)
        {
            CharTable = new Dictionary<char, string>();

            Node root = BuildHuffmanTree(input);
            BuildCharTable(root, "");
            WriteToFile(filePath, input);
        }

        private Node BuildHuffmanTree(string input)
        {
            PriorityQueue<Node, int> pq = new();
            var frequencies = input.GroupBy(x => x).ToDictionary(x => x.Key, x => x.Count());

            foreach (var kvp in frequencies)
                pq.Enqueue(new Node { Symbol = kvp.Key, Frequency = kvp.Value }, kvp.Value);

            while (pq.Count > 1)
            {
                Node left = pq.Dequeue();
                Node right = pq.Dequeue();

                // combine the two children's frequencies
                int combinedFreq = left.Frequency + right.Frequency;
                Node parentNode = new() { Symbol = '\0', Frequency = combinedFreq, Left = left, Right = right };

                pq.Enqueue(parentNode, combinedFreq);
            }

            return pq.Peek();
        }

        private void BuildCharTable(Node node, string prefix)
        {
            if (node == null)
                return;

            if (!CharTable.ContainsKey(node.Symbol))
                CharTable.Add(node.Symbol, prefix);

            if (node.Left != null)
                BuildCharTable(node.Left, prefix + "0");
            if (node.Right != null)
                BuildCharTable(node.Right, prefix + "1");
        }

        private string GenerateBinaryString(string input)
        {
            StringBuilder output = new();

            for (int i = 0; i < input.Length; i++)
                output.Append(CharTable[input[i]]);

            return output.ToString();
        }

        private void WriteCharTable(BinaryWriter writer)
        {
            writer.Write(CharTable.Count);
            foreach (var kvp in CharTable)
            {
                char symbol = kvp.Key;
                string code = kvp.Value;

                writer.Write(symbol);
                writer.Write(code.Length);
                writer.Write(Encoding.UTF8.GetBytes(kvp.Value));
            }
        }

        private void WriteBinaryStringToFile(BinaryWriter writer, string binstr)
        {
            // get padding size and pad right
            int paddingLength = (8 - (binstr.Length % 8)) % 8;
            binstr = binstr.PadRight(binstr.Length + paddingLength, '0');

            // write data size and padding length
            writer.Write(binstr.Length / 8);
            writer.Write(paddingLength);

            for (int i = 0; i < binstr.Length; i += 8)
            {
                string eightBits = binstr.Substring(i, Math.Min(8, binstr.Length - i));
                writer.Write(Convert.ToByte(eightBits, 2));
            }
        }

        private void WriteToFile(string filePath, string input)
        {
            string binstr = GenerateBinaryString(input);


            using (BinaryWriter writer = new(new FileStream(filePath, FileMode.Create)))
            {
                WriteCharTable(writer);
                WriteBinaryStringToFile(writer, binstr);
            }
        }
    }
}
