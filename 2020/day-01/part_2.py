from itertools import combinations

with open("input.txt", "r") as f:
    arr = [int(i) for i in f.readlines()]

arr = sorted([i for i in arr if i < 2020])

for a, b, c in combinations(arr, 3):
    if a + b + c == 2020:
        print(a * b * c)
        break
