# A service usef to keep track of all the existing websocket sessions

## Client
1. Emits regular location pings to the server containing longitude and latitude
2. 

## Server
1. Tracks client IDs currently connected to the server by storing them in the redis instance
2. Reads location pings from the client, stores in the cassandra instance and fetches as required by the location, map update, traffic update and graph processing services
3. 