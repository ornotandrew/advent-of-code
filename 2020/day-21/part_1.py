from food import foods, index

all_ingredients = set()
for ingredients, allergens in foods:
    all_ingredients = all_ingredients.union(set(ingredients))

ingredients_without_allergens = all_ingredients - set(list(v)[0] for v in index.values())

count = 0
for ingredients, allergens in foods:
    for i in ingredients_without_allergens:
        if i in ingredients:
            count += 1


print(count)
