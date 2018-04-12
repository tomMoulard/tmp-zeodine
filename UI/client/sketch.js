var cards;
var mouseIsPressedHere;
var thisCardsID
var loadedCards;
var nbCards;
var nbCardsMax;
var cardPos;
var timeToSave;
var bottomCards;

function setup() {
    h                  = windowHeight;
    w                  = windowWidth;
    createCanvas(w, h);
    cards              = []; //Personal cards
    loadedCards        = []; //all cards provided
    bottomCards        = []; //cards at the bottom of the page
    mouseIsPressedHere = false;
    nbCards            = floor(w / 100); // max card in a row
    nbCardsMax         = 5; //total cards provided
    // TODO: get this from the api
    cardPos            = 0; //Bottom first card
    timeToSave         = 0; //timer to save cards

    for (var i = 0; i < nbCardsMax; i++) {
        var url = "http://147.135.194.248/cards/" + i;
        // var cardJson = loadJSON(url, createCard);
        // append(loadedCards, -1);
        var tmpCard = new Card(-100, -100, "maxresdefault.jpg", "" + i);
        append(loadedCards, tmpCard);
    }
    buildBottomCards()
    console.log(loadedCards, bottomCards)
}

function draw() {
    background(51);
    // Showing all cards
    for (var i = 0; i < cards.length; i++) {
        cards[i].display();
    }
    for (var i = 0; i < bottomCards.length; i++) {
        bottomCards[i].display();
    }
    // showing preloaded cards starting with cardPos
    // first build a "smaller" list
    if (mouseIsPressedHere) {
        cards[thisCardsID].x = mouseX;
        cards[thisCardsID].y = mouseY;
    }
    if (timeToSave > 10000) {
        var res = [];
        for (var i = 0; i < cards.length; i++) {
            append(res, cards[i].save());
        }
        // console.log("Saving ...");
        // console.log({ res });
        timeToSave = -1;
    }
    timeToSave += 1;
}

function mousePressed() {
    if (mouseX < w && mouseY < h - 245) {
        cardID = isCard(mouseX, mouseY, cards);
        if (cardID == -1) {
            // var newCard = new Card(mouseX,
            //                        mouseY,
            //                        "maxresdefault.jpg",
            //                        "CardID:" + cards.length);
            // append(cards, newCard);
        } else {
            mouseIsPressedHere = true;
            thisCardsID        = cardID;
        }
    } else if (mouseX < w && mouseY < h) {
        cardID = isCard(mouseX, mouseY, bottomCards);
        console.log(cardID, bottomCards[cardID])
        var tmpCard = new Card(w / 2,
                               h / 2,
                               bottomCards[cardID].imgURL,
                               bottomCards[cardID].desc)
        console.log(cardID, tmpCard);
        append(cards, tmpCard);
    }
}

function mouseReleased() {
    mouseIsPressedHere = false;
}

function keyPressed() {
    if (keyCode == LEFT_ARROW) {
        cardPos += 1
        if (cardPos >= nbCardsMax) { cardPos = 0; }
    } else if (keyCode == RIGHT_ARROW) {
        cardPos -= 1
        if (cardPos < 0) { cardPos = nbCardsMax - 1; }
    }
    buildBottomCards();
}

function isCard(x, y, list) {
    console.log("Checking if (" + x + "," + y + ") is inside a card")
    console.log(list)
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
    var newCard = new Card(posX,
                           h - 185,
                           json.img,
                           json.text);
    append(cards, tmpCard);
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