defmodule Cernunnos.Orders.V1.Order do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3

  field :id, 1, type: :string
  field :company_id, 2, type: :string, json_name: "companyId"
  field :created_by, 3, type: :string, json_name: "createdBy"
  field :title, 5, type: :string
  field :created_at, 6, type: :uint64, json_name: "createdAt"
  field :updated_at, 7, type: :uint64, json_name: "updatedAt"
  field :deleted_at, 8, type: :uint64, json_name: "deletedAt"
  field :closed_at, 9, type: :uint64, json_name: "closedAt"
  field :price, 10, type: :float
  field :prime_cost, 11, type: :float, json_name: "primeCost"
end

defmodule Cernunnos.Orders.V1.NewOrderRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3

  field :order, 1, type: Cernunnos.Orders.V1.Order
end

defmodule Cernunnos.Orders.V1.NewOrderResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.GetOrdersRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3

  field :ids, 1, repeated: true, type: :string
end

defmodule Cernunnos.Orders.V1.GetOrdersResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3

  field :orders, 1, repeated: true, type: Cernunnos.Orders.V1.Order
end

defmodule Cernunnos.Orders.V1.UpdateOrderRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3

  field :new_order_state, 1, type: Cernunnos.Orders.V1.Order, json_name: "newOrderState"
end

defmodule Cernunnos.Orders.V1.UpdateOrderResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3
end

defmodule Cernunnos.Orders.V1.DeleteOrdersRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.12.0", syntax: :proto3

  field :ids, 1, repeated: true, type: :string
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