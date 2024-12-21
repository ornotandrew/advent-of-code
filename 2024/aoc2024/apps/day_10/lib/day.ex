defmodule Day do
  def get_parsed(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw) |> parse
      {:error, cause} -> raise "Error reading file: #{cause}"
    end
  end

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
            "." ->
              board

            char ->
              height = parse_int!(char)
              Board.with_height(board, height, p)
          end
        end)
      end
    )
  end

  def parse_int!(substr) do
    case Integer.parse(substr) do
      {int, _} -> int
      :error -> raise "Error parsing integer: #{substr}"
    end
  end
end
