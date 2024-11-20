# wait-for-infra.sh
#!/bin/bash

echo "Waiting for infrastructure services to be healthy..."

while true; do
    HEALTHY=$(docker-compose -f docker-compose-infras.yml ps --services --filter "status=running" | wc -l)
    TOTAL=$(docker-compose -f docker-compose-infras.yml ps --services | wc -l)

    if [ "$HEALTHY" -eq "$TOTAL" ]; then
        echo "All infrastructure services are healthy!"
        break
    fi

    echo "Waiting for services to become healthy..."
    sleep 5
done
