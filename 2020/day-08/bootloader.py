def parse_line(line):
    instruction, argument = line.split(" ")
    return instruction, int(argument)


with open("input.txt", "r") as f:
    instructions = [parse_line(l) for l in f.read().splitlines()]


def execute_step(instructions, state):
    program_counter, accumulator, seen_counters = state

    seen_counters.add(program_counter)

    instruction, argument = instructions[program_counter]
    if instruction == "acc":
        accumulator += argument
        program_counter += 1
    elif instruction == "jmp":
        program_counter += argument
    else:  # nop
        program_counter += 1

    return program_counter, accumulator, seen_counters


def run_until_loop_or_done(instructions):
    pc, acc, seen = 0, 0, set()
    while pc not in seen and pc < len(instructions):
        pc, acc, seen = execute_step(instructions, (pc, acc, seen))
    return pc, acc, seen
