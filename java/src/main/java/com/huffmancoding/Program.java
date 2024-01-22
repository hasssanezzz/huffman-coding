package com.huffmancoding;

import java.io.IOException;

class Program {
    public static void main(String[] args) {

        if (args.length < 2 || (args.length > 3 && args[1].equals("-d")) || (args.length > 2 && args[1].equals("-d"))) {
            System.err.println("Invalid CLI arguments");
            return;
        }

        try {

            if (args[0].equals("-d")) {
                HuffamnDecoder decoder = new HuffamnDecoder();
                decoder.decode(args[1], args[2]);
            } else {
                HuffmanEncoder encoder = new HuffmanEncoder(args[0], args[1]);
                encoder.encode();
            }            
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}