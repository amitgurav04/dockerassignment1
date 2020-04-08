# dockerassignment1
    This is server and client application written in golang.
### Steps to Run:
    1. Clone dockerassignment1 repository.
        git clone https://github.com/amitgurav04/dockerassignment1.git
    2. cd dockerassignment1
#### Run server application on docker container.
    3. Go inside server directory
        $cd server
    4. Build docker image of server
        $docker build -t amitgurav04/assignment_server .
    5. Run server container.
        $docker run --name assignment_server -d -p 8080:8080 amitgurav04/assignment_server
    6. See running containers.
        $ docker ps
        CONTAINER ID        IMAGE                           COMMAND             CREATED             STATUS              PORTS                    NAMES
        73e13ae58c8d        amitgurav04/assignment_server   "/server"           2 minutes ago       Up 2 minutes        0.0.0.0:8080->8080/tcp   assignment_server
    7. Create network myNetwork.
        $docker network create myNetwork
    8. Connect assignment_server container to myNetwork
        $docker network connect myNetwork assignment_server
    9. Inspect myNetwork
       $ docker network inspect myNetwork
         [
             {
                 "Name": "myNetwork",
                 "Id": "1358a28e345266cc2eec29d93579d1bbec238090d89e68f1332e5a2b76d59e77",
                 "Created": "2020-04-08T12:49:07.384400544+05:30",
                 "Scope": "local",
                 "Driver": "bridge",
                 "EnableIPv6": false,
                 "IPAM": {
                     "Driver": "default",
                     "Options": {},
                     "Config": [
                         {
                             "Subnet": "172.19.0.0/16",
                             "Gateway": "172.19.0.1"
                         }
                     ]
                 },
                 "Internal": false,
                 "Attachable": false,
                 "Ingress": false,
                 "ConfigFrom": {
                     "Network": ""
                 },
                 "ConfigOnly": false,
                 "Containers": {
                     "73e13ae58c8d9a5df30de6827b16b89afc22ae1b947ecb16d2c0bb597bce2737": {
                         "Name": "assignment_server",
                         "EndpointID": "07280f30f1b3e2c072ecd597e9166b01079454f3e5ec5309f175af225c776039",
                         "MacAddress": "02:42:ac:13:00:02",
                         "IPv4Address": "172.19.0.2/16",
                         "IPv6Address": ""
                     }
                 },
                 "Options": {},
                 "Labels": {}
             }
         ]

    10. Go inside client directory.
        $ cd ../client
    11. Build docker image of client
        $ docker build -t amitgurav04/assignment_client .
    12. Run assignment_client docker container
        $ docker run --name assignment_client -d -p 8090:8090 --env SERVER_CONTAINER_NAME=assignment_server amitgurav04/assignment_client
    13. Connect assignment_client container to myNetwork
        $docker network connect myNetwork assignment_client
    14. Inspect
        $docker network inspect myNetwork
         [
             {
                 "Name": "myNetwork",
                 "Id": "1358a28e345266cc2eec29d93579d1bbec238090d89e68f1332e5a2b76d59e77",
                 "Created": "2020-04-08T12:49:07.384400544+05:30",
                 "Scope": "local",
                 "Driver": "bridge",
                 "EnableIPv6": false,
                 "IPAM": {
                     "Driver": "default",
                     "Options": {},
                     "Config": [
                         {
                             "Subnet": "172.19.0.0/16",
                             "Gateway": "172.19.0.1"
                         }
                     ]
                 },
                 "Internal": false,
                 "Attachable": false,
                 "Ingress": false,
                 "ConfigFrom": {
                     "Network": ""
                 },
                 "ConfigOnly": false,
                 "Containers": {
                     "2c353f8e34d08f2764053e20a8402e4e8fec6ab60a0e993670181f20a8ae9db3": {
                         "Name": "assignment_client",
                         "EndpointID": "86d479b67026dcf29a68e324ab6c1594c314326c2f8fab5cd5779bd1d64d4532",
                         "MacAddress": "02:42:ac:13:00:03",
                         "IPv4Address": "172.19.0.3/16",
                         "IPv6Address": ""
                     },
                     "73e13ae58c8d9a5df30de6827b16b89afc22ae1b947ecb16d2c0bb597bce2737": {
                         "Name": "assignment_server",
                         "EndpointID": "07280f30f1b3e2c072ecd597e9166b01079454f3e5ec5309f175af225c776039",
                         "MacAddress": "02:42:ac:13:00:02",
                         "IPv4Address": "172.19.0.2/16",
                         "IPv6Address": ""
                     }
                 },
                 "Options": {},
                 "Labels": {}
             }
         ]
    15. Done.
    16. Now let's check.
        Open Browser and open http://localhost:8090/checksum this url.
        It will show following response:
            File data got from server is: Hello World!!!
            Checksum of above file data is: 0aa941b04274ae04dc5a9bd214f7d5214f36e6de
            Actual Checksum got from server is: 0aa941b04274ae04dc5a9bd214f7d5214f36e6de

