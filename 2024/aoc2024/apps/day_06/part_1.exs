Day.get_parsed("input.txt")
|> Board.walk()
|> then(&(MapSet.size(&1.route)))
|> IO.inspect()
