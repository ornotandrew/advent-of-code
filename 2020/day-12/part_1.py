def parse_line(line):
    return line[0], int(line[1:])


with open("input.txt", "r") as f:
    lines = [parse_line(l) for l in f.read().splitlines()]

directions_clockwise = ["N", "E", "S", "W"]
facing = "E"
ns, ew = 0, 0

for instruction, number in lines:
    if instruction == "N":
        ns += number
    elif instruction == "S":
        ns -= number
    elif instruction == "E":
        ew += number
    elif instruction == "W":
        ew -= number
    elif instruction == "R":
        times_to_rotate = (number / 90) % 4
        new_direction_index = (directions_clockwise.index(facing) + times_to_rotate) % 4
        facing = directions_clockwise[int(new_direction_index)]
    elif instruction == "L":
        times_to_rotate = (number / 90) % 4
        new_direction_index = (directions_clockwise.index(facing) - times_to_rotate) % 4
        facing = directions_clockwise[int(new_direction_index)]
    elif instruction == "F":
        if facing == "N":
            ns += number
        elif facing == "S":
            ns -= number
        elif facing == "E":
            ew += number
        elif facing == "W":
            ew -= number

print(abs(ns) + abs(ew))
