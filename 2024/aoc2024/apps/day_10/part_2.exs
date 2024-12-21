Day.get_parsed("input.txt")
|> Board.all_monotonic_routes()
|> IO.inspect(charlists: :as_lists)
