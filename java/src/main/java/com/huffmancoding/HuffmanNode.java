package com.huffmancoding;

public class HuffmanNode {
    int freq;
    char data = '\0';
    HuffmanNode left = null;
    HuffmanNode right = null;

    public HuffmanNode(int freq, char data, HuffmanNode left, HuffmanNode right) {
        this.freq = freq;
        this.data = data;
        this.left = left;
        this.right = right;
    }
}