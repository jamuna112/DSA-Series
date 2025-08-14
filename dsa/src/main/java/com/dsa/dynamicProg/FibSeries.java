package com.dsa.dynamicProg;

public class FibSeries {

    public int Fib(int num) {

        if (num == 0) {
            return 0;
        }
        if (num == 1){
            return 1;
        }

       int[] dp = new int[num + 1];
        dp[0] = 0;
        dp[1] = 1;

        for (int i=2; i <= num; i++) {
            dp[i] = dp[i-1] + dp[i-2];
            
        }
         
        return dp[num];
    }

     public int ClimbStair(int num) {

        int[] arr = new int[num + 1];

        arr[0] = 1;
        arr[1] = 1;

        for (int i = 2; i <= num; i++) {
            arr[i] = arr[i-1] + arr[i-2];
        }

        return arr[num];
    }

}
