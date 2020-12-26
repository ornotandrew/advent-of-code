from decks import p1, p2, score
from combat import play_until_done

p1, p2 = play_until_done(p1, p2)
winner = p1 or p2

print(score(winner))
