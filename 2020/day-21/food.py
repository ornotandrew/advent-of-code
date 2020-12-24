with open("input.txt", "r") as f:
    lines = f.read().splitlines()


def parse_line(line):
    lhs, rhs = line.split(" (contains ")
    return set(lhs.split(" ")), set(rhs.rstrip(")").split(", "))


foods = list(map(parse_line, lines))

index = {}
for ingredients, allergens in foods:
    for allergen in allergens:
        if not allergen in index:
            index[allergen] = set(ingredients)
        index[allergen] = index[allergen].intersection(ingredients)

known_ingredient_allergens = set(list(v)[0] for v in index.values() if len(v) == 1)
while any(v for v in index.values() if len(v) > 1):
    for allergen in index.keys():
        if len(index[allergen]) == 1:
            continue
        index[allergen] = [i for i in index[allergen] if i not in known_ingredient_allergens]
        if len(index[allergen]) == 1:
            known_ingredient_allergens.add(index[allergen][0])
