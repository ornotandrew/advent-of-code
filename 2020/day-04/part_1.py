with open("input.txt", "r") as f:
    passports = f.read().split("\n\n")


def is_valid(passport):
    for field in ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]:
        if field not in passport:
            return False
    return True


print([is_valid(p) for p in passports].count(True))
