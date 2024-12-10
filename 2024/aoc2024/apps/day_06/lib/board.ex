defmodule Point do
  defstruct row: 0, col: 0
  @type t :: %Point{row: number(), col: number()}

  @spec new(number(), number()) :: t
  def new(row, col) do
    %Point{row: row, col: col}
  end

  @spec line_between(Point.t(), Point.t()) :: [Point.t()]
  def line_between(from, to) do
    case from.row == to.row do
      true -> Enum.map(0..(to.col - from.col), &Point.new(from.row, from.col + &1))
      false -> Enum.map(0..(to.row - from.row), &Point.new(from.row + &1, from.col))
    end
  end

  @spec inside_bounds?(t, t) :: boolean()
  def inside_bounds?(point, dimensions) do
    {pr, pc, dr, dc} = {point.row, point.col, dimensions.row, dimensions.col}
    not (pr < 0 or pc < 0 or pr > dr or pc > dc)
  end
end

defmodule Guard do
  defstruct position: nil, direction: nil
  @type t :: %Guard{position: Point.t(), direction: Board.direction()}

  @spec new(Point.t(), Board.direction()) :: t
  def new(pos, dir) do
    %Guard{position: pos, direction: dir}
  end
end

defmodule Board do
  @type direction :: :north | :east | :south | :west

  defstruct obstructions: [],
            guard_start: nil,
            guard: nil,
            guard_stops: %MapSet{},
            cycle_points: %MapSet{},
            route: %MapSet{},
            dimensions: Point.new(0, 0)

  @type t :: %Board{
          obstructions: [Point.t()],
          guard: Guard.t(),
          guard_start: Point.t(),
          guard_stops: MapSet.t(Guard.t()),
          cycle_points: MapSet.t(Point.t()),
          route: [Point.t()],
          dimensions: Point.t()
        }

  @spec with_obstruction(t, Point.t()) :: t
  def with_obstruction(board, p) do
    %Board{board | obstructions: [p | board.obstructions]}
  end

  @spec with_guard(t, Guard.t()) :: t
  def with_guard(board, guard) do
    %Board{board | guard: guard, guard_start: guard.position}
  end

  @spec rotate(direction) :: direction
  def rotate(dir) do
    case dir do
      :north -> :east
      :east -> :south
      :south -> :west
      :west -> :north
    end
  end

  @spec next_point(t) :: {:oob, Point.t()} | {:ok, Point.t()}
  def next_point(board) do
    {gr, gc} = {board.guard.position.row, board.guard.position.col}

    case board.guard.direction do
      :north ->
        board.obstructions
        |> Enum.filter(&(&1.col == gc and gr > &1.row))
        |> then(fn obs ->
          case obs do
            [] -> {:oob, Point.new(0, gc)}
            list -> {:ok, Point.new(Enum.max_by(list, fn p -> p.row end).row + 1, gc)}
          end
        end)

      :east ->
        board.obstructions
        |> Enum.filter(&(&1.row == gr and gc < &1.col))
        |> then(fn obs ->
          case obs do
            [] -> {:oob, Point.new(gr, board.dimensions.col)}
            list -> {:ok, Point.new(gr, Enum.min_by(list, fn p -> p.col end).col - 1)}
          end
        end)

      :south ->
        board.obstructions
        |> Enum.filter(&(&1.col == gc and gr < &1.row))
        |> then(fn obs ->
          case obs do
            [] -> {:oob, Point.new(board.dimensions.row, gc)}
            list -> {:ok, Point.new(Enum.min_by(list, fn p -> p.row end).row - 1, gc)}
          end
        end)

      :west ->
        board.obstructions
        |> Enum.filter(&(&1.row == gr and gc > &1.col))
        |> then(fn obs ->
          case obs do
            [] -> {:oob, Point.new(gr, 0)}
            list -> {:ok, Point.new(gr, Enum.max_by(list, fn p -> p.col end).col + 1)}
          end
        end)
    end
  end

  @spec walk(t) :: t
  def walk(board) do
    next_stop = next_point(board)
    line = Point.line_between(board.guard.position, elem(next_stop, 1))

    case next_stop do
      {:ok, p} ->
        next_guard = Guard.new(p, rotate(board.guard.direction))

        next_board = %Board{
          board
          | route: MapSet.union(MapSet.new(line), board.route),
            guard: next_guard,
            guard_stops: MapSet.put(board.guard_stops, next_guard)
        }

        case Enum.member?(board.guard_stops, next_guard) do
          true -> {:cycle, next_board}
          false -> walk(next_board)
        end

      {:oob, _} ->
        %Board{
          board
          | route: MapSet.union(MapSet.new(line), board.route),
            guard: nil
        }
    end
  end

  @spec find_cycles(t) :: MapSet.t(Point.t())
  def find_cycles(board) do
    walked = walk(board)

    walked.route
    |> Enum.filter(fn p ->
      case Board.walk(%Board{board | obstructions: [p | board.obstructions]}) do
        {:cycle, _} -> true
        _ -> false
      end
    end)
    |> MapSet.new()
  end
end
