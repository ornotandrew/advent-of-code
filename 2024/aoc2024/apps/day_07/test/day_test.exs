defmodule DayTest do
  use ExUnit.Case
  doctest Day

  test "parse" do
    assert Day.get_parsed("input_1_small.txt") == [
             {190, [10, 19]},
             {3267, [81, 40, 27]},
             {83, [17, 5]},
             {156, [15, 6]},
             {7290, [6, 8, 6, 15]},
             {161_011, [16, 10, 13]},
             {192, [17, 8, 14]},
             {21037, [9, 7, 18, 13]},
             {292, [11, 6, 16, 20]}
           ]
  end

  test "is_valid?" do
    assert Day.is_valid?([:plus, :mult], {190, [10, 19]}) == true
    assert Day.is_valid?([:plus, :mult], {3267, [81, 40, 27]}) == true
    assert Day.is_valid?([:plus, :mult], {161_011, [16, 10, 13]}) == false
    assert Day.is_valid?([:plus, :mult], {156, [15, 6]}) == false

    assert Day.is_valid?([:plus, :mult, :concat], {156, [15, 6]}) == true
    assert Day.is_valid?([:plus, :mult, :concat], {7290, [6, 8, 6, 15]}) == true
  end
end
