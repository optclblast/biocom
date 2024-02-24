defmodule Cernunnos.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      Cernunnos.Repo,
      {DNSCluster, query: Application.get_env(:cernunnos, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: Cernunnos.PubSub},
      # Start the Finch HTTP client for sending emails
      {Finch, name: Cernunnos.Finch}
      # Start a worker by calling: Cernunnos.Worker.start_link(arg)
      # {Cernunnos.Worker, arg}
    ]

    Supervisor.start_link(children, strategy: :one_for_one, name: Cernunnos.Supervisor)
  end
end
