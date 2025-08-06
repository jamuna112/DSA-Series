package com.dsa;

import java.util.HashSet;

import com.dsa.recursion.Rec;
import com.dsa.searching.FruitSetExample;
import com.dsa.searching.Fruitmap;

public class Main {
    public static void main(String[] args) {

        HashSet<String> fruits = FruitSetExample.getFruitSet();
        
        System.out.println("Final fruits set: "+ fruits);

        Fruitmap fm = new Fruitmap();
        fm.FruitAddition();

        Rec rec = new Rec();

        int result = rec.Add(6);
        int facResult = rec.Factorial(6);
        int febonacciSeries = rec.FebSeries(6);

        System.out.println("Sum of given numbers is "+ result );   
        System.out.println("Factorial of given numbers is "+ facResult );  
        System.out.println("Fibonacci term of given numbers 4 is  "+ febonacciSeries );  
    }
}  