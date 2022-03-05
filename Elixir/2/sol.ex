defmodule M do
  def get_commands do
    ## read entire file
    {:ok, body} = File.read("input")

    ## create list of commands
    command_strs = String.split(body, "\n")
    commands = for c <- command_strs, do: String.split(c)

    ## drop empty newline
    Enum.drop(commands, -1)
  end

  def do_command(cmd) do
    case cmd do
      ["forward", dist] ->
        {String.to_integer(dist), 0}

      ["down", dist] ->
        {0, String.to_integer(dist)}

      ["up", dist] ->
        {0, String.to_integer(dist) * -1}
    end
  end

  ## recursively call do_command to sum distance and depth
  def do_commands_pt1([head | tail], dist, depth) do
    {d, v} = do_command(head)
    do_commands_pt1(tail, dist + d, depth + v)
  end

  def do_commands_pt1([], dist, depth) do
    {dist, depth}
  end

  ## recursively call do_command to sum distance calulate depth
  def do_commands_pt2([head | tail], dist, depth, angle) do
    {d, a} = do_command(head)

    if angle + a > 0 do
      do_commands_pt2(tail, dist + d, depth + (angle + a) * d, angle + a)
    else
      do_commands_pt2(tail, dist + d, depth, 0)
    end
  end

  def do_commands_pt2([], dist, depth, _angle) do
    {dist, depth}
  end

  def part1 do
    commands = get_commands()
    {distance, depth} = do_commands_pt1(commands, 0, 0)
    depth * distance
  end

  def part2 do
    commands = get_commands()
    {distance, depth} = do_commands_pt2(commands, 0, 0, 0)
    depth * distance
  end

  def main do
    ## print solutions
    IO.puts("Solution1 : #{part1()}")
    IO.puts("Solution2 : #{part2()}")
  end
end

M.main()
