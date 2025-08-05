package com.dsa.recursion;

public class Rec {

    public int Add(int n){
        if(n == 1) {
            return 1;
        }
        return n + Add(n-1);
    }

    public int Factorial(int n) {

        if (n == 0) {
            return 1;
        }
        return n * Factorial(n-1);
    }
    
}


