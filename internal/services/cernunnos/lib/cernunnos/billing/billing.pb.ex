defmodule Cernunnos.Orders.V1.NewBillRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.NewBillResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.CernunnosBillingAPI.Service do
  @moduledoc false

  use GRPC.Service,
    name: "cernunnos.orders.v1.CernunnosBillingAPI",
    protoc_gen_elixir_version: "0.12.0"

  rpc :NewBill, Cernunnos.Orders.V1.NewBillRequest, Cernunnos.Orders.V1.NewBillResponse
end

defmodule Cernunnos.Orders.V1.CernunnosBillingAPI.Stub do
  @moduledoc false

  use GRPC.Stub, service: Cernunnos.Orders.V1.CernunnosBillingAPI.Service
end