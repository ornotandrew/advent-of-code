from math import prod

with open("input.txt", "r") as f:
    lines = f.read().splitlines()

bus_indexes = [(i, int(bus)) for i, bus in enumerate(lines[1].split(",")) if bus != "x"]
N = prod(m for i, m in bus_indexes)

# This is a direct implementation of the CRT
# https://brilliant.org/wiki/chinese-remainder-theorem/
#
# Note: to find the a_i coefficients, we "shift" the sequence of n_i (bus ids)
# by the "offset" of bus_i in the set of all busses. We do this because we want
# the sequence to land in the correct place (mod n_i).

# fmt: off
print(
    sum( # ∑ for all i
        (n_i - offset) # a_i
        * N // n_i # y_i
        * pow(
            N // n_i, # y_i
            -1, # NOTE: this needs python 3.8
            n_i
        ) # z_i
        for offset, n_i in bus_indexes
    ) % N
)
# fmt: on
