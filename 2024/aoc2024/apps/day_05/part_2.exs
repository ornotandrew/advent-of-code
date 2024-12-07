day = Day.get_parsed("input.txt")

Day.check_validity(day)
|> then(&(&1[:invalid]))
|> Enum.map(&Day.fix_invalid_list(day, &1))
|> Enum.map(&Enum.at(&1, floor(length(&1)/2)))
|> Enum.sum()
|> IO.inspect()
