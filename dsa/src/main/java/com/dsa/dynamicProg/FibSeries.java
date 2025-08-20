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

     public int rob(int[] nums) {
         
        if (nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }

        int[] array = new int[nums.length];

        array[0] = nums[0];
        array[1] = Math.max(nums[0], nums[1]);

        for (int i = 2; i < nums.length; i++) {
            array[i] = Math.max(array[i-1], nums[i] + array[i-2]);
        }
        return array[nums.length - 1];
     }

     public int upStairs(int num) {

        if (num == 0 || num == 1) {
            return 1;
        }

        if (num == 2) {
            return 2;
        }

        int[] dp = new int[num + 1];
        dp[0] = 1;
        dp[1] = 1;
        dp[2] = 2; 

        for (int i=3; i<=num; i++) {
            dp[i] = dp[i-1]+dp[i-2]+dp[i-3];
        }
        return dp[num];
     }
}
