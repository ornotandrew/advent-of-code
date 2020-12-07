from rules import rules

# prime the cache with the leaf nodes
cache = {color: 1 for color, inner_colors in rules.items() if inner_colors is None}


def count_bags(color):
    if color in cache:
        return cache[color]

    total = 1  # we want to count _this_ bag, as well as all of the contained ones
    for inner_color, num in rules[color].items():
        total += num * count_bags(inner_color)
    cache[color] = total
    return total


# subtract one, because we want to exclude the shiny gold bag itself
print(count_bags("shiny gold") - 1)
