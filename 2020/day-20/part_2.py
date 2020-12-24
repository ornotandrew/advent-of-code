from tiles import build_line, corners, DOWN, find_matching, RIGHT, Tile

top_left_corner = next(c for c in corners if find_matching(c, DOWN) and find_matching(c, RIGHT))

left_column = build_line(top_left_corner, DOWN)
grid = [build_line(t, RIGHT) for t in left_column]

image = "\n".join(
    ["\n".join(["".join(l) for l in zip(*[t.without_border().splitlines() for t in row])]) for row in grid]
)

all_image_variations = [t.tile for t in Tile(0, image).all_orientations()]

monster = """
                  # 
#    ##    ##    ###
 #  #  #  #  #  #   
""".strip(
    "\n"
)

grid_size = len(image.splitlines())


def is_monster(image, row, col):
    if row > grid_size - 2 or col > grid_size - 20:
        return False
    for row_num, monster_row in enumerate(monster.splitlines()):
        for col_num, char in enumerate(monster_row):
            if char == "#" and image.splitlines()[row + row_num][col + col_num] != "#":
                return False
    return True


def count_monsters(image):
    monster_count = 0
    for row_num, row in enumerate(image.splitlines()):
        for col_num in range(len(row)):
            if is_monster(image, row_num, col_num):
                monster_count += 1
    return monster_count


monster_count = max(count_monsters(i) for i in all_image_variations)

print(image.count("#") - monster_count * monster.count("#"))
