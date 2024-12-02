Day.get_parsed("input.txt")
|> Enum.map(&Day.is_safe_with_removal?(&1))
|> Enum.reduce(0, fn result, acc ->
  case result do
    true -> acc + 1
    _ -> acc
  end
end)
|> IO.puts()
