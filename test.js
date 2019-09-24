const ping = require("./build/Release/node-ping");

for(var i = 1; i < 30; i++) {
  var resultStr = ping.pingHost("google.com", i, 2000, 1500);

  var result = JSON.parse(resultStr);
  console.log(result.number + ") " + result.address + " -> " + result.rtt);

  if(result.complete) {
    break;
  }
}
