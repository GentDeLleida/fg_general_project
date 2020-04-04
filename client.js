var PROTO_PATH = __dirname + '/game.proto'

var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var fg_proto = grpc.loadPackageDefinition(packageDefinition).fightinggame;


function main(){
    var client = new fg_proto.FightingGame('localhost:50051',
                                       grpc.credentials.createInsecure());
    var username = 'Superva69';
    var character = 'One Punch';
    
    client.StartGame({username: username, character: character}, function(err, response) {
        console.log('Aloha:', response.init_status);
      });

}

main();