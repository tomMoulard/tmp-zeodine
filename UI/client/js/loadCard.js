function cardClose() {
    for (var i = 0; i < cards.length; i++) {
        for (var j = i + 1; j < cards.length; j++) {
            if (Math.abs(cards[i].x - cards[j].x) < 20 &&
                Math.abs(cards[i].y - cards[j].y) < 20) {
                fusionCard(cards[i], cards[j])
            }
        }
    }
}

function fusionCard(card1, card2) {
    var gr = new Group(0, card1);
    gr.appendCard(card2)
    cards.splice(cards.indexOf(card1), 1);
    cards.splice(cards.indexOf(card2), 1);
    append(cards, gr);
    mouseIsPressedHere = false;
}

function isCard(x, y, list) {
    var res = [false, -1];
    for (var i = 0; i < list.length; i++) {
        res = list[i].check(x, y);
        if (res[0]) {
            return [i, res[1]]
        }
    }
    return [-1, -1];
}

function buildBottomCards() {
    bottomCards = [];
    if (nbCardsMax == -1) { return; }
    for (var i = 0; i < nbCards; i++) {
        if (cardPos >= nbCardsMax) { cardPos = 0; }
        var posX = map(i, 0, nbCards, 80, w + 20);

        var tmpCard = new Card(posX,
            h - 155,
            loadedCards[cardPos].imgURL,
            loadedCards[cardPos].desc);
        append(bottomCards, tmpCard);
        cardPos += 1;
    }
}

function errorFunction(err) {
    console.log("Erreur :",
        err);
    alert("Error " + err.toString());
}

function createCard(id, ev) {
    console.log(this.responseText)
    var card = JSON.parse(this.responseText);
    nbCardsMax++;
    var posX = map(nbCardsMax, 0, w, 0, nbCards);

    var newCard = new Card(posX,
        h - 185,
        card["img"],
        card["text"]["fr"],
        id);
    append(loadedCards, newCard);
    buildBottomCards()
}

function loadCard(ev) {
    console.log("load fait")
    json = JSON.parse(this.responseText)
    for (var i = 0; i < json["card_id"].length; i++) {
        var postData = { user_id: userid, ws_id: wsid, card_id: json["card_id"][i] };
        var req = new XMLHttpRequest();
        req.open("POST", "http://147.135.194.248/card/", true);

        req.addEventListener("load", createCard.bind(null, i));
        req.addEventListener("error", errorFunction);
        req.setRequestHeader("Content-Type", "application/json");
        req.send(JSON.stringify(postData));

    }
}

function saveCard(ev) {
    res = JSON.parse(this.responseText)
    console.log("save fait")
}