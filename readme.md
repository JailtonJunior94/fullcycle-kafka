# Full Cycle - Kafka

1. Criando tópico
   ```
   docker exec -it kafka bash

   kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3

   kafka-topics --list --bootstrap-server=localhost:9092
   ```

2. Detalhando tópico
   ```
   kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe
   ```

3. Consumindo e produzindo mensagens
   ```
   kafka-console-producer 

   kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste 

   kafka-console-consumer 

   kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

   kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning

   kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning --group=x
   ```

4. Consumer Groups
   ```
   ```