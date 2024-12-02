{ left, right } = Main.get_parsed("input.txt")

initial_counts = Enum.reduce(
  left,
  %{},
  fn val, acc -> Map.put(acc, val,
    case Map.get(acc, val) do
      nil -> {1, 0}
      { left, _ } -> {left + 1, 0}
    end)
end)

right
|> Enum.reduce(initial_counts, fn val, acc ->
    case Map.get(acc, val) do
      nil -> acc
      { left, right } -> Map.put(acc, val, {left, right + 1 })
    end
end)
|> Enum.map(fn {k, {left, right}} -> k*left*right end)
|> Enum.reduce(&(&1 + &2))
|> IO.puts()
