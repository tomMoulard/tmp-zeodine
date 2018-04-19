var allWorkspaces;
var wss;

function preload() {
    allWorkspaces = [];

    console.log("preload");
    wss = loadJSON("http://147.135.194.248:8081/ws/1524134993")
}

function setup() {
    for (var i = 0; i < wss.ws.length; i++) {
        var newwork = new Workspaces(wss.ws[i].ws_id, wss.ws[i].ws_name)
        document.body.appendChild(newwork.elm);
        console.log(newwork.elm);
        append(allWorkspaces, newwork);
    }
}