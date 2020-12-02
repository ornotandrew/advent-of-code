def parse_line(raw_line):
    policy, string = raw_line.split(": ")
    num_times, letter = policy.split(" ")
    position_1, position_2 = [int(i) for i in num_times.split("-")]
    return (position_1, position_2, letter, string)


def is_valid(line):
    position_1, position_2, letter, string = line
    return (string[position_1 - 1] == letter) ^ (string[position_2 - 1] == letter)


with open("input.txt", "r") as f:
    raw_lines = f.read().splitlines()

lines = [parse_line(l) for l in raw_lines]

print([is_valid(l) for l in lines].count(True))
