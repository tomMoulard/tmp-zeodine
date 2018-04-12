var cards;
var mouseIsPressedHere;
var thisCardsID
var loadedCards;
var nbCards;
var nbCardsMax;

function setup() {
    h = windowHeight - 16;
    w = windowWidth - 8;
    createCanvas(w, h);
    cards = [];
    loadedCards = [];
    mouseIsPressedHere = false;
    nbCards = floor(w / 100);
    nbCardsMax = 5
    var cardID = 1
    for (var i = 1; i < nbCards; i++) {
        if (cardID > nbCardsMax) { cardID = 1; }

        //tmp
        var posX = map(i, 0, nbCards, 0, w);
        var tmpCard = new Card(posX, h - 155, "maxresdefault.jpg", "" + posX + "\n" + cardID);
        console.log(posX)
        append(loadedCards, tmpCard);
        append(cards, tmpCard);


        var url = "http://147.135.194.248/cards/" + cardID;
        var cardJson = loadJSON(url, createCard);
        // append(loadedCards, -1);


        cardID += 1;
    }
}

function draw() {
    background(51);
    for (var i = 0; i < cards.length; i++) {
        cards[i].display();
    }
    for (var i = 0; i < loadedCards.length; i++) {
        loadedCards[i].display();
    }
    if (mouseIsPressedHere) {
        cards[thisCardsID].x = mouseX;
        cards[thisCardsID].y = mouseY;
    }
}

function mousePressed() {
    if (mouseX < w && mouseY < h) {
        cardID = isCard(mouseX, mouseY);
        if (cardID == -1) {
            var newCard = new Card(mouseX, mouseY, "maxresdefault.jpg", "CardID:" + cards.length);
            append(cards, newCard);
        } else {
            mouseIsPressedHere = true;
            thisCardsID = cardID;
        }
    }
}

function mouseReleased() {
    mouseIsPressedHere = false;
}

function isCard(x, y) {
    for (var i = 0; i < cards.length; i++) {
        if (x > cards[i].x - cards[i].w &&
            x < cards[i].x + cards[i].w &&
            y > cards[i].y - cards[i].h &&
            y < cards[i].y + cards[i].h)
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