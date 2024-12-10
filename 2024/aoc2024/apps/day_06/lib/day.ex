defmodule Day do
  def get_parsed(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw) |> parse
      {:error, cause} -> raise "Error reading file: #{cause}"
    end
  end

  @spec parse(String.t()) :: Board.t()
  def parse(file_contents) do
    lines = String.split(file_contents, "\n")
    [head | _] = lines
    dimensions = Point.new(length(lines) - 1, length(String.graphemes(head)) - 1)

    lines
    |> Enum.with_index()
    |> Enum.reduce(
      %Board{dimensions: dimensions},
      fn {line, row}, board ->
        String.graphemes(line)
        |> Enum.with_index()
        |> Enum.reduce(board, fn {char, col}, board ->
          p = Point.new(row, col)

          case char do
            "." -> board
            "#" -> Board.with_obstruction(board, p)
            "^" -> Board.with_guard(board, Guard.new(p, :north))
            ">" -> Board.with_guard(board, Guard.new(p, :east))
            "v" -> Board.with_guard(board, Guard.new(p, :south))
            "<" -> Board.with_guard(board, Guard.new(p, :west))
          end
        end)
      end
    )
  end
end
