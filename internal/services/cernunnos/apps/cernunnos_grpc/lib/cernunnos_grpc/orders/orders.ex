defmodule CernunnosGrpc.Orders do
  @moduledoc """
  Module CernunnosGrpc.Orders provides an orders interface
  """
  use Agent
  use Cernunnos.Orders.V1.NewOrderRequest

  def start_link(_) do
    Agent.start_link(
        fn -> {%{}, 1} end,
        name: __MODULE__,
      )
  end

  @doc """
  Arguments: Cernunnos.Orders.V1.NewOrderRequest
  """
  def create_order(new_order_req) do
    # TODO validate order request and write it into database
  end
end
