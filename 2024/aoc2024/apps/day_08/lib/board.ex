defmodule Point do
  defstruct row: 0, col: 0
  @type t :: %Point{row: number(), col: number()}

  @spec new(number(), number()) :: t
  def new(row, col) do
    %Point{row: row, col: col}
  end

  @spec inside_bounds?(t, t) :: boolean()
  def inside_bounds?(point, dimensions) do
    {pr, pc, dr, dc} = {point.row, point.col, dimensions.row, dimensions.col}
    not (pr < 0 or pc < 0 or pr > dr or pc > dc)
  end

  def plus(p1, p2) do
    %Point{row: p1.row + p2.row, col: p1.col + p2.col}
  end

  def minus(p1, p2) do
    %Point{row: p1.row - p2.row, col: p1.col - p2.col}
  end
end

defmodule Board do
  @type t :: %Board{
          antennae: map(),
          dimensions: Point.t()
        }
  defstruct antennae: %{}, dimensions: %Point{}

  @spec with_antenna(t, String.t(), Point.t()) :: t
  def with_antenna(board, freq, point) do
    %Board{
      board
      | antennae:
          Map.put(
            board.antennae,
            freq,
            MapSet.put(Map.get(board.antennae, freq, %MapSet{}), point)
          )
    }
  end

  @spec antinodes_in_direction(t(), [Point.t()], Point.t()) :: MapSet.t(Point.t())
  def antinodes_in_direction(board, acc, direction) do
    [p | _] = acc
    new_point = Point.plus(p, direction)

    if Point.inside_bounds?(new_point, board.dimensions) do
      antinodes_in_direction(board, [new_point | acc], direction)
    else
      acc |> MapSet.new()
    end
  end

  @spec antinodes(board :: Board.t(), p1 :: Point.t(), p2 :: Point.t(), repeats :: boolean()) :: [
          Point.t()
        ]
  def antinodes(board, p1, p2, repeats \\ false) do
    diff_p1 = Point.minus(p2, p1)
    diff_p2 = Point.minus(p1, p2)

    if not repeats do
      [Point.plus(p1, diff_p2), Point.plus(p2, diff_p1)]
      |> Enum.filter(&Point.inside_bounds?(&1, board.dimensions))
    else
      MapSet.union(
        antinodes_in_direction(board, [p1], diff_p2),
        antinodes_in_direction(board, [p2], diff_p1)
      )
    end
  end

  def all_antinodes(board, repeats \\ false) do
    board.antennae
    |> Enum.map(fn {_, points} ->
      Combinatorics.n_combinations(2, MapSet.to_list(points))
      |> Enum.map(fn [p1, p2] -> antinodes(board, p1, p2, repeats) end)
      |> Enum.reduce(%MapSet{}, fn points, acc ->
        Enum.reduce(points, acc, &MapSet.put(&2, &1))
      end)
    end)
    |> Enum.reduce(%MapSet{}, &MapSet.union(&2, &1))
  end
end
