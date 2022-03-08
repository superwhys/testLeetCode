#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# @File    ：main.py
# @IDE     ：PyCharm 
# @Author  ：SuperYong
# @Date    ：2022/2/14 12:23 
# @Summary : this is the summary 


def nextGreaterElement(nums2):
    res = {}
    stack = []
    for num in reversed(nums2):
        while stack and num >= stack[-1]:
            stack.pop()
        res[num] = stack[-1] if stack else -1
        stack.append(num)
        print(stack)
    print(res)


def twoSum(num, target):
    num_len = len(num)
    for idx in range(num_len):
        l, r = idx + 1, num_len - 1
        while l <= r:
            mid = l + (r - l) // 2
            if num[mid] + num[idx] == target:
                return [idx+1, mid+1]
            elif num[mid] + num[idx] > target:
                # in left
                r = mid - 1
            else:
                l = mid + 1

        return [-1, -1]


def twoSum2(num, target):
    m = {}
    for idx, item in enumerate(num):
        if target - item in m:
            return [m[target-item], idx+1]
        m[num[idx]] = idx+1
    return [-1, -1]


def twoSum3(numbers, target):
    for i in range(len(numbers)):
        for j in range(i+1, len(numbers)):
            if numbers[i] + numbers[j] == target:
                return [i+1, j+1]


print(twoSum3([2, 7, 11, 15], 9))
'''
              l      m       r
              l,m  r
'''