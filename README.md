# ams-push-worker

### How to build

#### Requirements
`- go version:` 1.15

#### Build Command

Inside the projects directory issue the following command:

```
go build
```

### About

AMS Push worker is a handy cli utility that let's you simulate AMS push functionality by pulling messages from an actual AMS project/subscription and pushing them to an endpoint in your local development environment. User should first have access to an AMS project and configure a standar subscription in pull mode. Then the user can run ams-push-worker with the correct parameters and target that project/subscription for pulling messages in pre-configured intervals(ms). Ams-push-worker will try to consume new messages from AMS and push them to the user's local endpoint.
This mode is usefull to test the AMS push capability in a local environment while working on implementing the receiving endpoint. Ams-push-worker can also be configured with an authorization header that the remote endpoint might expect for validating the push server. 

### How to run

#### Arguments

`--ams-host:` AMS host and ams port

`--ams-token:` AMS access token

`--ams-sub:` AMS subscription to be used

`--ams-project:` AMS project that the subscription belongs to

`--remote-endpoint:` Remote endpoint that we expect to receive messages to

`--pull-interval:` Interval in milliseconds between pulling the next message 

`--auth-header:` Expected Authorization header value from the remote endpoint

#### Example

```
./ams-push-worker
 --ams-host $AMS_FQDN:$AMS_PORT
 --ams-token $AMS_ACCESS_TOKEN
--ams-project ps-demo-project
 --ams-sub demo-sub
 --remote-endpoint https://localhost:5000/receive_here
 --pull-interval 1000
 --auth-header s3cr3et
```
