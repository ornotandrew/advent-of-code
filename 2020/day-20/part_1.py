from math import prod

from tiles import tiles, corners

print(prod([c.id for c in corners]))
