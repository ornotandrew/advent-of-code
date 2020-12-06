with open("input.txt", "r") as f:
    groups = f.read().split("\n\n")


def count_unique_chars(s):
    return len(set(s.replace("\n", "")))


print(sum([count_unique_chars(g) for g in groups]))
