## Examples
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