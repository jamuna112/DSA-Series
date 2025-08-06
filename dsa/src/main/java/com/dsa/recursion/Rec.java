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

    public int FebSeries(int n){
        //time complexity = O(2^n)
        if (n == 0 || n == 1) {
            return n;
        }

        return FebSeries(n-1) + FebSeries(n-2);
    }

    public Boolean isSorted(int[] arr, int n) {

        if (n == 0 || n == 1) {
            return true;
        }

        return arr[n-1] >= arr[n-2] && isSorted(arr, n - 1);
    }
    
}


