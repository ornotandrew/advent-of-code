from slope import Slope

slope = Slope()

count = 0
col = 0

for row in range(len(slope.trees)):
    if slope[row, col]:
        count += 1
    col += 3

print(count)
