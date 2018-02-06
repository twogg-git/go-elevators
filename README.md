
# Elevators Concurrency Problem
What happens when you use an elevator?
- You have an X user with an initial floor and a destination floor.
- An X quantity of elevators
- An X selected elevator that was on an Y floor, with Z passengers, and a max load size. 
- An X max size of floors available in your building.
- And a big concurrency problem...

Which elevator you should choose?, How many persons can use that elevator?, Which one is the next stop? Should the elevator go up or down?, How many users are waiting? etc...

Here is an example of how you can fix that problem! ᕙ(⇀‸↼‶)ᕗ

### Tools to use:
**Golang**, because of the great Goroutines.
**Mysql** for the database management, we need those ids relationships.
**Microservices** each elevator is a service, same goes for each user request.
**Docker**, easy to setup easy to run, plus each microservice is a container, then a service in k8s.
**K8s > Kubernetes** because why not, also thanks to that container management and escalation.

### Database setup
#### Docker Image
```ssh
docker run --name mysql -e MYSQL_ROOT_PASSWORD=admin -p 3300:3306 -d mysql
```

#### Tables
elevators: Manage all the info for each elevator, like max size and status.
requests: Here we will save each person request, including initial floor, and destination floor.
operations: Process the transactions between the elevators and the requests, the main table!.
configurations: Saves the delay time between, users requests and operations to be processed.

#### Configurations initial values
- floor_count 5
- elevator_delay 2 (In seconds)
- request_delay 10 (In seconds)

### Microservices
#### Requests
It will create randomly a user request, based on floor_count, and request_delay. We are going to setup this as a microservices because eventually, we will escalate this service to more that one request at the time. This microservice will include a goroutine.

Functions: 
- sets the requests.initial_floor, >= 1
- sets the requests.destination_floor, <= configurations.floor_count
- validates that requests.initial_floor and requests.destination_floor are not equal
- requests.current_floor starts at the same floor as requests.initial_floor

#### Elevators
It will manage the operations of one elevator, where it should stop, where it needs to go, accept only as many passengers as his max load size is capable, and it erases all finish operations.

Functions:

**the elevator goes to the next floor**
- creates a variable max_transactions = elevators.max_size by operations.elevator_id
- gets operations.current_floor and operations.is_going_up by operations.elevator_id
- creates a variable next_floor = operations.current_floor ++/-- depending on the direction (operations.is_going_up)  
- get a max_transactions items from requests by requests.elevator_id order by created_at asc by operations.elevator_id
- updates all requests.current_floor to next_floor with requests.elevator_id equals to operations.elevator_id

**passengers get out form the elevator**
- remove all requests items that have requests.destination_floor equal to requests.current_floor and requests.elevator_id equals to operations.elevator_id

**elevator validates if is it open for more passengers**
- creates a variable current_requests = gets the count of current requests items by operations.elevator_id
- if current_requests < max_transactions gets a current_requests items from requests by requests.elevator_id and requests.current_floor equals to next_floor and order by requests.created_at asc 
- updates requests.elevator_id = operations.elevator_id where requests id are in previous result list

**elevator closes**
- updates operations.current_floor = next_floor 
