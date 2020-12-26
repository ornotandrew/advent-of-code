class RecursiveConfiguration(Exception):
    def __init__(self, p1):
        self.p1 = p1
        super(Exception, self).__init__()


def p1_wins(p1, p2, seen_configurations, game_number, round_number=1):
    if (p1, p2) in seen_configurations:
        raise RecursiveConfiguration(p1)
    else:
        seen_configurations.append((p1, p2))

    if p1[0] >= len(p1) - 1 and p2[0] >= len(p2) - 1:
        new_p1, new_p2 = play_until_done(
            p1[1:], p2[1:], seen_configurations, game_number=game_number + 1, round_number=round_number
        )
        return (not not p1, seen_configurations)
    else:
        return p1[0] > p2[0], seen_configurations


def play_until_done(p1, p2, game_number=1, round_number=1):
    prefix = "\n" if game_number > 1 else ""
    seen_p1 = []
    seen_p2 = []

    while p1 and p2:
        if p1 in seen_p1 or p2 in seen_p2:
            raise RecursiveConfiguration(p1)
        else:
            seen_p1.append(p1.copy())
            seen_p2.append(p2.copy())

        c1, c2 = p1.pop(0), p2.pop(0)
        if c1 <= len(p1) and c2 <= len(p2):
            try:
                sub_p1, sub_p2 = play_until_done(p1[:c1], p2[:c2], game_number=game_number + 1)
                p1_wins = len(sub_p1) > 0
            except RecursiveConfiguration as e:
                p1_wins = True
        else:
            p1_wins = c1 > c2

        if p1_wins:
            p1 = p1 + [c1, c2]
        else:
            p2 = p2 + [c2, c1]

        round_number += 1

    return p1, p2
