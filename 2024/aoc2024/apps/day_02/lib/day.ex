defmodule Day do
  @spec get_parsed(String.t()) :: [[integer()]]
  def get_parsed(path) do
    get_file_contents(path)
    |> parse
  end

  def get_file_contents(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw)
      {:error, cause} -> throw("Error reading file: #{cause}")
    end
  end

  def parse(file_contents) do
    String.split(file_contents, "\n")
    |> Enum.map(fn line ->
      String.split(line, " ")
      |> Enum.map(&parse_int!(&1))
    end)
  end

  def parse_int!(substr) do
    case Integer.parse(substr) do
      {int, _} -> int
      :error -> raise "Error parsing integer: #{substr}"
    end
  end

  @spec combinations_without_1([integer()]) :: [[integer()]]
  def combinations_without_1(list) do
    Enum.reduce(0..(length(list) - 1), [], fn i, acc -> [List.delete_at(list, i) | acc] end)
    |> Enum.reverse()
  end

  @spec is_safe_with_removal?([integer()]) :: boolean()
  def is_safe_with_removal?(line) do
    Enum.any?(combinations_without_1(line), &is_safe?(&1))
  end

  @spec is_safe?([integer()]) :: boolean()
  def is_safe?(line) do
    [prev | rem] = line
    [cur | rem] = rem

    case compare(cur, prev) do
      {:ok, direction} ->
        is_safe?(rem, cur, direction)

      {:error, _} ->
        false
    end
  end

  @spec is_safe?([integer()], integer(), atom()) :: boolean()
  def is_safe?(line, prev, direction) when length(line) == 1 do
    [cur] = line

    case compare(cur, prev) do
      {:ok, ^direction} -> true
      _ -> false
    end
  end

  @spec is_safe?([integer()], integer(), atom()) :: boolean()
  def is_safe?(line, prev, direction) do
    [cur | rem] = line

    case compare(cur, prev) do
      {:ok, ^direction} ->
        is_safe?(rem, cur, direction)

      _ ->
        false
    end
  end

  @spec compare(integer(), integer()) :: {:ok | :error, :desc | :asc}
  defp compare(cur, prev) do
    diff = abs(cur - prev)

    direction =
      case cur - prev do
        x when x > 0 -> :asc
        _ -> :desc
      end

    case 1 <= diff and diff <= 3 do
      true ->
        {:ok, direction}

      false ->
        {:error, direction}
    end
  end
end
