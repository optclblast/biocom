defmodule CernunnosGRPC.MixProject do
  use Mix.Project

  def project do
    [
      app: :cernunnos_web,
      version: "0.1.0",
      build_path: "../../_build",
      config_path: "../../config/config.exs",
      deps_path: "../../deps",
      lockfile: "../../mix.lock",
      elixir: "~> 1.14",
      elixirc_paths: elixirc_paths(Mix.env()),
      start_permanent: Mix.env() == :prod,
      aliases: aliases(),
      deps: deps()
    ]
  end

  def application do
    [
      mod: {CernunnosGRPC.Application, []},
      extra_applications: [:logger, :runtime_tools]
    ]
  end

  defp elixirc_paths(:test), do: ["lib", "test/support"]
  defp elixirc_paths(_), do: ["lib"]

  defp deps do
    [
      {:grpc, github: "elixir-grpc/grpc"},
      {:cowlib, "~> 2.12.0", override: true},
      {:google_protos, "~> 0.4.0"},
      {:phoenix_live_view, ">= 0.0.0"}
    ]
  end

  defp aliases do
    [

    ]
  end
end
