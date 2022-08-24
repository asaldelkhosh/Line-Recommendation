// creating an express server for user to connect
const express = require("express");
const app = express();

// server port of user
const port = 3030;

// http server
const http = require("http");
const server = http.createServer(app);

// creating the app
app.use(express.static(__dirname + "/public"));

// listening on port
server.listen(port, () => console.log(`Server is running on port ${port}`));
