defmodule Grid do
  @type key() :: {integer, integer}

  defstruct index: %{}, rows: 0, cols: 0

  @type direction() ::
          :north | :east | :south | :west | :north_east | :south_east | :south_west | :north_west

  def from_string(str) do
    lines = String.split(str, "\n")

    index =
      lines
      |> Enum.reduce({0, %{}}, fn line, {line_number, index} ->
        {
          line_number + 1,
          String.graphemes(line)
          |> Enum.reduce({0, index}, fn char, {col_number, index} ->
            {col_number + 1, Map.put(index, {line_number, col_number}, char)}
          end)
          |> elem(1)
        }
      end)
      |> elem(1)

    %Grid{index: index, rows: length(lines), cols: length(String.graphemes(Enum.at(lines, 0)))}
  end

  @spec find_indexes(%Grid{}, String.grapheme()) :: [key()]
  def find_indexes(grid, char) do
    Enum.filter(grid.index, fn {_, v} -> v == char end)
    |> Enum.map(fn {k, _} -> k end)
  end

  @spec count_xmas_occurrences(%Grid{}) :: number()
  def count_xmas_occurrences(grid) do
    find_indexes(grid, "X")
    |> Enum.map(&search_xmas_8(grid, &1))
    |> Enum.sum()
  end

  @spec get_pos(%Grid{}, key(), direction()) :: key() | nil
  def get_pos(grid, key, direction) do
    case direction do
      :north when elem(key, 0) > 0 ->
        {elem(key, 0) - 1, elem(key, 1)}

      :north_east when elem(key, 0) > 0 and elem(key, 1) < grid.cols - 1 ->
        {elem(key, 0) - 1, elem(key, 1) + 1}

      :east when elem(key, 1) < grid.cols - 1 ->
        {elem(key, 0), elem(key, 1) + 1}

      :south_east when elem(key, 0) < grid.rows - 1 and elem(key, 1) < grid.cols - 1 ->
        {elem(key, 0) + 1, elem(key, 1) + 1}

      :south when elem(key, 0) < grid.rows - 1 ->
        {elem(key, 0) + 1, elem(key, 1)}

      :south_west when elem(key, 0) < grid.rows - 1 and elem(key, 1) > 0 ->
        {elem(key, 0) + 1, elem(key, 1) - 1}

      :west when elem(key, 1) > 0 ->
        {elem(key, 0), elem(key, 1) - 1}

      :north_west when elem(key, 0) > 0 and elem(key, 1) > 0 ->
        {elem(key, 0) - 1, elem(key, 1) - 1}

      _ ->
        nil
    end
  end

  def search_xmas_8(grid, start_pos) do
    Enum.reduce(
      [:north, :east, :south, :west, :north_east, :south_east, :south_west, :north_west],
      0,
      fn direction, acc ->
        case is_string(grid, start_pos, direction, String.graphemes("XMAS")) do
          true -> acc + 1
          false -> acc
        end
      end
    )
  end

  @spec is_string(%Grid{}, key(), direction(), [String.t()]) :: boolean()
  def is_string(_, _, _, remaining_chars) when length(remaining_chars) == 0 do
    true
  end

  def is_string(_, pos, _, _) when pos == nil do
    false
  end

  def is_string(grid, pos, direction, remaining_chars) when length(remaining_chars) > 0 do
    if pos == nil do
      false
    end

    [cur | rem] = remaining_chars

    case pos != nil and Map.get(grid.index, pos) == cur do
      false -> false
      true -> is_string(grid, get_pos(grid, pos, direction), direction, rem)
    end
  end

  @spec count_xmas_occurrences(%Grid{}) :: number()
  def count_x_mas_occurrences(grid) do
    find_indexes(grid, "A")
    |> Enum.filter(&is_x_mas_4?(grid, &1))
    |> length
  end

  @spec is_x_mas_4?(%Grid{}, key()) :: boolean()
  def is_x_mas_4?(grid, start_pos) do
    x =
      [:north_east, :south_west, :north_west, :south_east]
      |> Enum.map(&Map.get(grid.index, get_pos(grid, start_pos, &1)))

    case x do
      ["M", "S", "M", "S"] -> true
      ["M", "S", "S", "M"] -> true
      ["S", "M", "M", "S"] -> true
      ["S", "M", "S", "M"] -> true
      _ -> false
    end
  end
end
