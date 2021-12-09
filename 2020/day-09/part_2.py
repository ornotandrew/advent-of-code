with open("input.txt", "r") as f:
    numbers = [int(n) for n in f.read().splitlines()]

invalid_number = 31161678  # this matches my input specifically

left, right = 0, 0
while True:
    subset = numbers[left : right + 1]
    total = sum(subset)
    if total == invalid_number:
        print(min(subset) + max(subset))
        break
    elif total < invalid_number:
        # we need more numbers
        right += 1
    else:  # total > invalid_number
        # we need less numbers
        left += 1
