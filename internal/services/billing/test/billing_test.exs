defmodule BillingTest do
  use ExUnit.Case
  doctest Billing

  test "greets the world" do
    assert Billing.hello() == :world
  end
end
