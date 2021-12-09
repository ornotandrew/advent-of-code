with open("input.txt", "r") as f:
    # the .strip() is important here, since the file contains a trailing newline
    groups = f.read().strip().split("\n\n")


def count_common_chars(s):
    lines = s.split("\n")
    return len(set(lines[0]).intersection(*lines[1:]))


print(sum([count_common_chars(g) for g in groups]))
