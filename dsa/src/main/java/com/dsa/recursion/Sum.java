package com.dsa.recursion;

public class Sum {

    public int Add(int n){
        if(n == 1) {
            return 1;
        }
        return n + Add(n-1);
    }
    
}
