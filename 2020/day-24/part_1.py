from floor import lines, small_lines, follow_directions

tile_counts = {}

for line in lines:
    tile = follow_directions(line)
    if not tile in tile_counts:
        tile_counts[tile] = 0
    tile_counts[tile] += 1

black_tiles = [k for k, v in tile_counts.items() if v % 2 == 1]

print(len(black_tiles))
