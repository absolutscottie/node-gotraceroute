const ping = require("./build/Release/node-ping");
console.log("Round trip time to host: " + ping.pingHost("www.flirt4free.com", 5));
