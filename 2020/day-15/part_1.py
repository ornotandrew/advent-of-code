def find_nth_number(starting_numbers, n):
    seen_at = {num: [i] for i, num in enumerate(starting_numbers)}
    previous_num = starting_numbers[-1]

    for i in range(len(starting_numbers), n):
        next_num = 0 if len(seen_at[previous_num]) == 1 else seen_at[previous_num][1] - seen_at[previous_num][0]

        if next_num not in seen_at:
            seen_at[next_num] = []
        elif len(seen_at[next_num]) == 2:
            seen_at[next_num].pop(0)  # free up some memory

        seen_at[next_num].append(i)

        previous_num = next_num

    return previous_num


with open("input.txt", "r") as f:
    numbers = [int(i) for i in f.read().split(",")]

print(find_nth_number(numbers, 2020))
