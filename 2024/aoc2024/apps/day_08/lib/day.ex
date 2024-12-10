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
          case char do
            "." -> board
            freq -> Board.with_antenna(board, freq, Point.new(row, col))
          end
        end)
      end
    )
  end
end
