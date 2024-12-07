#!/bin/bash
day_number=$1
dir="day_${day_number}"

cd aoc2024/apps
mix new "$dir"
cd $dir

touch input.txt input_1_small.txt input_2_small.txt
rm lib/*

echo 'defmodule Day do
  def get_parsed(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw) |> parse
      {:error, cause} -> raise "Error reading file: #{cause}"
    end
  end

  def parse(file_contents) do
    String.split(file_contents, "\\n")
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
end' > lib/day.ex

echo 'Day.get_parsed("input_1_small.txt")
|> IO.inspect(charlists: :as_lists)' > part_1.exs
cat part_1.exs > part_2.exs

rm test/day_*_test.exs
echo 'defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "greets the world" do
    assert Day.hello() == :world
  end
end' > test/day_test.exs

echo "Created $dir"'
