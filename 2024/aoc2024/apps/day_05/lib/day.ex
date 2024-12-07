defmodule Day do
  defstruct bad_orderings: %{}, lists: []

  def get_parsed(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw) |> parse
      {:error, cause} -> raise "Error reading file: #{cause}"
    end
  end

  @spec parse(String.t()) :: %Day{}
  def parse(file_contents) do
    [raw_rules, raw_lists] = String.split(file_contents, "\n\n")

    rules =
      String.split(raw_rules, "\n")
      |> Enum.map(fn line ->
        String.split(line, "|")
        |> Enum.map(&parse_int!(&1))
        |> then(fn [lhs, rhs] -> {lhs, rhs} end)
      end)

    lists =
      String.split(raw_lists, "\n")
      |> Enum.map(fn line ->
        String.split(line, ",")
        |> Enum.map(&parse_int!(&1))
      end)

    %Day{bad_orderings: index_not_before(rules), lists: lists}
  end

  def parse_int!(substr) do
    case Integer.parse(substr) do
      {int, _} -> int
      :error -> raise "Error parsing integer: #{substr}"
    end
  end

  def index_not_before(rules) do
    Enum.reduce(rules, %{}, fn {lhs, rhs}, acc ->
      existing = Map.get(acc, rhs, [])
      Map.put(acc, rhs, [lhs | existing])
    end)
  end

  @spec check_validity(%Day{}) :: [[number]]
  def check_validity(day) do
    Enum.reduce(day.lists, %{:valid => [], :invalid => []}, fn list, acc ->
      case is_valid?(day, Enum.with_index(list)) do
        :valid -> %{acc | :valid => [list | acc[:valid]]}
        {:invalid, _, _} -> %{acc | :invalid => [list | acc[:invalid]]}
      end
    end)
  end

  def is_valid?(_, list) when length(list) == 0 do
    :valid
  end

  def is_valid?(day, list_with_idx) do
    [{head, head_idx} | tail_with_idx] = list_with_idx
    not_allowed_to_come_after = Map.get(day.bad_orderings, head, [])

    mismatch =
      Enum.find(tail_with_idx, fn {tail_item, _} ->
        Enum.member?(not_allowed_to_come_after, tail_item)
      end)

    case mismatch do
      nil -> is_valid?(day, tail_with_idx)
      {_, tail_idx} -> {:invalid, head_idx, tail_idx}
    end
  end

  def fix_invalid_list(day, list) do
    case is_valid?(day, Enum.with_index(list)) do
      :valid ->
        list

      {:invalid, p0, p1} ->
        fix_invalid_list(day, swap(list, p0, p1))
    end
  end

  defp swap(a, i1, i2) do
    {first, [e1 | middle]} = Enum.split(a, i1)
    {middle, [e2 | rest]} = Enum.split(middle, i2 - i1 - 1)
    List.flatten([first, e2, middle, e1, rest])
  end
end
