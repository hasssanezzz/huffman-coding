package com.huffmancoding;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.HashMap;

public class HuffamnDecoder {
    public String readBinaryData(DataInputStream dataInputStream) throws IOException {
        StringBuilder binaryData = new StringBuilder();

        int paddingSize = dataInputStream.readInt();
        int byteCount = dataInputStream.readInt();

        for (int i = 0; i < byteCount; i++) {
            byte b = dataInputStream.readByte();
            String binaryString = Integer.toBinaryString(b & 0xFF);

            while (binaryString.length() < 8)
                binaryString = "0" + binaryString;

            binaryData.append(binaryString);
        }

        binaryData.setLength(binaryData.length() - paddingSize);

        return binaryData.toString();
    }

    public HashMap<String, Character> readTable(DataInputStream dataInputStream) throws IOException {
        HashMap<String, Character> charTable = new HashMap<>();

        int tableSize = dataInputStream.readInt();

        for (int i = 0; i < tableSize; i++) {
            String data = dataInputStream.readUTF();
            dataInputStream.readInt();
            String code = dataInputStream.readUTF();
            charTable.put(code, data.charAt(0));
        }

        return charTable;
    }

    public void decodeAndWriteBinaryData(String binaryData, HashMap<String, Character> charTable, String outFilepath) {
        try (DataOutputStream dataOutputStream = new DataOutputStream(new FileOutputStream(outFilepath))) {
            StringBuilder sb = new StringBuilder();

            for (int i = 0; i < binaryData.length(); i++) {
                sb.append(binaryData.charAt(i));

                if (charTable.containsKey(sb.toString())) {
                    dataOutputStream.writeChar(charTable.get(sb.toString()));
                    sb.setLength(0);
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public void decode(String filepath, String outFilepath) throws IOException {
        DataInputStream dataInputStream = new DataInputStream(new FileInputStream(filepath));
        HashMap<String, Character> charTable = readTable(dataInputStream);
        String binaryData = readBinaryData(dataInputStream);
        decodeAndWriteBinaryData(binaryData, charTable, outFilepath);
    }
}
