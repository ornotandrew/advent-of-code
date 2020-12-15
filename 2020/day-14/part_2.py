import re

with open("input.txt", "r") as f:
    lines = f.read().splitlines()

parse_write = lambda line: (int(i) for i in re.match("mem\[(\d+)\] = (\d+)", line).groups())
set_char_at_pos = lambda string, char, pos: string[:pos] + char + string[pos + 1 :]


def merge_chars(address, mask):
    if mask == "X":
        return "X"
    return address if mask == "0" else mask


def explode_addresses(address, mask):
    address_bin = bin(address).replace("0b", "").rjust(36, "0")
    initial_address = "".join([merge_chars(a, m) for a, m in zip(address_bin, mask)])
    addresses = [initial_address]
    for i, char in enumerate(initial_address):
        if char == "X":
            addresses = [set_char_at_pos(m, "1", i) for m in addresses] + [
                set_char_at_pos(m, "0", i) for m in addresses
            ]
            string_addresses = "\n".join(addresses)
    return [int(a, 2) for a in addresses]


memory = {}
mask = None

for line in lines:
    if line.startswith("mask"):
        mask = line.split(" ")[-1]
        continue

    # otherwise, we're writing to memory
    address, value = parse_write(line)
    addresses = explode_addresses(address, mask)
    for a in addresses:
        memory[a] = value

print(sum(memory.values()))
