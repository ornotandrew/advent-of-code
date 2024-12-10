Day.get_parsed("input.txt")
|> Board.find_cycles()
|> MapSet.size()
|> IO.inspect()
