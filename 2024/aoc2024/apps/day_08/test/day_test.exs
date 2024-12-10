defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "parse" do
    assert Day.get_parsed("input_1_small.txt") == %Board{
             dimensions: Point.new(11, 11),
             antennae: %{
               "0" =>
                 MapSet.new([
                   Point.new(1, 8),
                   Point.new(2, 5),
                   Point.new(3, 7),
                   Point.new(4, 4)
                 ]),
               "A" =>
                 MapSet.new([
                   Point.new(5, 6),
                   Point.new(8, 8),
                   Point.new(9, 9)
                 ])
             }
           }
  end

  test "antinodes" do
    board = Day.get_parsed("input_1_small.txt")

    assert Board.antinodes(board, Point.new(2, 5), Point.new(1, 8)) == [
             Point.new(3, 2),
             Point.new(0, 11)
           ]

    assert Board.antinodes(board, Point.new(5, 6), Point.new(8, 8)) == [
             Point.new(2, 4),
             Point.new(11, 10)
           ]

    assert Board.antinodes(board, Point.new(5, 6), Point.new(9, 9)) == [
             Point.new(1, 3)
           ]
  end

  test "all_antinodes" do
    board = Day.get_parsed("input_1_small.txt")

    assert Board.all_antinodes(board) ==
             MapSet.new([
               Point.new(0, 6),
               Point.new(0, 11),
               Point.new(1, 3),
               Point.new(2, 4),
               Point.new(2, 10),
               Point.new(3, 2),
               Point.new(4, 9),
               Point.new(5, 1),
               Point.new(5, 6),
               Point.new(6, 3),
               Point.new(7, 0),
               Point.new(7, 7),
               Point.new(10, 10),
               Point.new(11, 10)
             ])

    assert Board.all_antinodes(board, true) |> MapSet.size() == 34
  end

  test "antinodes_in_direction" do
    board = Day.get_parsed("input_1_small.txt")
    {p1, p2} = {Point.new(2, 5), Point.new(3, 7)}
    diff = Point.minus(p1, p2)

    assert Board.antinodes_in_direction(board, [p2], diff) ==
             MapSet.new([
               Point.new(2, 5),
               Point.new(1, 3),
               Point.new(0, 1),
               p2
             ])
  end
end
