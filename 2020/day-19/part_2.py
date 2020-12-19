from validate import validate, rules, messages

rules[8] = [(42,), (42, 8)]
rules[11] = [(42, 31), (42, 11, 31)]

print([len(m) in validate(rules[0], m) for m in messages].count(True))
