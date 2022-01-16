# dynamodb-local


### dynamodb setup on local and run sample 

```
docker-compose up -d
./run.sh 
docker-compose down
```

### use aws cli

install and setup aws cli
```
http://aws.amazon.com/cli
aws configure
```

create test table 
```
aws dynamodb create-table \
    --table-name chat \
    --attribute-definitions \
        AttributeName=UserID,AttributeType=S \
        AttributeName=PostTime,AttributeType=S \
        AttributeName=RoomID,AttributeType=S \
    --key-schema AttributeName=UserID,KeyType=HASH AttributeName=PostTime,KeyType=RANGE \
    --local-secondary-indexes '[{ "IndexName": "local_index", "Projection": { "ProjectionType": "ALL" }, "KeySchema": [{ "AttributeName": "UserID", "KeyType": "HASH" }, { "AttributeName": "RoomID", "KeyType": "RANGE" }]}]' \
    --global-secondary-indexes '[{ "IndexName": "global_index", "Projection": { "ProjectionType": "ALL" }, "KeySchema": [{ "AttributeName": "RoomID", "KeyType": "HASH" }, { "AttributeName": "PostTime", "KeyType": "RANGE" }], "ProvisionedThroughput": { "ReadCapacityUnits": 10, "WriteCapacityUnits": 10 }}]' \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url http://localhost:8000

aws dynamodb list-tables --endpoint-url http://localhost:8000

aws dynamodb put-item --table-name chat \
    --item '{"UserID":{"S":"1234"},"PostTime":{"S":"2021-12-05 10:11:10"},"RoomID":{"S":"111111"},"Message":{"S":"good morning"}}' \
    --endpoint-url http://localhost:8000
aws dynamodb put-item --table-name chat \
    --item '{"UserID":{"S":"1234"},"PostTime":{"S":"2021-12-05 10:11:12"},"RoomID":{"S":"111111"},"Message":{"S":"hello"}}' \
    --endpoint-url http://localhost:8000
aws dynamodb put-item --table-name chat \
    --item '{"UserID":{"S":"1234"},"PostTime":{"S":"2021-12-05 10:11:15"},"RoomID":{"S":"111111"},"Message":{"S":"good night"}}' \
    --endpoint-url http://localhost:8000

aws dynamodb get-item --table-name chat \
    --key '{"UserID":{"S":"1234"},"PostTime":{"S":"2021-12-05 10:11:12"}}' \
    --endpoint-url http://localhost:8000

aws dynamodb query --table-name chat \
    --key-condition-expression 'UserID = :UserID and PostTime >= :PostTime' \
    --expression-attribute-values '{ ":UserID": { "S": "1234" }, ":PostTime": { "S": "2021-12-05 10:11:13" } }' \
    --endpoint-url http://localhost:8000


```