package com.huffmancoding;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.io.IOException;
import java.io.DataOutputStream;
import java.io.FileOutputStream;
import java.util.PriorityQueue;
import java.util.HashMap;
import java.util.Comparator;

public class HuffmanEncoder {

    private String filepath;
    private String text;
    private HashMap<Character, String> charTable;

    HuffmanEncoder(String infilepath, String outfilepath) throws IOException {
        this.text = Files.readString(Paths.get(infilepath));
        this.filepath = outfilepath;
        charTable = new HashMap<>();
    }

    private HuffmanNode buildTree() {
        HashMap<Character, Integer> mp = new HashMap<>();

        for (int i = 0; i < text.length(); i++)
            mp.put(text.charAt(i), mp.getOrDefault(text.charAt(i), 0) + 1);

        PriorityQueue<HuffmanNode> pq = new PriorityQueue<>(Comparator.comparingInt(c -> c.freq));

        // push the chars
        mp.forEach((data, freq) -> {
            pq.add(new HuffmanNode(freq, data, null, null));
        });

        // do the magic
        while (pq.size() > 1) {
            HuffmanNode left = pq.poll();
            HuffmanNode right = pq.poll();
            pq.add(new HuffmanNode(left.freq + right.freq, '\0', left, right));
        }

        return pq.peek();
    }

    private void buildCharTable(HuffmanNode root, String code) {
        if (root != null) {
            if (root.data != '\0')
                charTable.put(root.data, code);

            buildCharTable(root.left, code + "0");
            buildCharTable(root.right, code + "1");
        }
    }

    private void writeCharTable() throws IOException {
        int tableSize = charTable.size();

        try (DataOutputStream dataOutputStream = new DataOutputStream(new FileOutputStream(filepath))) {

            dataOutputStream.writeInt(tableSize);
            for (char data : charTable.keySet()) {
                String code = charTable.get(data);
                dataOutputStream.writeUTF(Character.toString(data));
                dataOutputStream.writeInt(code.length());
                dataOutputStream.writeUTF(code);
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void writeBinaryData(String binstr) throws IOException {
        // get paddding size
        int paddingSize = (8 - binstr.length() % 8) % 8;

        // add padding
        StringBuilder sb = new StringBuilder(binstr);
        for (int i = 0; i < paddingSize; i++)
            sb.append('0');
        binstr = sb.toString();

        try (DataOutputStream dataOutputStream = new DataOutputStream(new FileOutputStream(filepath, true))) {
            dataOutputStream.writeInt(paddingSize);
            dataOutputStream.writeInt(binstr.length() / 8);
            // create byte array
            for (int i = 0; i < binstr.length(); i += 8) {
                String byteString = binstr.substring(i, i + 8);
                byte b = (byte) Integer.parseInt(byteString, 2);
                dataOutputStream.writeByte(b);
            }

        }
    }

    private String encodeText() {
        StringBuilder sb = new StringBuilder();

        for (int i = 0; i < text.length(); i++)
            sb.append(charTable.get(text.charAt(i)));

        return sb.toString();
    }

    public void encode() throws IOException {
        HuffmanNode tree = buildTree();
        buildCharTable(tree, "");
        writeCharTable();
        writeBinaryData(encodeText());
    }
}