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
end

defimpl String.Chars, for: Point do
  def to_string(point), do: "(#{point.row},#{point.col})"
end

defmodule Board do
  defstruct heights: %{}, dimensions: %Point{}

  def with_height(board, height, point) do
    %Board{board | heights: Map.put(board.heights, point, height)}
  end

  def trailheads(board) do
    Enum.filter(board.heights, &(elem(&1, 1) == 0))
    |> Enum.map(&elem(&1, 0))
    |> MapSet.new()
  end

  def neighbours(board, point) do
    [
      Point.new(point.row - 1, point.col),
      Point.new(point.row, point.col + 1),
      Point.new(point.row + 1, point.col),
      Point.new(point.row, point.col - 1)
    ]
    |> Enum.filter(&Point.inside_bounds?(&1, board.dimensions))
  end

  def monotonic_neighbours(board, point) do
    h = Map.get(board.heights, point)

    neighbours(board, point)
    |> Enum.filter(&(Map.get(board.heights, &1) == h + 1))
    |> MapSet.new()
  end

  def monotonic_destinations(board, start) do
    if Map.get(board.heights, start) == 9 do
      [start]
    else
      monotonic_neighbours(board, start)
      |> Enum.map(&monotonic_destinations(board, &1))
      |> List.flatten()
    end
  end

  def unique_monotonic_destinations(board, start) do
    monotonic_destinations(board, start)
    |> Enum.uniq()
  end

  def all_unique_monotonic_destinations(board) do
    Board.trailheads(board)
    |> Enum.map(&(unique_monotonic_destinations(board, &1) |> length))
    |> Enum.sum()
  end

  def all_monotonic_routes(board) do
    Board.trailheads(board)
    |> Enum.map(&monotonic_destinations(board, &1))
    |> Enum.map(&length(&1))
    |> Enum.sum()
  end
end
