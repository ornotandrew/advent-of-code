Day.get_parsed("input.txt", true)
|> Enum.reduce({ true, 0 }, fn elem, {enabled, num} ->
  case elem do
    :enable -> {true, num}
    :disable -> {false, num}
    {lhs, rhs} -> {
      enabled,
      case enabled do
        true -> num + (lhs * rhs)
        false -> num
      end
    }
  end
end)
|> then(fn {_, result} -> result end)
|> IO.puts()
