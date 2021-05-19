cd `dirname $0` && pwd
cd ./main_service
go build -ldflags="-w -s" -buildmode=pie -o ../output/main_service
cd ../room_service
go build -ldflags="-w -s" -buildmode=pie -o ../output/room_service
cd ../user_service
go build -ldflags="-w -s" -buildmode=pie -o ../output/user_service
cd ..
cp script/* output/
cp conf/server.conf output/server.conf
