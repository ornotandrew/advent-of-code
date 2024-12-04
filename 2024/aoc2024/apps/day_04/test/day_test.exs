defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "find_indexes" do
    grid = Day.get_parsed("input_1_small.txt")
    assert length(Grid.find_indexes(grid, "X")) == 19
  end

  test "search_xmas_8" do
    grid = Day.get_parsed("input_1_small.txt")
    assert Grid.search_xmas_8(grid, {9, 3}) == 2
    assert Grid.search_xmas_8(grid, {9, 5}) == 3
  end

  test "count_xmas_occurrences" do
    grid = Day.get_parsed("input_1_small.txt")
    assert Grid.count_xmas_occurrences(grid) == 18
  end

  test "search_x_mas_4" do
    grid = Day.get_parsed("input_2_small.txt")
    assert Grid.is_x_mas_4?(grid, {7, 1}) == true
  end

  test "count_x_mas_occurrences" do
    grid = Day.get_parsed("input_2_small.txt")
    assert Grid.count_x_mas_occurrences(grid) == 9

    grid = Day.get_parsed("input.txt")
    assert Grid.count_x_mas_occurrences(grid) == 1978
  end
end
