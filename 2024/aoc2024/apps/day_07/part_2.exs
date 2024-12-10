Day.get_parsed("input.txt")
|> Enum.filter(&Day.is_valid?([:plus, :mult, :concat], &1))
|> Enum.map(&(elem(&1, 0)))
|> Enum.sum()
|> IO.inspect()
