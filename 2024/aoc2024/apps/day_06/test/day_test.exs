defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "get_parsed" do
    board = Day.get_parsed("input_1_small.txt")

    assert board == %Board{
             obstructions: [
               Point.new(9, 6),
               Point.new(8, 0),
               Point.new(7, 8),
               Point.new(6, 1),
               Point.new(4, 7),
               Point.new(3, 2),
               Point.new(1, 9),
               Point.new(0, 4)
             ],
             guard: %Guard{position: Point.new(6, 4), direction: :north},
             guard_start: Point.new(6, 4),
             dimensions: Point.new(9, 9)
           }
  end
end
