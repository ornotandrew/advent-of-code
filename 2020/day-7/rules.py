import re

with open("input.txt", "r") as f:
    lines = f.read().splitlines()


def parse_inner_bag(bag):
    num, color = bag.split(" ", 1)
    color = re.sub("s$", "", color)  # remove plurals
    return (color, int(num))


def parse_rule(line):
    outer, inner = line.split(" bags contain ")
    if inner == "no other bags.":
        return outer, None
    inner_bags = re.sub(" bags?\.?", "", inner).split(", ")
    return outer, dict(parse_inner_bag(bag) for bag in inner_bags)


rules = {}
for line in lines:
    outer, inner_rules = parse_rule(line)
    rules[outer] = inner_rules
