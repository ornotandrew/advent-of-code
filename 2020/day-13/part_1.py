import math
import sys

with open("input.txt", "r") as f:
    lines = f.read().splitlines()

earliest_time = int(lines[0])
bus_ids = [int(bus) for bus in lines[1].split(",") if bus != "x"]


def find_first_multiple(number, greater_than):
    return math.ceil(greater_than / number) * number


bus_id, time_of_departure = min(((i, find_first_multiple(i, earliest_time)) for i in bus_ids), key=lambda t: t[1])
print(bus_id * (time_of_departure - earliest_time))
