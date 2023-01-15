const { app, BrowserWindow } = require('electron')
// include the Node.js 'path' module at the top of your file
const path = require('path')
import WebSocket from 'ws'

// modify your existing createWindow() function
const createWindow = () => {
    // Create the browser window.
    const mainWindow = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            preload: path.join(__dirname, 'preload.js')
        }
    })

    // and load the index.html of the app.
    mainWindow.loadFile('index.html')
}

const openSocket = () => {


    
    // var conn;
    // var msg = document.getElementById("msg");
    // var log = document.getElementById("log");

    // function appendLog(item) {
    //     var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    //     log.appendChild(item);
    //     if (doScroll) {
    //         log.scrollTop = log.scrollHeight - log.clientHeight;
    //     }
    // }

    // document.getElementById("form").onsubmit = function () {
    //     if (!conn) {
    //         return false;
    //     }
    //     if (!msg.value) {
    //         return false;
    //     }
    //     conn.send(msg.value);
    //     msg.value = "";
    //     return false;
    // };



    // if (window["WebSocket"]) {
    //     conn = new WebSocket("ws://" + document.location.host + "/ws");
    //     conn.onclose = function (evt) {
    //         var item = document.createElement("div");
    //         item.innerHTML = "<b>Connection closed.</b>";
    //         appendLog(item);
    //     };
    //     conn.onmessage = function (evt) {
    //         var messages = evt.data.split('\n');
    //         for (var i = 0; i < messages.length; i++) {
    //             var item = document.createElement("div");
    //             item.innerText = messages[i];
    //             appendLog(item);
    //         }
    //     };
    // } else {
    //     var item = document.createElement("div");
    //     item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    //     appendLog(item);
    // }
}

app.whenReady().then(() => {
    createWindow()

    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) createWindow()
    })
})

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') app.quit()
})