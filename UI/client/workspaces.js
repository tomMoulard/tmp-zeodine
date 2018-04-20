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

    divVue.className = "container"
    divVue.style.width = "202px";
    divVue.style.height = "202px";

    card1.style.width = "100px";
    card1.style.height = "100px";
    card1.style.background = "rgb(239, 239, 239)";

    card2.style.width = "100px";
    card2.style.height = "100px";
    card2.style.background = "rgb(239, 239, 239)";

    card3.style.width = "100px";
    card3.style.height = "100px";
    card3.style.background = "rgb(239, 239, 239)";

    card4.style.width = "100px";
    card4.style.height = "100px";
    card4.style.background = "rgb(239, 239, 239)";

    text.innerHTML = this.name;
}