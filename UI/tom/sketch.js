var cards;
var mouseIsPressedHere;
var thisCardsID
var loadedCards;
var nbCards;
var nbCardsMax;
var cardPos;
var timeToSave;

function setup() {
    h = windowHeight;
    w = windowWidth;
    createCanvas(w, h);
    cards = [];
    loadedCards = [];
    mouseIsPressedHere = false;
    nbCards = floor(w / 100);
    nbCardsMax = 5;
    cardPos = 0;
    timeToSave = 0;

    for (var i = 0; i < nbCardsMax; i++) {
        var url = "http://147.135.194.248/cards/" + i;
        // var cardJson = loadJSON(url, createCard);
        // append(loadedCards, -1);
        var tmpCard = new Card(0, h - 155, "maxresdefault.jpg", "" + i);
        append(loadedCards, tmpCard);
    }
    console.log(loadedCards)
}

function draw() {
    background(51);
    // Showing all cards
    for (var i = 0; i < cards.length; i++) {
        if (cards[i]) {
            cards[i].display();
        }
    }
    // showing preloaded cards starting with cardPos
    var cardPosTmp = cardPos;
    for (var i = cardPos; i < cardPos + nbCards; i++) {
        if (cardPosTmp >= nbCardsMax) { cardPosTmp = 0; }
        var posX = map(i, 0, nbCards, 80, w-20)
        // var tmpCard = new Card(posX, h - 155, "maxresdefault.jpg", "CardID:" + cardPosTmp)
        loadedCards[cardPosTmp].x = posX;
        loadedCards[cardPosTmp].display();
        cardPosTmp += 1;
    }
    if (mouseIsPressedHere) {
        cards[thisCardsID].x = mouseX;
        cards[thisCardsID].y = mouseY;
    }
    if (timeToSave > 1000) {
        console.log("Saving ...");
        var res = [];
        for (var i = 0; i < cards.length; i++) {
            append(res, cards[i].save());
        }
        console.log({ res });
        timeToSave = -1;
    }
    timeToSave += 1;
}

function mousePressed() {
    if (mouseX < w && mouseY < h - 245) {
        cardID = isCard(mouseX, mouseY, cards);
        if (cardID == -1) {
            var newCard = new Card(mouseX, mouseY, "maxresdefault.jpg", "CardID:" + cards.length);
            append(cards, newCard);
        } else {
            mouseIsPressedHere = true;
            thisCardsID = cardID;
        }
    } else if (mouseX < w && mouseY < h){
        cardID = isCard(mouseX, mouseY, loadedCards);
        var tmpCard = new Card(w/2, h/2, loadedCards[cardID].imgURL, loadedCards[cardID].desc)
        console.log(cardID, tmpCard);
        append(cards, tmpCard);
    }
}

function mouseReleased() {
    mouseIsPressedHere = false;
}

function isCard(x, y, list) {
    for (var i = 0; i < list.length; i++) {
        if (x > list[i].x - list[i].w &&
            x < list[i].x + list[i].w &&
            y > list[i].y - list[i].h &&
            y < list[i].y + list[i].h)
            return i;
    }
    return -1;
}

function createCard(json) {
    console.log(json);
    var posX = map(json.id, 0, w, 0, nbCards);
    var newCard = new Card(posX, h - 185, json.img, json.text);
    append(cards, tmpCard);
}