var input = document.getElementById("input");
var output = document.getElementById("output");


var socket = new WebSocket("ws://localhost:8081/ws");
socket.onopen = () => {
    output.innerHTML += "Status: Connected\n";
};
socket.onmessage = (e) => {
    output.innerHTML += "Server: " + e.data + "\n";
};

let send = () => {
    socket.send(input.value);
    input.value = "";
};
