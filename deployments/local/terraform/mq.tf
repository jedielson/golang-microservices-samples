provider "rabbitmq" {
  endpoint = "http://localhost:15672"
  username = "guest"
  password = "guest"
}

resource "rabbitmq_exchange" "order_placed" {
  name  = "order.placed"
  vhost = "/"

  settings {
    type        = "headers"
    durable     = true
    auto_delete = false
  }
}

resource "rabbitmq_queue" "order_placed_order" {
  name  = "order.placed.order"
  vhost = "/"

  settings {
    durable     = true
    auto_delete = false
  }
}

resource "rabbitmq_binding" "order_placed_order_placed_order" {
  source           = rabbitmq_exchange.order_placed.name
  vhost            = "/"
  destination      = rabbitmq_queue.order_placed_order.name
  destination_type = "queue"
}