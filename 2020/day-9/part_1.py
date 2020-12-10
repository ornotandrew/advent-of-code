with open("input.txt", "r") as f:
    numbers = [int(i) for i in f.read().splitlines()]


def breakdown(target, possible_numbers):
    for n in possible_numbers:
        if (target - n) in possible_numbers:
            return n, target - n
    return None


# This can be done without brute-force by keeping track of all the possible
# sums as we move our index. I've done this the simple way, because this is
# still O(25 * n) ~= O(n) i.e. it's not too bad.
for i in range(25, len(numbers) - 25):
    if not breakdown(numbers[i], numbers[i - 25 : i]):
        print(numbers[i])
        break
