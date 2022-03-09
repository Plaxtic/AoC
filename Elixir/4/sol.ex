#!/bin/elixir

defmodule Board do
  defstruct numbers: [[]],
            win: false
end

defmodule BingoNumber do
  defstruct value: -1,
            strike: false
end

defmodule M do
  @dimensions 5

  def get_inputs(input) do
    input
    |> String.split("\n")
    |> List.first()
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
  end

  def get_boards(input) do
    input
    |> String.split("\n")
    |> Enum.drop(1)
    |> Enum.map(&String.split/1)
    |> Enum.filter(fn x ->
      x != []
    end)
    |> Enum.map(fn x ->
      for n <- x,
          do: %BingoNumber{
            value: String.to_integer(n),
            strike: false
          }
    end)
    |> Enum.chunk_every(@dimensions)
    |> Enum.map(fn x ->
      %Board{
        numbers: x,
        win: false
      }
    end)
  end

  def get_input_file do
    File.read!("input")
    # File.read!("practice")
  end

  def winner(board) do
    if Enum.reduce(board.numbers, false, fn line, acc ->
         Enum.reduce(line, true, fn b_num, acc2 ->
           # if any are unstriked, all subsequent will be false
           b_num.strike && acc2

           # if any are true, all subsequent will be also
         end) || acc
       end) ||
         Enum.reduce(List.zip(board.numbers), false, fn {b1, b2, b3, b4, b5}, acc ->
           (b1.strike && b2.strike && b3.strike && b4.strike && b5.strike) || acc
         end) do
      true
    else
      false
    end
  end

  def winning_board([head | tail], _) do
    if winner(head) do
      winning_board([], head)
    else
      winning_board(tail, 0)
    end
  end

  def winning_board([], acc) do
    acc
  end

  def get_game(input) do
    {get_boards(input), get_inputs(input)}
  end

  def strike_board(board, num) do
    Enum.map(board.numbers, fn line ->
      Enum.map(line, fn b_num ->
        if b_num.value == num do
          %BingoNumber{
            value: b_num.value,
            strike: true
          }
        else
          b_num
        end
      end)
    end)
  end

  def strike_all(boards, num) do
    Enum.map(boards, fn board ->
      %Board{
        numbers: strike_board(board, num)
      }
    end)
  end

  #  @spec 
  def play_bingo([head | tail], boards, _) do
    new_boards = strike_all(boards, head)
    winner = winning_board(new_boards, [])

    if winner != 0 do
      play_bingo([], [], {head, winner})
    else
      play_bingo(tail, new_boards, [])
    end
  end

  def play_bingo([], _, acc) do
    acc
  end

  def play_until_looser([head | tail], boards, _) do

    ## remove winners
    new_boards = 
      strike_all(boards, head)
      |> Enum.filter(fn board -> 
        !winner(board) 
      end)

    ## play last board until it wins
    if length(new_boards) == 1 do
      play_bingo(tail, new_boards, [])
    else
      play_until_looser(tail, new_boards, {0, 0})
    end
  end

  def score_board(numbers) do
    Enum.reduce(numbers, 0, fn line, acc ->
      Enum.reduce(line, 0, fn b_num, acc2 ->
        if b_num.strike == false do
          b_num.value + acc2
        else
          acc2
        end
      end) + acc
    end)
  end

  def part1 do
    {boards, inputs} =
      get_input_file()
      |> get_game

    {final_input, winner} = play_bingo(inputs, boards, [])
    score_board(winner.numbers) * final_input
  end

  def part2 do
    {boards, inputs} =
      get_input_file()
      |> get_game

    {final_input, winner} = play_until_looser(inputs, boards, [])

    score_board(winner.numbers) * final_input

  end

  def main do
    ## print solutions
    IO.puts("First winner score : #{part1()}")
    IO.puts("Last winner score  : #{part2()}")
  end
end

M.main()
