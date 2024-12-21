Day.get_parsed("input.txt")
|> Board.all_unique_monotonic_destinations()
|> IO.inspect(charlists: :as_lists)
