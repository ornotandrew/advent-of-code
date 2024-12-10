Day.get_parsed("input.txt")
|> Board.all_antinodes()
|> MapSet.size()
|> IO.inspect()
