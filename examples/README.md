## How it works

```sh
 curl -X POST http://hari:@0.0.0.0:3000/project                                                                 
{"message":"Forbidden"}
curl -X POST http://alice:@0.0.0.0:3000/project                                                              
"ok project post"
curl -X POST http://bob:@0.0.0.0:3000/project                                                                  
{"message":"Forbidden"}
curl -X GET http://bob:@0.0.0.0:3000/project                                                                   
"ok project get"
curl -X GET http://alice:@0.0.0.0:3000/project                                                               
"ok project get"
curl -X POST http://bob:@0.0.0.0:3000/channel                                                                  
{"message":"Forbidden"}
curl -X POST http://gt:@0.0.0.0:3000/channel                                                                  
{"message":"Forbidden"}
```