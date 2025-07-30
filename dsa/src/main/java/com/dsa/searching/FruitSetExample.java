package com.dsa.searching;

import java.util.HashSet;
import java.util.Iterator;

public class FruitSetExample {
    
    public static HashSet<String> getFruitSet(){
        HashSet<String> fruits = new HashSet<>();

        // Add elements
        fruits.add("Apple");
        fruits.add("Banana");
        fruits.add("Mango");
        fruits.add("Apple"); // Duplicate, will be ignored

        // Check if an element exists
        if (fruits.contains("Banana")) {
            System.out.println("Banana is in the set");
        }

        // Remove an element
        fruits.remove("Apple");

        Iterator<String> it = fruits.iterator();

        while (it.hasNext()) {
            System.out.println(it.next());
        }

        // Return final set for testing
        return fruits;
    }
}
