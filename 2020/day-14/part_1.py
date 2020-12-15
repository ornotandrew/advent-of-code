import re

with open("input.txt", "r") as f:
    lines = f.read().splitlines()

parse_write = lambda line: (int(i) for i in re.match("mem\[()\] = (\d+)", line).groups())
char_mask = lambda char, mask: "".join(["1" if x == char else "0" for x in mask])

memory = {}
to_add, to_subtract = 0, 0

for line in lines:
    if line.startswith("mask"):
        mask = line.split(" ")[-1]
        to_add = int(char_mask("1", mask), 2)
        to_subtract = int(char_mask("0", mask), 2)
        continue

    # otherwise, we're writing to memory
    address, value = parse_write(line)
    # the middle term here subtracts the cases where both the mask AND the
    # input are true (we don't want to double-count)
    memory[address] = value + (to_add - (to_add & value)) - (to_subtract & value)

print(sum(memory.values()))
