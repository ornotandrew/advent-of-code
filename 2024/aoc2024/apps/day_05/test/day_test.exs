defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "is_valid?" do
    day = Day.get_parsed("input_1_small.txt")
    assert(Day.is_valid?(day, Enum.with_index([61, 13, 29])) == {:invalid, 1, 2})
  end

  test "fix_invalid_list" do
    day = Day.get_parsed("input_1_small.txt")

    assert Day.fix_invalid_list(day, [75, 97, 47, 61, 53]) == [
             97,
             75,
             47,
             61,
             53
           ]

    assert Day.fix_invalid_list(day, [61, 13, 29]) == [61, 29, 13]

    assert Day.fix_invalid_list(day, [97, 13, 75, 29, 47]) == [
             97,
             75,
             47,
             29,
             13
           ]
  end
end
