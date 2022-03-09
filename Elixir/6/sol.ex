#!/bin/elixir

defmodule Point do
  defstruct X: -1,
            Y: -1,
            lines: 0
end

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

  def get_ages(input) do
    ages =
      input
      |> String.trim("\n")
      |> String.split(",")
      |> Enum.map(&String.to_integer/1)
      |> Enum.frequencies()

    Enum.map(0..8, fn n ->
      if ages[n] == nil do
        %{n => 0}
      else
        %{n => ages[n]}
      end
    end)
    |> Enum.reduce(%{}, fn x, acc ->
      Map.merge(acc, x)
    end)
  end

  def one_day(lanterns) do
    Enum.reduce(0..8, %{}, fn n, acc ->
      Map.merge(
        acc,
        case n do
          8 ->
            %{n => lanterns[0]}

          6 ->
            %{n => lanterns[n + 1] + lanterns[0]}

          _ ->
            %{n => lanterns[n + 1]}
        end
      )
    end)
  end

  def sum_lanterns(lanterns) do
    lanterns
    |> Enum.reduce(0, fn {_, v}, acc ->
      acc + v
    end)
  end

  def pass_days(lanterns, num_days) do
    Enum.reduce(1..num_days, lanterns, fn _, acc ->
      one_day(acc)
    end)
  end

  def part1 do
    get_input_file()
    |> get_ages()
    |> pass_days(80)
    |> sum_lanterns()
  end

  def part2 do
    get_input_file()
    |> get_ages()
    |> pass_days(256)
    |> sum_lanterns()
  end

  def main do
    ## print solutions

    IO.puts("Lanterns after 80 days : #{part1()}")
    IO.puts("Lanterns after 256 days : #{part2()}")
  end
end

M.main()
