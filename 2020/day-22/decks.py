with open("input.txt", "r") as f:
    decks = [[int(c) for c in d] for d in [d.splitlines()[1:] for d in f.read().split("\n\n")]]

p1, p2 = decks


def score(deck):
    total = 0
    for i, c in enumerate(deck[::-1]):
        total += (i + 1) * c
    return total
