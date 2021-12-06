sudo docker-compose up --build -d test_db
#process_id=$!
#wait $process_id

cd ./server/handlers
go test -v
cd ../..

sudo docker-compose stop test_db
sudo docker rm -f test_myapp_db

sudo rm -rf ./test_data/