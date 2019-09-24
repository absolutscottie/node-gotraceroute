//include the ping module
const ping = require("./build/Release/node-ping");

//run the exported pingHost function, gather
//results. quit when output indicates that the
//trace is complete 
for(var i = 1; i < 30; i++) {
  var resultStr = ping.pingHost("google.com", i, 2000, 1500);

  var result = JSON.parse(resultStr);
  console.log(result.number + ") " + result.address + " -> " + result.rtt);

  if(result.complete) {
    break;
  }
}
