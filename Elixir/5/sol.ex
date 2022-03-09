#!/bin/elixir

defmodule Point do
  defstruct X: -1,
            Y: -1,
            lines: 0
end

defmodule M do
  def get_vectors(input) do
    input
    |> String.split("\n")
    |> Enum.map(fn vec ->
      String.split(vec, " -> ")
      |> Enum.filter(fn vec -> vec != "" end)
      |> Enum.map(fn points ->
        pts =
          String.split(points, ",")
          |> Enum.map(&String.to_integer/1)

        %Point{
          X: List.first(pts),
          Y: List.last(pts),
          lines: 1
        }
      end)
    end)
  end

  def draw_diagonal(vec) do
    v1 = List.first(vec)
    v2 = List.last(vec)
    low_x = v1."X"
    high_x = v2."X"
    low_y = v1."Y"
    high_y = v2."Y"

    x_range = for n <- low_x..high_x, do: n
    y_range = for n <- low_y..high_y, do: n

    for {x, y} <- List.zip([x_range, y_range]),
        do: %Point{
          X: x,
          Y: y,
          lines: 1
        }
  end

  def draw_hoz_vec(vec) do
    v1 = List.first(vec)
    v2 = List.last(vec)
    low = v1."X"
    high = v2."X"

    for x <- low..high,
        do: %Point{
          X: x,
          Y: v1."Y",
          lines: 1
        }
  end

  def draw_vert_vec(vec) do
    v1 = List.first(vec)
    v2 = List.last(vec)
    low = v1."Y"
    high = v2."Y"

    for y <- low..high,
        do: %Point{
          X: v1."X",
          Y: y,
          lines: 1
        }
  end

  def remove_point(points, point) do
    Enum.filter(points, fn p ->
      p."X" != point."X" ||
        p."Y" != point."Y"
    end)
  end

  def in_points(points, point) do
    in_points(points, point, [])
  end

  def in_points([head | tail], point, _) do
    if head."X" == point."X" && head."Y" == point."Y" do
      in_points([], [], head)
    else
      in_points(tail, point, [])
    end
  end

  def in_points([], _, acc) do
    acc
  end

  def reduce_points(points) do
    reduce_points(points, points)
  end

  def reduce_points([head | tail], points) do
    match_point = in_points(points, head)

    if match_point != [] do
      new_points = remove_point(points, head)

      new_point = %Point{
        X: match_point."X",
        Y: match_point."Y",
        lines: match_point.lines + 1
      }

      reduce_points(tail, [new_point | new_points])
    else
      reduce_points(tail, points)
    end
  end

  def reduce_points([], points) do
    points
  end

  def get_input_file do
    File.read!("input")
  end

  def map_vectors(input, part) do
    input
    |> get_vectors()
    |> Enum.map(fn vec ->
      v1 = List.first(vec)
      v2 = List.last(vec)

      cond do
        v1 == nil ->
          nil

        v1."Y" == v2."Y" ->
          draw_hoz_vec(vec)

        v1."X" == v2."X" ->
          draw_vert_vec(vec)

        (v1."X" - v2."X" == v1."Y" - v2."Y" ||
           v2."X" - v1."X" == v1."Y" - v2."Y") && part > 1 ->
          draw_diagonal(vec)

        true ->
          nil
      end
    end)
    |> Enum.filter(fn p ->
      p != nil
    end)
  end

  def calc_overlapping(vectors) do
    vectors
    |> List.flatten()
    |> Enum.frequencies()
    |> Enum.count(fn {_, freq} ->
      freq > 1
    end)
  end

  def part1 do
    get_input_file()
    |> map_vectors(1)
    |> calc_overlapping()
  end

  def part2 do
    get_input_file()
    |> map_vectors(2)
    |> calc_overlapping()
  end

  def main do
    ## print solutions

    IO.puts("Overlapping lines                 : #{part1()}")
    IO.puts("Overlapping lines (inc diagonals) : #{part2()}")
  end
end

M.main()
