defmodule CernunnosTest do
  use ExUnit.Case
  doctest Cernunnos

  test "greets the world" do
    assert Cernunnos.hello() == :world
  end
end
