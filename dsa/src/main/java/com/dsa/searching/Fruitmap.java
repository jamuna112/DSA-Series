package com.dsa.searching;

import java.util.HashMap;

public class Fruitmap {

    public HashMap<String, Integer> FruitAddition(){

    HashMap<String, Integer> fm = new HashMap<>();
    
    fm.put("Mango", 345);
    fm.put("Apple", 123);
    fm.put("Orange", 345);

    System.out.println(fm);
    System.out.println(fm.get("Apple"));
    System.out.println(fm.containsKey("Mango"));
    System.out.println(fm.containsValue(845));

        return fm;
}
}
    
    
