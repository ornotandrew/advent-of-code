def play_round(p1, p2):
    if p1[0] > p2[0]:
        return p1[1:] + [p1[0], p2[0]], p2[1:]
    else:
        return p1[1:], p2[1:] + [p2[0], p1[0]]


def play_until_done(p1, p2):
    while p1 and p2:
        p1, p2 = play_round(p1, p2)
    return p1, p2
