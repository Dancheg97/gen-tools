package devops

import "dancheg97.ru/templates/gen-tools/templates/utils"

func GenerateNats() {
	utils.AppendToCompose(NatsYaml)
}

const NatsYaml = `  nats:
    image: nats
    ports:
      - 4222:4222
    command: "--cluster_name NATS --cluster nats://0.0.0.0:6222 --http_port 8222"
    
  nats-ui:
    image: piotrpersona/nats-streaming-ui:latest
    ports:
      - 8282:8282 
    environment:
      - STAN_URL=http://nats:4222
      - STAN_MONITOR_URL=http://nats:8222
      - STAN_CLUSTER=test-cluster

`
