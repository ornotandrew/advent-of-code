import sys
from bootloader import instructions, execute_step, run_until_loop_or_done


def set_instruction_at_position(pos, new_instruction):
    modified_instructions = instructions.copy()
    _, arg = modified_instructions[pos]
    modified_instructions[pos] = (new_instruction, arg)
    return modified_instructions


# the instruction we need to flip _must_ be before the loop
possible_solutions = []

pc, seen = 0, set()
while True:
    current_instruction = instructions[pc][0]
    if current_instruction == "nop":
        possible_solutions.append(set_instruction_at_position(pc, "jmp"))
    elif current_instruction == "jmp":
        possible_solutions.append(set_instruction_at_position(pc, "nop"))

    pc, _, seen = execute_step(instructions, (pc, 0, seen))
    if pc in seen:
        print(pc, seen)
        break

# brute force... :(
for soln_instructions in possible_solutions:
    pc, acc, seen = run_until_loop_or_done(soln_instructions)
    if pc >= len(soln_instructions):
        print(acc)
        break
