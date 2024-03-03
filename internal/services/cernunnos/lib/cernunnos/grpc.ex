defmodule Cernunnos.GRPC.Orders.Server do
  use GRPC.Server, service: Cernunnos.Orders.V1.CernunnosOrdersAPI.Service

  def new_order(request, stream) do
    :logger.warn("request: #{request}")
  end
end
