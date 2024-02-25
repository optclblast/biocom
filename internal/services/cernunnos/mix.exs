defmodule Cernunnos.Umbrella.MixProject do
  use Mix.Project

  def project do
    [
      apps_path: "apps",
      version: "0.1.0",
      start_permanent: Mix.env() == :prod,
      deps: deps(),
      aliases: aliases()
    ]
  end

  defp deps do
    [
      {:grpc, github: "elixir-grpc/grpc"},
      {:protobuf, "~> 0.10.0"},
      {:cowlib, "~> 2.12.0", override: true},
      {:google_protos, "~> 0.4.0"},
      {:phoenix_live_view, ">= 0.0.0"}
    ]
  end

  defp aliases do
    [
      # run `mix setup` in all child apps
      setup: ["cmd mix setup"]
    ]
  end
end
