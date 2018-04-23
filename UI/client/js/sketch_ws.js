var allWorkspaces;
var createWs;
var middle;
var wss;

function preload() {
    allWorkspaces = [];
    middle = document.getElementById("middle");
    createWs = document.getElementById("NEW");

    console.log("preload");
    wss = loadJSON("http://147.135.194.248/ws/" + userid)

    var postData = { user_id: userid };
    // httpPost("http://147.135.194.248:8080/ws/", 'json', postData, function(result) {
    //     wss = result;
    // }, errorFunction);
}

function setup() {
    console.log(userid);
    createWs.addEventListener("click", createNewWorkspace)

    for (var i = 0; i < wss.ws.length; i++) {
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
    // var postData = { user_id: userid, ws_name: txt.value };
    // httpPost("http://147.135.194.248:8080/createws/", 'json', postData, function(result) {
    //     txt.value = "";
    //     location.reload();
    // }, errorFunction);

    resReq = loadJSON("http://147.135.194.248:8081/createws/" + userid + "/" + txt.value);
    txt.value = "";
    location.reload();
};

function errorFunction(err) {
    console.log(err);
    alert("Error");
};