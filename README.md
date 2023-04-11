## Testing GOLANG local Lambda
This is based on the aws documentation on Runtime Interface Emultor ([RIE](https://docs.aws.amazon.com/lambda/latest/dg/images-test.html))


### Testing the lambda

To build the image run 
```
docker compose up
```
After container has built successfully open a separate terminal window and curl
```
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":<yourName>}'
```
Substitute `<yourName>` with your name 😃

<b>NB:</b> Postman can also be used