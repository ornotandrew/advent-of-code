from bootloader import instructions, run_until_loop_or_done


_, acc, _ = run_until_loop_or_done(instructions)
print(acc)
