defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "parse" do
    assert Day.parse("012\n345\n678") == %Board{
             dimensions: Point.new(2, 2),
             heights:
               Map.new([
                 {Point.new(0, 0), 0},
                 {Point.new(0, 1), 1},
                 {Point.new(0, 2), 2},
                 {Point.new(1, 0), 3},
                 {Point.new(1, 1), 4},
                 {Point.new(1, 2), 5},
                 {Point.new(2, 0), 6},
                 {Point.new(2, 1), 7},
                 {Point.new(2, 2), 8}
               ])
           }
  end

  test "monotonic_neighbours" do
    assert Day.get_parsed("input_2_small.txt")
           |> Board.monotonic_neighbours(Point.new(3, 5)) ==
             MapSet.new([
               Point.new(2, 5),
               Point.new(4, 5),
               Point.new(3, 4)
             ])
  end

  test "all_unique_monotonic_destinations" do
    assert Board.all_unique_monotonic_destinations(Day.get_parsed("input_1_small.txt")) == 1
    assert Board.all_unique_monotonic_destinations(Day.get_parsed("input_3_small.txt")) == 2
    assert Board.all_unique_monotonic_destinations(Day.get_parsed("input_4_small.txt")) == 4
    assert Board.all_unique_monotonic_destinations(Day.get_parsed("input_5_small.txt")) == 3
    assert Board.all_unique_monotonic_destinations(Day.get_parsed("input_6_small.txt")) == 36
  end

  test "all_monotonic_routes" do
    assert Board.all_monotonic_routes(Day.get_parsed("input_7_small.txt")) == 3
    assert Board.all_monotonic_routes(Day.get_parsed("input_6_small.txt")) == 81
  end
end
