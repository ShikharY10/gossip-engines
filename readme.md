# GOSSIP-ENGINE

A microservice that get the request forwarded by websocket gateway and process them. It get the request by a rabbitmq channel/queue.

It get the request in proper format and used Protobuf protocol for communication with other services.

The main function of this microservice is to pypass the chat message that are comming from users.

This services garranties that the chat message will be delivered to target user.