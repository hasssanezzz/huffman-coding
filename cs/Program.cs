using System.IO;

namespace HuffmanTree
{

    internal class Program
    {
        static void Main(string[] args)
        {
            if (args.Length < 2 || (args.Length > 3 && args[1] == "-d") || (args.Length > 2 && args[1] == "-d"))
            {
                Console.WriteLine("Invalid CLI arguments");
                return;
            }


            try
            {
                if (args[0] == "d" || args[0] == "decode")
                {
                    _ = new HuffmanDecoder(args[1], args[2]);

                    FileInfo compressedFile = new(args[1]);
                    FileInfo originalFile = new(args[2]);

                    Console.WriteLine("Original file size: " + originalFile.Length);
                    Console.WriteLine("Compressed file size: " + compressedFile.Length);
                }
                else
                {
                    _ = new HuffmanEncoder(File.ReadAllText(args[0]), args[1]);

                    FileInfo originalFile = new(args[0]);
                    FileInfo compressedFile = new(args[1]);

                    double compressionRate = (1 - compressedFile.Length / (double)originalFile.Length) * 100;

                    Console.WriteLine("Original file size: " + originalFile.Length);
                    Console.WriteLine("Compressed file size: " + compressedFile.Length);
                    Console.WriteLine($"Compression rate: {compressionRate:F2}");
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e.Message);
            }
        }
    }

}
