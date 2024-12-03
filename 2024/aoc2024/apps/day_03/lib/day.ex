defmodule Day do
  def get_parsed(path, with_conditionals) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw) |> parse(with_conditionals)
      {:error, cause} -> raise "Error reading file: #{cause}"
    end
  end

  def parse(file_contents, with_conditionals) do
    regex =
      case with_conditionals do
        true -> ~r/mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)/
        false -> ~r/mul\((\d{1,3}),(\d{1,3})\)/
      end

    Regex.scan(regex, file_contents)
    |> Enum.map(fn match_group ->
      case length(match_group) do
        1 ->
          case Enum.at(match_group, 0) do
            "do()" -> :enable
            "don't()" -> :disable
            _ -> raise "Unexpected match: #{match_group}"
          end

        _ ->
          parse_mult(match_group)
      end
    end)
  end

  def parse_mult(match_group) do
    [_ | nums] = match_group
    [lhs, rhs] = nums |> Enum.map(&parse_int!(&1))
    {lhs, rhs}
  end

  def parse_int!(substr) do
    case Integer.parse(substr) do
      {int, _} -> int
      :error -> raise "Error parsing integer: #{substr}"
    end
  end
end
