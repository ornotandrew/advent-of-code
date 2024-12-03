Day.get_parsed("input.txt", false)
|> Enum.reduce(0, fn pair, acc ->
  acc + (elem(pair, 0) * elem(pair, 1))
end)
|> IO.puts()
