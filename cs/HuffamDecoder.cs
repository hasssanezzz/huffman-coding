using System.Text;

namespace HuffmanTree
{
    public class HuffmanDecoder
    {
        public Dictionary<string, char> CharTable { get; set; }

        public HuffmanDecoder(string filePath, string outputPath)
        {
            CharTable = new Dictionary<string, char>();

            using (FileStream fileStream = new(filePath, FileMode.Open))
            using (BinaryReader reader = new(fileStream, Encoding.UTF8))
            {
                ReadCharTableFromFile(reader);

                string binstr = ReadBinaryData(reader);
                string result = DecodeBinaryString(binstr);

                // write decoded data
                File.WriteAllText(outputPath, result);
            }
        }

        public void ReadCharTableFromFile(BinaryReader reader)
        {
            int tableCount = reader.ReadInt32();

            for (int i = 0; i < tableCount; i++)
            {
                char key = reader.ReadChar();
                int length = reader.ReadInt32();
                string value = Encoding.UTF8.GetString(reader.ReadBytes(length));

                // NOTE: reversed the key value pair for faster search
                CharTable[value] = key;
            }
        }

        public string ReadBinaryData(BinaryReader reader)
        {
            int binaryDataSize = reader.ReadInt32();
            int paddingLength = reader.ReadInt32();

            byte[] byteArr = reader.ReadBytes(binaryDataSize);

            StringBuilder binstr = new(byteArr.Length * 8);

            foreach (var b in byteArr)
                binstr.Append(Convert.ToString(b, 2).PadLeft(8, '0'));

            // remove padding zeros
            binstr.Length -= paddingLength;

            return binstr.ToString();
        }

        private string DecodeBinaryString(string binstr)
        {
            StringBuilder result = new();
            StringBuilder current = new();

            for (int i = 0; i < binstr.Length; i++)
            {
                current.Append(binstr[i]);

                if (CharTable.ContainsKey(current.ToString()))
                {
                    result.Append(CharTable[current.ToString()]);
                    current.Clear();
                }
            }

            return result.ToString();
        }
    }
}
