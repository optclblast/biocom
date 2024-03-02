defmodule Cernunnos.MixProject do
  use Mix.Project

  def project do
    [
      app: :cernunnos,
      version: "0.1.0",
      elixir: "~> 1.15",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger],
      mod: {Cernunnos.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:protobuf, "~> 0.10.0"},
      {:google_protos, "~> 0.1"},
      {:grpc, "~> 0.6", hex: :grpc_fresha}
    ]
  end
end
