def parse_line(raw_line):
    policy, string = raw_line.split(": ")
    num_times, letter = policy.split(" ")
    min_times, max_times = [int(i) for i in num_times.split("-")]
    return (min_times, max_times, letter, string)


def is_valid(line):
    min_times, max_times, letter, string = line
    return min_times <= string.count(letter) <= max_times


with open("input.txt", "r") as f:
    raw_lines = f.read().splitlines()

lines = [parse_line(l) for l in raw_lines]

print([is_valid(l) for l in lines].count(True))
