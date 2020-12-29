example_labels = [3, 8, 9, 1, 2, 5, 4, 6, 7]
labels = [2, 8, 4, 5, 7, 3, 9, 6, 1]


def build_index(labels):
    index = {labels[i]: labels[i + 1] for i in range(len(labels) - 1)}
    index[labels[-1]] = labels[0]
    return index


def unravel_index(starting_number, index):
    labels = [starting_number]
    for i in range(len(index) - 1):
        labels.append(index[labels[-1]])
    return labels


def play_round(current_label, index):
    removed = [index[current_label]]
    for i in range(2):
        removed.append(index[removed[-1]])
    index[current_label] = index[removed[-1]]

    destination = current_label - 1
    while destination in removed or destination == 0:
        destination -= 1
        if destination <= 0:
            destination = len(index)
    destination_next = index[destination]

    index[destination] = removed[0]
    index[removed[-1]] = destination_next
    return index[current_label], index


example_game = [
    [3, 8, 9, 1, 2, 5, 4, 6, 7],
    [2, 8, 9, 1, 5, 4, 6, 7, 3],
    [5, 4, 6, 7, 8, 9, 1, 3, 2],
    [8, 9, 1, 3, 4, 6, 7, 2, 5],
    [4, 6, 7, 9, 1, 3, 2, 5, 8],
    [1, 3, 6, 7, 9, 2, 5, 8, 4],
    [9, 3, 6, 7, 2, 5, 8, 4, 1],
    [2, 5, 8, 3, 6, 7, 4, 1, 9],
    [6, 7, 4, 1, 5, 8, 3, 9, 2],
    [5, 7, 4, 1, 8, 3, 9, 2, 6],
    [8, 3, 7, 4, 1, 9, 2, 6, 5],
]
for current, expected in zip(example_game[:-1], example_game[1:]):
    actual = unravel_index(*play_round(current[0], build_index(current)))
    assert actual == expected, "\n".join(map(lambda l: " ".join(map(str, l)), [current, actual, expected]))


def hash_labels(labels):
    """
    Starting after the cup labeled 1, collect the other cups' labels clockwise
    into a single string with no extra characters; each number except 1 should
    appear exactly once.
    """
    one_index = labels.index(1)
    return int("".join(map(str, labels[one_index + 1 :] + labels[:one_index])))


assert hash_labels([8, 3, 7, 4, 1, 9, 2, 6, 5]) == 92658374
