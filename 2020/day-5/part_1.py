from seats import get_seat_id

with open("input.txt", "r") as f:
    all_boarding_passes = f.read().splitlines()

print(max([get_seat_id(p) for p in all_boarding_passes]))
