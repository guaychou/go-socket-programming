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

<img align="center" width="668" height="427" src="images/client1.png">
<br><br>
<p align="center">
From the image above: *client* send the data<br><br>
</p> 

![Client server check](images/client-server-if-not-reachable.png)<br>
If server not reachable for 5 times, the client will be exited.
### Server<br>
![Server](images/server1.png)<br>
From the image above: *server* print the data