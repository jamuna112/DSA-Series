package com.dsa;

import java.util.HashSet;

import com.dsa.searching.FruitSetExample;
import com.dsa.searching.Fruitmap;

public class Main {
    public static void main(String[] args) {

        HashSet<String> fruits = FruitSetExample.getFruitSet();
        
        System.out.println("Final fruits set: "+ fruits);

        Fruitmap fm = new Fruitmap();
        fm.FruitAddition();
    }
}  