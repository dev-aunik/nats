# NATS Order and Inventory Example

Small Go example that shows service-to-service messaging with NATS.

The order service publishes an `orders.created` event every few seconds. The inventory service subscribes to that event and reduces stock for the ordered product.

## Stack

- Go
- NATS
- Docker Compose

## Run

```bash
docker compose up --build
```

You should see the order service publishing events and the inventory service receiving them.

## Services

| Service | Purpose |
| --- | --- |
| `nats` | Message broker |
| `order-service` | Publishes new order events |
| `inventory-service` | Subscribes to order events and updates stock |

## Event

Subject:

```text
orders.created
```

Payload:

```json
{
  "order_id": 101,
  "product_id": 51,
  "quantity": 1
}
```

## Notes

- `NATS_URL` is configured in `docker-compose.yml`.
- Local logs and binaries are ignored by Git.
- This is intentionally small so the message flow is easy to follow.

## Troubleshooting

- If a service exits early, run `docker compose logs <service-name>`.
- If events are not received, confirm both services use `NATS_URL=nats://nats:4222`.
- If port `4222` is already in use, change the exposed port in `docker-compose.yml`.
