# toolbox

I'm new go coder, toolbox is small util project for my work.

any question you can send mail to  wei.su@stevensu.me


rand provide simple Rand generater func , like string, number, hex

crypto/aes  provide easy useful aes func

orm/mgo provide mongo ORM , you can use DAO desgin pattern


net provide simple way use golang network libraray
	
 ```go
 //HTTPClient
 //RESTClient
 	var err error
 	resp := &map[string]interface{}{}
 	client := NewHTTPSimpleREST()
	client.Headers["X-Identity-Domain"] = "cncsmtrail3578"
	err = client.Put("http://localhost:9999", data, resp)
	err = client.Post("http://localhost:9999", data, resp)
	
 ```