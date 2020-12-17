from itertools import product

with open("input.txt", "r") as f:
    lines = f.read().splitlines()


def read_initial_grid(dimensions):
    grid = set()
    for y, line in enumerate(lines):
        for x, char in enumerate(line):
            if char == "#":
                grid.add((x, y) + (0,) * (dimensions - 2))
    return grid


def vec_add(vec_a, vec_b):
    return tuple(a + b for a, b in zip(vec_a, vec_b))


def surrounding_points(vec):
    surround = set()
    for diff in product([-1, 0, 1], repeat=len(vec)):
        surround.add(vec_add(vec, diff))

    surround.remove(vec)
    return surround


def step(grid):
    neighbour_count = {}
    for alive_cell in grid:
        for vec in surrounding_points(alive_cell):
            if vec not in neighbour_count:
                neighbour_count[vec] = 0
            neighbour_count[vec] += 1

    next_grid = set()
    for vec, count in neighbour_count.items():
        if vec in grid and count in [2, 3]:
            next_grid.add(vec)
        elif vec not in grid and count == 3:
            next_grid.add(vec)

    return next_grid
