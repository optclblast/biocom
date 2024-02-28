defmodule CernunnosGrpc.GRPC do
  use Agent

  def start_link(_) do
    Agent.start_link(
      fn -> {%{}, 1} end,
      name: __MODULE__,
      )
  end
end
