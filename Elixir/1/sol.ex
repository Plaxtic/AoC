defmodule M do
  def get_depths do
    ## read entire file
    {:ok, body} = File.read("input") #"practice")

    ## create list of depths
    depths = String.split(body)

    ## convert to integer list
    Enum.map(depths, &String.to_integer/1)
  end

  def shift_one(list) do
    Enum.drop(list, 1) ++ [0]
  end

  def calc_increases(list) do

    ## create offset list for comparison
    current = shift_one(list)

    ## conjoin 
    offset_depths = List.zip([list, current])

    ## calculate total increases in depth
    length(for {p, c} <- offset_depths, p < c, do: 1)
  end

  def part1 do 

    ## pass depths to increase calculator
    depths = get_depths()
    calc_increases(depths)
  end

  def part2 do 

    ## get depths as integer list
    depths = get_depths()

    ## create offsets
    depths_offset_1 = shift_one(depths)
    depths_offset_2 = shift_one(depths_offset_1)

    ## conjoin 
    offset_depths = List.zip([depths, depths_offset_1, depths_offset_2])

    ## get sum off offsets
    window_sum = for {d, d1, d2} <- offset_depths, do: d+d1+d2

    ## pass window to increase calculator
    calc_increases(window_sum)
  end

  def main do

    ## print solutions
    IO.puts "Number of raw inceases                 : #{part1()}"
    IO.puts "Number of inceases with sliding window : #{part2()}"
  end
end

M.main()
