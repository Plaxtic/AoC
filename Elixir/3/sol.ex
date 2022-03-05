defmodule M do
  def get_bin do
    File.read!("input")
    |> String.split()
    |> Enum.map(&to_bit_array/1)
  end

  def to_bit_array(str) do
    str
    |> String.split("")
    |> Enum.filter(fn x -> x != "" end)
    |> Enum.map(&String.to_integer/1)
  end

  def add_arrays(a, b) do
    Enum.map(List.zip([a, b]), fn {a, b} -> a + b end)
  end

  def sum_bins([], accumulator) do
    accumulator
  end

  def sum_bins([head | tail], accumulator) do
    sum_bins(tail, add_arrays(head, accumulator))
  end

  def bitarray_to_int(bitarray) do
    Enum.join(bitarray)
    |> String.to_integer(2)
  end

  def gamma(bin) do
    bin
    |> most_common_bin()
    |> bitarray_to_int()
  end

  def epsilon(bin) do
    bin
    |> most_common_bin()
    |> Enum.map(fn
      1 -> 0
      0 -> 1
    end)
    |> bitarray_to_int()
  end

  def most_common_bin(bin) do
    bin
    |> sum_bins(for _ <- 1..length(List.first(bin)), do: 0)
    |> Enum.map(fn x ->
      trunc(Float.round(x / length(bin)))
    end)
  end

  def get_rating([], _index, _f, acc) do
    acc
    |> List.first()
    |> bitarray_to_int()
  end

  def get_rating(bin, index, f, _) do
    ## guard case
    if length(bin) == 1 do
      get_rating([], index, f, bin)
    else
      bit =
        bin
        |> most_common_bin()
        |> Enum.at(index)

      reduced_bin = f.(bin, index, bit)
      get_rating(reduced_bin, index + 1, f, reduced_bin)
    end
  end

  def ox_generator(bin) do
    get_rating(
      bin,
      0,
      fn bin, index, bit ->
        Enum.filter(bin, fn x ->
          Enum.at(x, index) == bit
        end)
      end,
      []
    )
  end

  def o2_scrubber(bin) do
    get_rating(
      bin,
      0,
      fn bin, index, bit ->
        Enum.filter(bin, fn x ->
          Enum.at(x, index) != bit
        end)
      end,
      []
    )
  end

  def part1 do
    ## get file as list of int lists
    bin = get_bin()
    gamma(bin) * epsilon(bin)
  end

  def part2 do
    ## get file as list of int lists
    bin = get_bin()
    o2_scrubber(bin) * ox_generator(bin)
  end

  def main do
    ## print solutions
    IO.puts("Power consumption   : #{part1()}")
    IO.puts("Life support rating : #{part2()}")
  end
end

M.main()
