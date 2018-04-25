var allWorkspaces;
var createWs;
var middle;
var wss;

function preload() {
    allWorkspaces = [];
    middle = document.getElementById("middle");
    createWs = document.getElementById("NEW");
    //wss = loadJSON("http://147.135.194.248/ws/" + userid)
}

function setup() {
    createWs.addEventListener("click", createNewWorkspace)

    var postData = { user_id: userid };
    console.log(postData);

    // httpPost("http://147.135.194.248/ws/", 'json', postData, function(result) {
    //     wss = result;
    // }, errorFunction);

    var req = new XMLHttpRequest();

    req.open("POST", "http://147.135.194.248/ws/", true);
    req.addEventListener("load", wsLoader);
    req.addEventListener("error", errorFunction);
    req.setRequestHeader("Content-Type", "application/json");
    req.send(JSON.stringify(postData));
}

function wsLoader(ev) {
    var tab = JSON.parse(this.responseText);
    wss = tab;
    console.log(wss)

    for (var i = 0; i < wss["ws"].length; i++) {
        var newwork = new Workspaces(wss.ws[i].ws_id, wss.ws[i].ws_name)
        middle.appendChild(newwork.elm);
        console.log(newwork.elm);
        append(allWorkspaces, newwork);
    }
}

function createNewWorkspace(ev) {
    console.log(ev);
    var txt = document.getElementById("newName");
    var errTxt = document.getElementById("textErrorName");
    console.log(txt.value);

    if (txt.value == "") {
        errTxt.innerHTML = "Erreur : veuillez rentrez un nom pour votre nouveau workspace";
        return;
    }
    errTxt.innerHTML = "";
    var postData = { user_id: userid, ws_name: txt.value };

    var req = new XMLHttpRequest();
    req.open("POST", "http://147.135.194.248/createws/", true);

    req.addEventListener("load", function(result) {
        txt.value = "";
        location.reload();
    });
    req.addEventListener("error", errorFunction);
    req.setRequestHeader("Content-Type", "application/json");
    req.send(JSON.stringify(postData));


    // httpPost("http://147.135.194.248/createws/", 'json', postData, function(result) {
    //     txt.value = "";
    //     location.reload();
    // }, errorFunction);

};

function errorFunction(err) {
    console.log("Erreur :",
        err);
    alert("Error " + err.toString());
};