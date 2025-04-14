Title: Find All Pairs with a Given Sum

Objective: Given an array of integers and a target sum, find all unique pairs of integers in the array that add up to the target sum. The solution should maintain a time complexity of O(n2)O(n2) or better.
Input:
An array of integers, arr.
An integer, target_sum.
Output:
A list of tuples, where each tuple contains a pair of integers from the array that add up to the target sum.
Constraints:
The array can contain both positive and negative integers.
The same element cannot be used twice in a pair.
The pairs should be unique, i.e., (a, b) and (b, a) are considered the same and should be counted only once.
Example
Input:
arr = [1, 2, 3, 4, 3, 5, -1, 0]target_sum = 4
Output:
[(1, 3), (4, 0), (5, -1)]

Input:
arr = [1, 2, 3, 4, 3, 5, -1, 0]
target_sum = 4

Output:
[(1, 3), (4, 0), (5, -1)]