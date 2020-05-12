# go-socket-programming
This is examples of implementation socket programming with golang using net library

## Data structure <br>
| Field  |      Type      |  Description |
|----------|:-------------:|:------:|
| ID | int | Identifier of client connection |
| Nama |    string   |   Name of student |
| Nim | String |    Identifier of each Student |

## Screenshot
### Client<br>
![Client](images/client.png)<br>
From the image above: *client* send the data

![Client server check](images/client-server-if-not-reachable.png)<br>
If server not reachable for 5 times, the client will be exited.
### Server<br>
![Server](images/server.png)<br>
From the image above: *server* print the data