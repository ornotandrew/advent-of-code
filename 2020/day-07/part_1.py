from rules import rules

# build up an inverse index of the rules, so we can walk backwards
inverse_index = {}
for outer, inner in rules.items():
    if inner is None:
        continue
    inner_colors = inner.keys()
    for color in inner_colors:
        if not color in inverse_index:
            inverse_index[color] = []
        inverse_index[color].append(outer)

possible_outer_colors = []

queue = inverse_index["shiny gold"]
while queue:
    color = queue.pop(0)
    if color not in possible_outer_colors:
        possible_outer_colors.append(color)
    if color in inverse_index:
        queue += inverse_index[color]

print(len(possible_outer_colors))
