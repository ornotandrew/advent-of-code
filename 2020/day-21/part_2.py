from food import index

ingredients = []
for allergen in sorted(index.keys()):
    ingredients.append(list(index[allergen])[0])

print(",".join(ingredients))
