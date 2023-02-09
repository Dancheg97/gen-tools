package devops

import "dancheg97.ru/templates/gen-tools/templates/utils"

func GenerateRedis() {
  utils.AppendToCompose(RedisYaml)
}

const RedisYaml = `
  redis:
    image: redis:6.2-alpine
    restart: always
    command: redis-server --save 20 1 --loglevel warning
    healthcheck:
      test: [ "CMD", "redis-cli", "ping" ]
      interval: 1s
      timeout: 3s
      retries: 30

`
