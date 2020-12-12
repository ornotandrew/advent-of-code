from waiting_area import WaitingArea

waiting_area = WaitingArea(use_visible_rules=True)
state = str(waiting_area)
while True:
    next_state = next(waiting_area)
    if next_state == state:
        print(state.count("#"))
        break
    state = next_state
