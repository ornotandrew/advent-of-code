# This solution is horrible, but works
UP, DOWN, LEFT, RIGHT = "u", "d", "l", "r"

opposite_sides = {UP: DOWN, DOWN: UP, LEFT: RIGHT, RIGHT: LEFT}


class Tile:
    def __init__(self, id, tile):
        self.id = id
        self.tile = tile

    def __repr__(self):
        return f"{self.id}\n{self.tile}"

    def __eq__(self, other):
        return repr(self) == repr(other)

    def __hash__(self):
        return hash(repr(self))

    def side(self, direction):
        if direction == UP:
            return self.tile.splitlines()[0]
        elif direction == DOWN:
            return self.tile.splitlines()[-1]
        elif direction == LEFT:
            return "".join([row[0] for row in self.tile.splitlines()])
        elif direction == RIGHT:
            return "".join([row[-1] for row in self.tile.splitlines()])

    def flip_y(self):
        return Tile(self.id, "\n".join([row[::-1] for row in self.tile.splitlines()]))

    def flip_x(self):
        return Tile(self.id, "\n".join(self.tile.splitlines()[::-1]))

    def rotate_clockwise(self):
        rows = self.tile.splitlines()
        cols = ["".join([r[c] for r in rows][::-1]) for c in range(len(rows[0]))]
        return Tile(self.id, "\n".join(cols))

    def all_orientations(self):
        tiles = [self]
        for i in range(1, 4):
            tiles.append(tiles[-1].rotate_clockwise())
        tiles.append(self.flip_x())
        for i in range(1, 4):
            tiles.append(tiles[-1].rotate_clockwise())
        tiles.append(self.flip_y())
        for i in range(1, 4):
            tiles.append(tiles[-1].rotate_clockwise())
        return list(set(tiles))

    def without_border(self):
        return "\n".join([row[1:-1] for row in self.tile.splitlines()[1:-1]])


with open("input.txt", "r") as f:
    tiles_input = f.read().split("\n\n")


def parse_tile_input(tile_input):
    lines = tile_input.splitlines()
    id = int(lines[0].split(" ")[-1][:-1])
    tile = "\n".join(lines[1:])
    return (id, tile)


tiles = [Tile(*parse_tile_input(t)) for t in tiles_input]
tile_size = len(tiles[0].tile.splitlines())
side_length = len(tiles) ** 0.5


def find_matching(tile, side):
    tile_orientations = tile.all_orientations()
    for t in tiles:
        if t in tile_orientations:
            continue
        for o in t.all_orientations():
            if tile.side(side) == o.side(opposite_sides[side]):
                return o


def build_line(tile, direction):
    remaining_ids = [t.id for t in tiles]
    remaining_ids.remove(tile.id)
    line = [tile]
    while len(line) < side_length:
        match = find_matching(line[-1], direction)
        remaining_ids.remove(match.id)
        line.append(match)
    return line


corners = [t for t in tiles if len([d for d in opposite_sides.keys() if find_matching(t, d)]) == 2]
