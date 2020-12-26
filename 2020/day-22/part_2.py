from decks import p1, p2, score
from recursive_combat import play_until_done, RecursiveConfiguration

try:
    p1, p2 = play_until_done(p1, p2)
    winner = p1 or p2
except RecursiveConfiguration as e:
    winner = e.p1

print(score(winner))
