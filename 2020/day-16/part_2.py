import math

from tickets import is_valid_for_any_field, my_ticket, nearby_tickets, rules

valid_nearby_tickets = [t for t in nearby_tickets if all(is_valid_for_any_field(value) for value in t)]
valid_tickets = valid_nearby_tickets + [my_ticket]

possible_positions_for_field = {field: list(range(len(my_ticket))) for field in rules.keys()}


def invalid_positions_for_ruleset(ruleset, ticket):
    invalid_positions = set()
    for min_, max_ in ruleset:
        for position, value in enumerate(ticket):
            if not (min_ <= value <= max_):
                # this field cannot possibly be in this position
                invalid_positions.add(position)
    return invalid_positions


def is_valid_for_rule(ruleset, value):
    (min_a, max_a), (min_b, max_b) = ruleset
    return any(min_ <= value <= max_ for min_, max_ in ruleset)


# 1. Use the ticket values to narrow down the options
for field, ruleset in rules.items():
    for ticket in valid_tickets:
        for position, value in enumerate(ticket):
            if position not in possible_positions_for_field[field]:
                # we've already determined that this combination is impossible
                continue

            if not is_valid_for_rule(ruleset, value):
                # this field cannot possibly be in this position
                possible_positions_for_field[field].remove(position)
                break


# 2. There should now be at least one position with only one option. Use that
# to continue removing possible options from the other fields.
to_remove = [x[0] for x in possible_positions_for_field.values() if len(x) == 1]
while to_remove:
    value = to_remove.pop(0)
    for field, possible_positions in possible_positions_for_field.items():
        if len(possible_positions) == 1:
            continue

        if value in possible_positions:
            possible_positions.remove(value)

        if len(possible_positions) == 1:
            to_remove.append(possible_positions[0])


for field, possible_positions in possible_positions_for_field.items():
    assert len(possible_positions) == 1

field_positions = {field: possible_positions[0] for field, possible_positions in possible_positions_for_field.items()}

print(
    math.prod(
        [my_ticket[pos] for pos in [pos for field, pos in field_positions.items() if field.startswith("departure")]]
    )
)
