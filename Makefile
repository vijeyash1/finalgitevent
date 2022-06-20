.PHONY: agent client
agent:
	echo "8489" | sudo -S service docker start && export NATS_TOKEN="UfmrJOYwYCCsgQvxvcfJ3BdI6c8WBbnD" && export NATS_ADDRESS="nats://localhost:4222" && docker run -d -p 4222:4222 nats:latest -js && go run ./agent
	
	
client:
	export DB_ADDRESS="127.0.0.1" && export DB_PORT="9000" && export NATS_TOKEN="UfmrJOYwYCCsgQvxvcfJ3BdI6c8WBbnD" && export NATS_ADDRESS="nats://localhost:4222" && go run ./client
	
