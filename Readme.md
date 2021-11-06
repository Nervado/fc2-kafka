**Run bash on kafka Container**

docker exec -it <container_name> bash

**Create topic**

kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3 --replication-factor=1

**Listen topic**

kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3 --replication-factor=1
