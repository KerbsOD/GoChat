// api/index.js
var socket = new WebSocket("ws://localhost:8080/ws");

// Funcion para conectarse con el backend
let connect = (callback, username) => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

  // Cada vez que el back mande un mensaje esta funcion lo recibe.
  socket.onmessage = msg => {
    console.log(msg);
    callback(msg)
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

// Funcion para enviar mensajes al backend
let sendMsg = msg => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
