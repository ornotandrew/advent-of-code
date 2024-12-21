defmodule Day do
  def get_parsed(path) do
    case File.read(path) do
      {:ok, raw} -> String.trim(raw) |> parse
      {:error, cause} -> raise "Error reading file: #{cause}"
    end
  end

  def parse(file_contents) do
    String.graphemes(file_contents)
    |> Enum.map(&parse_int!(&1))
    |> classify()
  end

  def parse_int!(substr) do
    case Integer.parse(substr) do
      {int, _} -> int
      :error -> raise "Error parsing integer: #{substr}"
    end
  end

  def classify(raw) do
    Enum.reduce(raw, {0, []}, fn char, {idx, acc} ->
      if acc == [] do
        {0, [{:file, {idx, char}}]}
      else
        case hd(acc) do
          {:free, _} -> {idx + 1, [{:file, {idx + 1, char}} | acc]}
          {:file, _} -> {idx, [{:free, char} | acc]}
        end
      end
    end)
    |> then(&elem(&1, 1))
    |> Enum.reverse()
    |> Enum.filter(&(not match?({:free, 0}, &1)))
  end

  def first_free_index(disk, min_size \\ 0) do
    Enum.find_index(disk, fn {type, len} -> type == :free and len >= min_size end)
  end

  def last_file_index(disk) do
    Enum.reverse(disk)
    |> Enum.find_index(fn {type, _} -> type == :file end)
    |> then(&(length(disk) - &1 - 1))
  end

  def print(disk) do
    Enum.reduce(disk, [], fn slot, acc ->
      case slot do
        {:free, length} -> List.duplicate(".", length) ++ acc
        {:file, {idx, length}} -> List.duplicate(idx, length) ++ acc
      end
    end)
    |> Enum.reverse()
    |> Enum.join("")
  end

  def pop_n_files_from_end(disk, n) do
    Enum.reverse(disk)
    |> Enum.reduce({0, [], []}, fn slot, {num_files, files, acc} ->
      {type, details} = slot

      if num_files == n or type == :free do
        {num_files, files, [slot | acc]}
      else
        {idx, len} = details
        remaining = n - num_files

        case len > remaining do
          true ->
            {n, [{:file, {idx, remaining}} | files], [{:file, {idx, len - remaining}} | acc]}

          false ->
            {num_files + len, [{:file, {idx, len}} | files], acc}
        end
      end
    end)
    |> then(&{Enum.reverse(elem(&1, 1)), elem(&1, 2)})
  end

  def move_files_left(disk) do
    free_idx = first_free_index(disk)
    {head, [{:free, free_blocks} | tail]} = Enum.split(disk, free_idx)

    {files, tail} = pop_n_files_from_end(tail, free_blocks)

    if files == [] do
      disk
    else
      move_files_left(head ++ files ++ tail ++ [{:free, free_blocks}])
    end
  end

  def move_whole_files_left(disk) when length(disk) == 0 do
    disk
  end

  def move_whole_files_left(disk) do
    {head, [last_file | tail]} = Enum.split(disk, last_file_index(disk))
    {:file, {_, last_file_len}} = last_file

    case first_free_index(head, last_file_len) do
      nil ->
        move_whole_files_left(head) ++ [last_file | tail]

      free_idx ->
        {head, [free_slot | rest]} = Enum.split(head, free_idx)
        {:free, free_len} = free_slot

        replacement =
          case free_len - last_file_len do
            gap when gap > 0 -> [last_file, {:free, gap}]
            gap when gap == 0 -> [last_file]
          end

        move_whole_files_left(head ++ replacement ++ rest) ++ [{:free, last_file_len} | tail]
    end
  end

  def checksum(disk) do
    Enum.reduce(disk, {0, 0}, fn slot, {pos, acc} ->
      case slot do
        {:free, len} ->
          {pos + len, acc}

        {:file, {idx, length}} ->
          {pos + length,
           Enum.reduce(pos..(pos + length - 1), acc, fn pos, acc ->
             acc + idx * pos
           end)}
      end
    end)
    |> then(&elem(&1, 1))
  end
end
