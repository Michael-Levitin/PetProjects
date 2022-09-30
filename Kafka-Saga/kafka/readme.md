wurstmeister kafka
https://hub.docker.com/r/wurstmeister/kafka/

Start containers sequentially from docker_compose.yaml (kafka-ui is optional), 
wait ~10 sec between starting each container 



sudo rm -r Kafka-Saga/kafka/kafka_data/

docker volume prune
docker system prune -a	

