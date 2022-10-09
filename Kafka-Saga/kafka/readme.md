wurstmeister kafka
https://hub.docker.com/r/wurstmeister/kafka/

If kafka containers fall shortly after start - stop all 5 containers; erase kafka_data folder
sudo rm -r Kafka-Saga/kafka/kafka_data/

Start containers sequentially from docker_compose.yaml (kafka-ui is optional), 
wait ~10 sec between starting each container 

docker volume prune
docker system prune -a	

