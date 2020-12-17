from tickets import nearby_tickets, is_valid_for_any_field


invalid_values = []

for field_values in nearby_tickets:
    invalid_values += [val for val in field_values if not is_valid_for_any_field(val)]

print(sum(invalid_values))
