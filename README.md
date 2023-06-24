# go-casbin
A simple solution to apply Basic Role-Based HTTP Authorization in Go with Casbin using Redis and PostgreSQL.


More details are in this [blog](https://dev.to/girishg4t/basic-role-based-http-authorization-in-go-with-casbin-redis-postgresql-1eh2) post

# How to run the project

```sh
go run *.go
```

Results:
```sh
curl -X POST http://bob:@0.0.0.0:3000/channel                                                                 
{"message":"Forbidden"}
curl http://bob:@0.0.0.0:3000/channel                                                                         
"ok channel get"
curl -X POST http://alice:@0.0.0.0:3000/channel                                                               
"ok channel post"
curl -X GET http://alice:@0.0.0.0:3000/channel                                                                
"ok channel get"
curl -X POST http://bob:@0.0.0.0:3000/project                                                                 
{"message":"Forbidden"}
```
