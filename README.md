# go-gorutinesfun
Some gorutines examples

script inicial DB
un elevador
variables configuracion
floor_count 5
eleator_delay 2 *seconds

go rutines

1 users_requetest ()
donde_estoy: ramdon de 1 hasta floor_count
a_donde_voy: ramdon de 1 hasta floor_count diferente de donde_estoy

2 elevators_transaction
validar el registro actual del elevador designado
valida la direccion y siguiente piso dependiendo del floor_count
actializa y remueve los registro en users_requetest dependiendo del destination_floor
valida la cantidad de peticiones en users_requetest para el current_floor del elevador
si el load_size permite toma load_size -  current de users_requetest en el orden acs 
actualiza los registros en users_requetest
actualiza el registro en elevators_transaction


tablas

users_requetest
request_id
person_id
destination_floor
request_timestamp
elevetor_id

elevators_transaction
elevator_id
load_size
direction
current_floor

configuration
conf_key
conf_value

elevator_keeper
elevator_id
max_size
status

dbs
mysql cc
mongo alex
