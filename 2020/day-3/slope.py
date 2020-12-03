class Slope:
    def __init__(self, filename="input.txt"):
        with open("input.txt", "r") as f:
            lines = f.read().splitlines()

        self.trees = [[col == "#" for col in l] for l in lines]

    def __getitem__(self, key):
        row, col = key
        num_rows = len(self.trees)
        num_cols = len(self.trees[0])
        # we wrap to the left/right, but not up/down
        return self.trees[row][col % num_cols]
