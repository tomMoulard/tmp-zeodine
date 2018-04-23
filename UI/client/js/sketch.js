var cards;
var mouseIsPressedHere;
var doubleClick;
var thisCardsID
var loadedCards;
var nbCards;
var nbCardsMax;
var cardPos;
var timeToSave;
var bottomCards;
var ws;

function preload() {

    h = windowHeight;
    w = windowWidth;
    cards = []; //Personal cards
    loadedCards = []; //all cards provided
    bottomCards = []; //cards at the bottom of the page
    mouseIsPressedHere = false;
    doubleClick = false;
    nbCards = floor(w / 100); // max card in a row
    nbCardsMax = 5; //total cards provided
    // TODO: get this from the api
    // nbCardsMax = loadJSON("http://147.135.194.248/nbcard/1524134993/1524135042")
    cardPos = 0; //Bottom first card
    timeToSave = 0; //timer to save cards

    for (var i = 0; i < nbCardsMax; i++) {
        var url = "http://147.135.194.248/cards/" + i;
        var json = loadJSON(url, createCard)
    }
}

function setup() {
    console.log(userid, wsid)
    createCanvas(w, h);
    buildBottomCards()
}


function windowResized() {
    resizeCanvas(windowWidth, windowHeight);
    h = windowHeight;
    w = windowWidth;
    bottomCards = [];
    nbCards = floor(w / 100);
    buildBottomCards();
}

function draw() {
    background(51);
    // Showing all cards
    if (doubleClick) {
        cards[thisCardsID[0]].globalView(thisCardsID[1], [0, 0, 0]);
        return;
    }
    for (var i = 0; i < cards.length; i++) {
        cards[i].display([0, 0, 0]);
    }
    for (var i = 0; i < bottomCards.length; i++) {
        bottomCards[i].display([0, 0, 0]);
    }
    if (mouseIsPressedHere) {
        cards[thisCardsID[0]].move(mouseX, mouseY, thisCardsID[1]);
    }
    if (timeToSave > 1000) { // var req1 = new XMLHttpRequest();
        // req1.open("GET",url,true);
        // req1.addEventListener("load",createCard);
        // req1.addEventListener("error",erreur);
        // req1.send(null);

        `{"user_id": 1,"ws_id": 1,"groupes": [{"groupe_id": 12,"cards": [{"card_id": 1,"card": {"card_content": "{}"}},{"card_id": 2,"card": {"card_content": "{}"}},{"card_id": 3,"card": {"card_content": "{}"}},{"card_id": 4,"card": {"card_content": "{}"}}]},{"groupe_id": 21,"cards": [{"card_id": 5,"card": {"card_content": "{\"card_pos\":12}"}}]}]}`
        var res = [];
        for (var i = 0; i < cards.length; i++) {
            append(res, cards[i].save());
        }
        console.log("Saving ...");
        console.log({ res });
        timeToSave = -1;
    }
    cardClose()
    timeToSave += 1;
}

function doubleClicked() {
    if (doubleClick) {
        doubleClick = false;
        cards[thisCardsID[0]].rmvChild();
        return;
    }
    if (mouseX < w && mouseY < h - 245) {
        cardID = isCard(mouseX, mouseY, cards);
        if (cardID[0] != -1) { //Si il y a une carte ou un groupe selectionner par la sourie
            thisCardsID = cardID;
            doubleClick = true;
        }
    }
}

function mousePressed() {
    if (doubleClick) { return; }

    if (mouseX < w && mouseY < h - 245) {
        cardID = isCard(mouseX, mouseY, cards);
        if (cardID[0] != -1) { //Si il y a une carte ou un groupe selectionner par la sourie
            if (mouseButton === RIGHT) { // si on fait un click doit sur une carte
                cards[cardID[0]].removeCard(cardID[1]);
            } else {
                mouseIsPressedHere = true;
                thisCardsID = cardID;
            }
        }
    } else if (mouseX < w && mouseY < h) {
        cardID = isCard(mouseX, mouseY, bottomCards);
        var tmpCard = new Card(w / 2,
            h / 2,
            bottomCards[cardID[0]].imgURL,
            bottomCards[cardID[0]].desc)
        append(cards, tmpCard);
    }
}

function mouseReleased() {
    mouseIsPressedHere = false;
}

function keyPressed() {
    if (doubleClick) { return; }
    if (keyCode == LEFT_ARROW) {
        cardPos += 1
        if (cardPos >= nbCardsMax) { cardPos = 0; }
    } else if (keyCode == RIGHT_ARROW) {
        cardPos -= 1
        if (cardPos < 0) { cardPos = nbCardsMax - 1; }
    }
    buildBottomCards();
}

function createCard(json) {
    var posX = map(json.id, 0, w, 0, nbCards);
    var newCard = new Card(posX,
        h - 185,
        json.img,
        json.text);
    append(loadedCards, newCard);
}

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