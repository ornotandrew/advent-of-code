import re


def parse_passport(passport):
    return dict([p.split(":") for p in passport.split()])


def is_valid(passport):
    for field in ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]:
        if field not in passport:
            return False

    # years
    if not (1920 <= int(passport["byr"]) <= 2002):
        return False
    if not (2010 <= int(passport["iyr"]) <= 2020):
        return False
    if not (2020 <= int(passport["eyr"]) <= 2030):
        return False

    # hgt
    hgt = re.match("^(\d+)(cm|in)$", passport["hgt"])
    if not hgt:
        return False
    hgt_num, hgt_unit = hgt.groups()
    if hgt_unit == "cm" and not (150 <= int(hgt_num) <= 193):
        return False
    if hgt_unit == "in" and not (59 <= int(hgt_num) <= 76):
        return False

    # hcl
    if not re.match("^#[0-9a-f]{6}$", passport["hcl"]):
        return False

    # ecl
    if not passport["ecl"] in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]:
        return False

    # pid
    if not re.match("^\d{9}$", passport["pid"]):
        return False

    return True


with open("input.txt", "r") as f:
    passports = f.read().split("\n\n")

print([is_valid(parse_passport(p)) for p in passports].count(True))
