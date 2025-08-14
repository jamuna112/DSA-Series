package com.dsa;

import java.util.HashSet;
import java.util.List;

import com.dsa.dynamicProg.FibSeries;
import com.dsa.recursion.Backtracking;
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
        int[] array = {4, 5, 12, 7, 8};
        int target = 19;

        int result = rec.Add(6);
        int facResult = rec.Factorial(6);
        int febonacciSeries = rec.FebSeries(6);
        Boolean isSortArray = rec.isSorted(array, array.length);
        int index = rec.search(array, target);


        System.out.println("Sum of given numbers is "+ result );   
        System.out.println("Factorial of given numbers is "+ facResult );  
        System.out.println("Fibonacci term of given numbers 4 is  "+ febonacciSeries );  
        System.out.printf("It is %b that given array is sorted.\n", isSortArray ); 
         System.out.printf("Target %d value found in index %d\n", target, index ); 

         int[] arr1 = {1, 2, 3};

         Backtracking bt = new Backtracking();
        List<List<Integer>> list =  bt.subsets(arr1);

        System.out.println(list);

        FibSeries fs = new FibSeries();
        int num = 8;
        int res = fs.Fib(num);
        System.out.printf("Fib series of given %d num is %d\n", num, res);
         
        int steps = 5;
        int output = fs.ClimbStair(steps) ;
        System.out.printf("There are %d ways to climb %d steps\n", output, steps);
    }
}  