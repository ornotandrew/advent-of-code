defmodule Day do
  def get_parsed(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw) |> parse
      {:error, cause} -> raise "Error reading file: #{cause}"
    end
  end

  def parse(file_contents) do
    String.split(file_contents, "\n")
    |> Enum.map(fn line ->
      [left, right] = String.split(line, ": ")

      {
        parse_int!(left),
        String.split(right, " ")
        |> Enum.map(&parse_int!(&1))
      }
    end)
  end

  def parse_int!(substr) do
    case Integer.parse(substr) do
      {int, _} -> int
      :error -> raise "Error parsing integer: #{substr}"
    end
  end

  def is_valid?(operators, {target, numbers}) do
    is_valid?(operators, target, 0, numbers)
  end

  def is_valid?(_operators, target, current, remaining) when length(remaining) == 0 do
    target == current
  end

  @spec is_valid?(
          operators :: [:atom],
          target :: integer,
          current :: integer,
          remaining :: [integer]
        ) :: boolean
  def is_valid?(operators, target, current, remaining) do
    Enum.any?(operators, fn op ->
      [num | rem] = remaining

      next =
        case op do
          :mult -> current * num
          :plus -> current + num
          :concat -> parse_int!(Integer.to_string(current) <> Integer.to_string(num))
        end

      case next > target do
        true -> false
        false -> is_valid?(operators, target, next, rem)
      end
    end)
  end
end
