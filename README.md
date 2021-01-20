# telegram-webhookinfo-exporter

This is a Prometheus exporter for Telegram Bot API getWebhookInfo.

## Installation

### Binary

Binaries are available on the GitHub [releases](https://github.com/MoonLiightz/telegram-webhookinfo-exporter/releases) page.

### Docker (recommended)

There is also a [Docker Image](https://hub.docker.com/r/moonliightz/telegram-webhookinfo-exporter). You can run it using the following minimal example by passing the telegram bot token via an environment variable:

```bash
$ docker run \
  -e 'BOT_TOKEN=<your bot token>' \
  -p 2112:2112 \
  moonliightz/telegram-webhookinfo-exporter:latest
```

Or by using the way with more options through a configuration file:

```bash
$ docker run \
  -v $PWD/config.yml:/config.yml
  -p 2112:2112
  moonliightz/telegram-webhookinfo-exporter:latest
```

Example of `config.yml`:

```yml
app:
  interval: 10

telegram:
  token: <your bot token>

prometheus:
  namespace: any_namespace
  subsystem: any_subsystem
  name: pending_update_count

# http:
#   addr: 0.0.0.0
#   port: 2112
```

## License

telegram-webhookinfo-exporter is released under the [MIT license](LICENSE).
