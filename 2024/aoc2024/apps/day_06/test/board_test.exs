defmodule BoardTest do
  use ExUnit.Case
  doctest Board

  test "next_point" do
    board = Day.get_parsed("input_1_small.txt")

    [
      # :ok
      {Guard.new(Point.new(6, 4), :north), {:ok, Point.new(1, 4)}},
      {Guard.new(Point.new(4, 2), :east), {:ok, Point.new(4, 6)}},
      {Guard.new(Point.new(6, 4), :west), {:ok, Point.new(6, 2)}},
      {Guard.new(Point.new(0, 2), :south), {:ok, Point.new(2, 2)}},
      # :oob
      {Guard.new(Point.new(6, 5), :north), {:oob, Point.new(0, 5)}},
      {Guard.new(Point.new(6, 4), :east), {:oob, Point.new(6, 9)}},
      {Guard.new(Point.new(5, 4), :west), {:oob, Point.new(5, 0)}},
      {Guard.new(Point.new(6, 4), :south), {:oob, Point.new(9, 4)}}
    ]
    |> Enum.each(fn {guard, expected} ->
      assert Board.next_point(Board.with_guard(board, guard)) == expected
    end)
  end

  test "line_between" do
    [
      {Point.new(6, 4), Point.new(1, 4),
       [
         Point.new(6, 4),
         Point.new(5, 4),
         Point.new(4, 4),
         Point.new(3, 4),
         Point.new(2, 4),
         Point.new(1, 4)
       ]},
      {Point.new(4, 1), Point.new(4, 6),
       [
         Point.new(4, 1),
         Point.new(4, 2),
         Point.new(4, 3),
         Point.new(4, 4),
         Point.new(4, 5),
         Point.new(4, 6)
       ]}
    ]
    |> Enum.each(fn {p1, p2, expected} ->
      assert Point.line_between(p1, p2) == expected
    end)
  end

  test "walk" do
    board =
      Day.get_parsed("input_1_small.txt")
      |> Board.walk()

    assert MapSet.size(board.route) == 41
  end

  test "find_cycles" do
    board = Day.get_parsed("input_1_small.txt")

    assert Board.find_cycles(board) ==
             MapSet.new([
               Point.new(6, 3),
               Point.new(7, 6),
               Point.new(7, 7),
               Point.new(8, 1),
               Point.new(8, 3),
               Point.new(9, 7)
             ])

    board = Day.get_parsed("input_2_small.txt")
    assert Board.find_cycles(board) == MapSet.new([Point.new(0, 0)])

    board = Day.get_parsed("input_3_small.txt")

    assert Board.find_cycles(board) == MapSet.new([Point.new(1, 3)])

    board = Day.get_parsed("input.txt")
    cycles = MapSet.size(Board.find_cycles(board))
    assert cycles > 1052
    assert cycles < 1534
  end
end
