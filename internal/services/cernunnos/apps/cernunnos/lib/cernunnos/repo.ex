defmodule Cernunnos.Repo do
  use Ecto.Repo,
    otp_app: :cernunnos,
    adapter: Ecto.Adapters.Postgres
end
