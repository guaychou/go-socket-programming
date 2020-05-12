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
<p align="center">
  <img src="images/client1.png" width="668" height="427" alt="Client" title="Client"></a>
  <br>
<strong><i>Client</i> send the data</strong>
<br><br>
  <img src="images/client-server-if-not-reachable.png" width="668" height="427" alt="Client-1" title="Client-1"></a>
  <br>
  <Strong>If server not reachable for 5 times, the client will be exited. </strong>
</p>

### Server<br>
![Server](images/server1.png)<br>
<p align="center">
  <img src="images/client1.png" width="668" height="427" alt="Server" title="Server"></a>
  <br>
<strong><i>Server</i> get the data</strong>
</p>
