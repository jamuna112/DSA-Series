package com.dsa.recursion;

import java.util.ArrayList;
import java.util.List;

public class Backtracking {
    private void backtrack(int index, int[] nums, List<Integer> current, List<List<Integer>> result) {
        // Add the current subset to the result list
        result.add(new ArrayList<>(current));

        // Explore further elements
        for (int i = index; i < nums.length; i++) {
            // Choose the current number
            current.add(nums[i]);

            // Explore with the chosen number
            backtrack(i + 1, nums, current, result);

            // Backtrack: remove the last added number
            current.remove(current.size() - 1);
        }
    }

    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(0, nums, new ArrayList<>(), result);
        return result;
    }
}
