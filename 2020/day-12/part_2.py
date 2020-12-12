import math


def parse_line(line):
    return line[0], int(line[1:])


with open("input.txt", "r") as f:
    lines = [parse_line(l) for l in f.read().splitlines()]


def rotate(coords, degrees):
    x, y = coords
    r = math.radians(degrees)
    return x * math.cos(r) - y * math.sin(r), x * math.sin(r) + y * math.cos(r)


y, x = 1, 10
y_ship, x_ship = 0, 0


for instruction, number in lines:
    if instruction == "N":
        y += number
    elif instruction == "S":
        y -= number
    elif instruction == "E":
        x += number
    elif instruction == "W":
        x -= number
    elif instruction == "R":
        x, y = rotate((x, y), -number)
    elif instruction == "L":
        x, y = rotate((x, y), number)
    elif instruction == "F":
        x_ship += number * x
        y_ship += number * y

print(round(abs(x_ship) + abs(y_ship)))
