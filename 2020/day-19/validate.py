with open("input.txt", "r") as f:
    rules, messages = f.read().split("\n\n")


def parse_sequence(pair):
    return tuple(int(g) for g in pair.split(" "))


def parse_rule(line):
    key, rule = line.split(": ")

    parsed = None
    if "|" in rule:  # "OR" cases are lists
        parsed = [parse_sequence(p) for p in rule.split(" | ")]
    elif '"' in rule:  # single letter cases are strings
        parsed = rule.replace('"', "")
    else:  # sequences of rules are tuples
        parsed = parse_sequence(rule)

    return (int(key), parsed)


rules = {key: rule for key, rule in [parse_rule(line) for line in rules.splitlines()]}

messages = messages.splitlines()


def validate(rule, message):
    """
    If the message doesn't match the rule, return [].

    If it does, return the possible next positions after the end of the match.
    NB: this can be multiple values if BOTH sides of an OR statement match.
    """
    if len(message) == 0:
        return []
    if isinstance(rule, str):
        return [1] if message[0] == rule else []
    elif isinstance(rule, tuple):
        possible_positions = [0]
        for sub_rule_key in rule:
            # for each position we _could_ be on, try the rule
            next_positions = []
            for current_pos in possible_positions:
                positions = validate(rules[sub_rule_key], message[current_pos:])
                if positions:
                    next_positions += [current_pos + next_pos for next_pos in positions]
            if not next_positions:
                return []
            possible_positions = next_positions
        return possible_positions
    elif isinstance(rule, list):
        first_option, second_option = rule
        # concat the two lists together
        return validate(first_option, message) + validate(second_option, message)
