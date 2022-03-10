#!/bin/elixir

defmodule M do
  def get_input_file do
    case File.read("input") do
      {:ok, contents} ->
        contents

      {:error, reason} ->
        IO.puts("missing file: #{reason}")
        exit(:shutdown)
    end
  end

  def get_positions(input) do
    input
    |> String.trim("\n")
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
    |> Enum.frequencies()
  end

  def calc_fuel(amount) do
    Enum.reduce(0..amount, 0, fn price, acc ->
      acc + price
    end)
  end

  def total_fuel_required(crabs, position, fuel) do
    crabs
    |> Enum.reduce(0, fn {pos, num_crabs}, acc ->
      cond do
        pos > position ->
          acc + fuel.(pos - position) * num_crabs

        pos < position ->
          acc + fuel.(position - pos) * num_crabs

        true ->
          acc
      end
    end)
  end

  def get_position_map() do
    get_input_file()
    |> get_positions()
  end

  def get_min_fuel(max_position, crabs, fuel) do
    Enum.map(0..max_position, fn n ->
      total_fuel_required(crabs, n, fuel)
    end)
    |> Enum.min()
  end

  def get_max_position(crabs) do
    crabs
    |> Enum.max()
    |> elem(0)
  end

  def part1 do
    crabs = get_position_map()

    crabs
    |> get_max_position()
    |> get_min_fuel(crabs, fn x -> x end)
  end

  def part2 do
    crabs = get_position_map()

    crabs
    |> get_max_position()
    |> get_min_fuel(crabs, fn x -> 
      calc_fuel(x) 
    end)
  end

  def main do
    ## print solutions

    IO.puts("Closest position               : #{part1()}")
    IO.puts("Closest position (scarce fuel) : #{part2()}")
  end
end

M.main()
