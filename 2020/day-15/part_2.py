from part_1 import find_nth_number

with open("input.txt", "r") as f:
    numbers = [int(i) for i in f.read().split(",")]

print(find_nth_number(numbers, 30000000))
