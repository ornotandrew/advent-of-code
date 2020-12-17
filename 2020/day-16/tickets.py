with open("input.txt", "r") as f:
    sections = f.read().split("\n\n")


def parse_rule(line):
    field, ranges = line.split(": ")
    ranges = [(int(a), int(b)) for a, b in [r.split("-") for r in ranges.split(" or ")]]
    return (field, ranges)


def parse_ticket(line):
    return [int(i) for i in line.split(",")]


rules = {field: ranges for field, ranges in [parse_rule(line) for line in sections[0].splitlines()]}
my_ticket = parse_ticket(sections[1].splitlines()[1])
nearby_tickets = [parse_ticket(line) for line in sections[2].splitlines()[1:]]


def is_valid_for_any_field(number):
    for ruleset in rules.values():
        for min_, max_ in ruleset:
            if min_ <= number <= max_:
                return True
    return False
