defmodule Cernunnos.Application do
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    :logger.warning("STARTING")

    children = [
      # Starts a worker by calling: Cernunnos.Worker.start_link(arg)
      {Cernunnos.Worker, []},
      {GRPC.Server.Supervisor, [{Cernunnos.GRPC.Worker, 50051}]}
    ]

    opts = [strategy: :one_for_one, name: Cernunnos.Supervisor]
    Supervisor.start_link(children, opts)
  end
end

defmodule Cernunnos.Worker do
  def child_spec(opts) do
    %{
      id: __MODULE__,
      start: {__MODULE__, :start_link, [opts]},
      type: :worker,
      restart: :permanent,
      shutdown: 500
    }
  end

  def start_link(_arg) do
    loop_hello()
  end

  def loop_hello() do
    IO.puts("hello, i am #{__MODULE__}") # DEBUG

    :timer.sleep(5000)

    loop_hello()
  end
end

defmodule Cernunnos.GRPC.Worker do
  use GRPC.Endpoint

  intercept GRPC.Logger.Server
  run Cernunnos.Orders.V1.CernunnosOrdersAPI.Service
end
