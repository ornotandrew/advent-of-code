day = Day.get_parsed("input.txt")

Day.check_validity(day)
|> then(&(&1[:valid]))
|> Enum.map(&Enum.at(&1, floor(length(&1)/2)))
|> Enum.sum()
|> IO.inspect()
