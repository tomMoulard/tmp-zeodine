function Workspaces(id, name) {
    this.id = id;
    this.name = name;
    this.elm = document.createElement("div");

    var aElm = document.createElement("a");
    var divVue = document.createElement("div");
    var text = document.createElement("div");
    var card1 = document.createElement("div");
    var card2 = document.createElement("div");
    var card3 = document.createElement("div");
    var card4 = document.createElement("div");

    aElm.setAttribute('href', "/use/" + this.id);

    this.elm.appendChild(aElm);
    divVue.appendChild(card1);
    divVue.appendChild(card2);
    divVue.appendChild(card3);
    divVue.appendChild(card4);
    aElm.appendChild(divVue);
    aElm.appendChild(text);

    this.elm.className = "ws";
    this.elm.id = this.id;
    text.className = "title";

    divVue.className = "container"

    text.innerHTML = this.name;
}