package com.dsa;

import java.util.HashSet;

import com.dsa.searching.FruitSetExample;

public class Main {
    public static void main(String[] args) {

        HashSet<String> fruits = FruitSetExample.getFruitSet();
        
        System.out.println("Final fruits set: "+ fruits);
    }
}  