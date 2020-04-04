const {EchoRequest, EchoResponse} = require('./game_pb.js');
//const {EchoServiceClient} = require('./echo_grpc_web_pb.js');

var echoService = new EchoServiceClient('http://localhost:8080');

var request = new EchoRequest();
request.setMessage('Hello World!');

echoService.echo(request, {}, function(err, response) {
  // ...
});