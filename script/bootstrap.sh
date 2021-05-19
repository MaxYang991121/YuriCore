cd `dirname $0` && pwd
./room_service &
./user_service &
sleep 1
./main_service &
