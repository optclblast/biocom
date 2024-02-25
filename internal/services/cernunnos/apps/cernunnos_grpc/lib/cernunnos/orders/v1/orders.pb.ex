defmodule Cernunnos.Orders.V1.NewOrderRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.NewOrderResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.GetOrdersRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.GetOrdersResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.UpdateOrderRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.UpdateOrderResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.DeleteOrdersRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.DeleteOrdersResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.CernunnosOrdersAPI.Service do
  @moduledoc false

  use GRPC.Service,
    name: "cernunnos.orders.v1.CernunnosOrdersAPI",
    protoc_gen_elixir_version: "0.12.0"

  rpc :NewOrder, Cernunnos.Orders.V1.NewOrderRequest, Cernunnos.Orders.V1.NewOrderResponse

  rpc :GetOrders, Cernunnos.Orders.V1.GetOrdersRequest, Cernunnos.Orders.V1.GetOrdersResponse

  rpc :UpdateOrder,
      Cernunnos.Orders.V1.UpdateOrderRequest,
      Cernunnos.Orders.V1.UpdateOrderResponse

  rpc :DeleteOrders,
      Cernunnos.Orders.V1.DeleteOrdersRequest,
      Cernunnos.Orders.V1.DeleteOrdersResponse
end

defmodule Cernunnos.Orders.V1.CernunnosOrdersAPI.Stub do
  @moduledoc false

  use GRPC.Stub, service: Cernunnos.Orders.V1.CernunnosOrdersAPI.Service
end
