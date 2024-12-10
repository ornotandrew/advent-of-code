Day.get_parsed("input.txt")
|> Board.all_antinodes(true)
|> MapSet.size()
|> IO.inspect()
