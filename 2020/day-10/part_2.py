with open("input.txt", "r") as f:
    adaptors = sorted([int(n) for n in f.read().splitlines()])

joltages = [0] + adaptors + [max(adaptors) + 3]

total = 0
dp = [1]
for i in range(1, len(joltages) - 1):
    valid_combinations = 0
    for j in range(1, 4):
        if i - j < 0:
            continue
        if joltages[i] - joltages[i - j] <= 3:
            valid_combinations += dp[-j]

    dp.append(valid_combinations)

print(dp[-1])
