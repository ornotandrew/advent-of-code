from game import labels, example_labels, play_round, hash_labels

for i in range(100):
    labels = play_round(labels)

print(hash_labels(labels))
