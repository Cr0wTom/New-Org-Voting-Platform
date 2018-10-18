cd ../basic-network
./stop.sh
docker kill $(docker ps -q)
docker rm $(docker ps -aq)
docker rmi $(docker images dev-* -q)
cd ../New-Org-Voting-Platform
./startFabric.sh
rm -r hfc-key-store
node enrollAdmin.js
node registerUser.js
