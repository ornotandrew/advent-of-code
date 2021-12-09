def get_seat_id(boarding_pass):
    rows = range(128)
    for i in boarding_pass[:7]:  # the first 7 chars deal with row-finding
        current_len = int(len(rows) / 2)
        rows = rows[: current_len + 1] if i == "F" else rows[current_len:]

    cols = range(8)
    for i in boarding_pass[7:]:  # the last 3 chars deal with col-finding
        current_len = int(len(cols) / 2)
        cols = cols[: current_len + 1] if i == "L" else cols[current_len:]

    return rows[0] * 8 + cols[0]
