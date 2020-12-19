from validate import validate, rules, messages

print([len(m) in validate(rules[0], m) for m in messages].count(True))
