from game import labels, example_labels, play_round, build_index

million = 1_000_000


def play(labels):
    extended_labels = labels + list(range(max(labels) + 1, million + 1))
    current, index = labels[0], build_index(extended_labels)

    for i in range(10 * million):
        current, index = play_round(current, index)

    return index[1] * index[index[1]]


print(play(labels))
