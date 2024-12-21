defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "get_parsed" do
    assert Day.get_parsed("input_1_small.txt") == [
             {:file, {0, 1}},
             {:free, 2},
             {:file, {1, 3}},
             {:free, 4},
             {:file, {2, 5}}
           ]

    assert Day.get_parsed("input_1_small.txt")
           |> Day.print() == "0..111....22222"

    assert Day.get_parsed("input_2_small.txt")
           |> Day.print() == "00...111...2...333.44.5555.6666.777.888899"
  end

  test "last_file_index" do
    assert Day.last_file_index([
             {:file, {0, 1}},
             {:free, 2},
             {:file, {1, 3}},
             {:free, 4},
             {:file, {2, 5}}
           ]) == 4

    assert Day.last_file_index([
             {:file, {0, 1}},
             {:free, 2},
             {:file, {1, 3}},
             {:free, 4},
             {:file, {2, 5}},
             {:free, 4}
           ]) == 4

    assert Day.last_file_index([
             {:file, {0, 1}},
             {:free, 2},
             {:file, {1, 3}},
             {:free, 4},
             {:free, 4},
             {:free, 4}
           ]) == 2
  end

  test "pop_n_files_from_end" do
    disk = [{:file, {0, 1}}, {:free, 2}, {:file, {1, 3}}, {:free, 4}, {:file, {2, 5}}]

    assert Day.pop_n_files_from_end(disk, 5) == {
             [{:file, {2, 5}}],
             [{:file, {0, 1}}, {:free, 2}, {:file, {1, 3}}, {:free, 4}]
           }

    assert Day.pop_n_files_from_end(disk, 6) == {
             [{:file, {2, 5}}, {:file, {1, 1}}],
             [{:file, {0, 1}}, {:free, 2}, {:file, {1, 2}}, {:free, 4}]
           }

    assert Day.pop_n_files_from_end([{:file, {2, 5}}], 6) == {
             [{:file, {2, 5}}],
             []
           }
  end

  test "move_files_left" do
    assert Day.get_parsed("input_1_small.txt")
           |> Day.move_files_left()
           |> Day.print() == "022111222......"

    assert Day.get_parsed("input_2_small.txt")
           |> Day.move_files_left()
           |> Day.print() == "0099811188827773336446555566.............."
  end

  test "move_whole_files_left" do
    assert Day.get_parsed("input_1_small.txt")
           |> Day.move_whole_files_left()
           |> Day.print() == "0..111....22222"

    assert Day.get_parsed("input_2_small.txt")
           |> Day.move_whole_files_left()
           |> Day.print() == "00992111777.44.333....5555.6666.....8888.."
  end

  test "checksum" do
    assert Day.checksum(
             file: {0, 1},
             file: {2, 2},
             file: {1, 3},
             file: {2, 3},
             free: 6
           ) == 60

    assert Day.checksum(
             file: {0, 2},
             file: {9, 2},
             file: {2, 1},
             file: {1, 3},
             file: {7, 3},
             free: 1,
             file: {4, 2},
             free: 1,
             file: {3, 3},
             free: 1,
             free: 2,
             free: 1,
             file: {5, 4},
             free: 1,
             file: {6, 4},
             free: 1,
             free: 3,
             free: 1,
             file: {8, 4},
             free: 2
           ) == 2858

    assert Day.get_parsed("input_2_small.txt")
           |> Day.move_files_left()
           |> Day.checksum() == 1928
  end

  test "part 1" do
    assert Day.get_parsed("input.txt")
           |> Day.move_files_left()
           |> Day.checksum() == 6_607_511_583_593
  end
end
