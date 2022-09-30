export MY_IP=0.0.0.0
up:
	docker-compose up -d
down:
	docker-compose down

clear:
	docker-compose rm -svf

TOPIC=foo
create_topic:
	docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic ${TOPIC} --partitions 2 --replication-factor 2 --if-not-exists --zookeeper localhost:32181

create_topic_w:
	docker run --net=host --rm wurstmeister/cp-kafka:5.0.0 kafka-topics --create --topic ${TOPIC} --partitions 2 --replication-factor 2 --if-not-exists --zookeeper localhost:32181

kafka_ui:
	docker run -p 8080:8080 \
	-e KAFKA_CLUSTERS_0_NAME=local \
	-e KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka:9092 \
	-d provectuslabs/kafka-ui:latest
#  http://localhost:8080/

stop_all:
	docker kill $(docker ps -q)


