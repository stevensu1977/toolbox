# toolbox

I'm new go coder, toolbox is small util project for my work.

any question you can send mail to  wei.su@stevensu.me


rand provide simple Rand generater func , like string, number, hex
```go
//rand
//RandString string length
//prefix like
// Number,NumberNoZero,Hex
// Upper,Lower,,MixAll,MixUpper MixLower
//RandString(<length>)
//RandString(<length>,<prefix>)
password:=RandString(16,MixAll)

//RandInt(<length>)
phoneNumber:=RandInt(11)
```
crypto/aes  provide easy useful aes func
```go
//aes package provide simple aes encrypt/decrypt func
 payload:="Welcome1"
 ePayload,err:=AESEncrypt(payload)
 //dPayload will be "Welcome1"
 dPayload,err:=AESDecrypt(ePayload)
 //err handler if err!=nil....

```
orm/mgo provide mongo ORM , you can use DAO desgin pattern
```go
//more example see orm/mgo/context_test.go
```

net provide simple way use golang network libraray
	
 ```go
 //net package provider some network libraray
 //HTTPClient
 //RESTClient
 	var err error
 	resp := &map[string]interface{}{}
 	client := NewHTTPSimpleREST()
	client.Headers["X-Identity-Domain"] = "cncsmtrail3578"
	err = client.Put("http://localhost:9999", data, resp)
	err = client.Post("http://localhost:9999", data, resp)
	
 ```