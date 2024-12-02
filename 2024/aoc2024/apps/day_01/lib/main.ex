defmodule Main do
  def get_parsed(path) do
    get_file_contents(path)
    |> parse
  end

  def get_file_contents(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw)
      {:error, cause} -> throw("Error reading file: #{cause}")
    end
  end

  def parse(file_contents) do
    String.split(file_contents, "\n")
    |> Enum.reduce({[], []}, fn line, {leftAcc, rightAcc} ->
      [left, right] =
        String.split(line, "   ")
        |> Enum.map(
          &case Integer.parse(&1) do
            {num, _} -> num
            :error -> throw("Problem parsing int from: #{&1}")
          end
        )

      {[left | leftAcc], [right | rightAcc]}
    end)
  end
end
