with open("input.txt", "r") as f:
    adaptors = sorted([int(n) for n in f.read().splitlines()])

joltages = [0] + adaptors + [max(adaptors) + 3]
differences = [b - a for a, b in zip(joltages, joltages[1:])]
print(differences.count(1) * differences.count(3))
