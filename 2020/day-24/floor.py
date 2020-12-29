from itertools import product

vec_mapping = {
    "ne": (1, 0, -1),
    "e": (1, -1, 0),
    "se": (0, -1, 1),
    "sw": (-1, 0, 1),
    "w": (-1, 1, 0),
    "nw": (0, 1, -1),
}

vec_add = lambda a, b: tuple(map(float, map(sum, zip(a, b))))

with open("input.txt", "r") as f:
    lines = f.read().splitlines()

with open("input_small.txt", "r") as f:
    small_lines = f.read().splitlines()


def follow_directions(line):
    total = (0, 0, 0)
    while line:
        if line[0] in ["e", "w"]:
            vec = vec_mapping[line[0]]
            line = line[1:]
        else:
            vec = vec_mapping[line[:2]]
            line = line[2:]
        total = vec_add(total, vec)
    return total


def surrounding_points(vec):
    directions = [
        (1, -1, 0),
        (1, 0, -1),
        (0, 1, -1),
        (-1, 1, 0),
        (-1, 0, 1),
        (0, -1, 1),
    ]
    surround = set()
    for d in directions:
        surround.add(vec_add(vec, d))

    return surround


def step(grid):
    neighbour_count = {}
    for black_tile in grid:
        for vec in surrounding_points(black_tile):
            if vec not in neighbour_count:
                neighbour_count[vec] = 0
            neighbour_count[vec] += 1

    next_grid = set()
    for vec, count in neighbour_count.items():
        if vec in grid and count in [1, 2]:
            next_grid.add(vec)
        elif vec not in grid and count == 2:
            next_grid.add(vec)

    return next_grid
