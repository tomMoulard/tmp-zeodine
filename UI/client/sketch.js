var cards;
var mouseIsPressedHere;
var thisCardsID
var loadedCards;
var nbCards;
var nbCardsMax;
var cardPos;
var timeToSave;
var bottomCards;

function preload() {

    h                  = windowHeight;
    w                  = windowWidth;
    cards              = []; //Personal cards
    loadedCards        = []; //all cards provided
    bottomCards        = []; //cards at the bottom of the page
    mouseIsPressedHere = false;
    nbCards            = floor(w / 100); // max card in a row
    nbCardsMax         = 5; //total cards provided
    // TODO: get this from the api
    // nbCardsMax = loadJSON("http://147.135.194.248/nbcards/")
    cardPos            = 0; //Bottom first card
    timeToSave         = 0; //timer to save cards
    
    for (var i = 0; i < nbCardsMax; i++) {
        var url = "http://localhost:8080/card/" + i;
        
        var json = loadJSON(url, createCard)
    }
}

function setup() {
    createCanvas(w,h);
    buildBottomCards()
}

function createCard(json) {
    var posX = map(json.id, 0, w, 0, nbCards);
    var newCard = new Card(posX,
            h - 185,
            json.img,
            json.text);
    append(loadedCards, newCard);
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
    if (mouseIsPressedHere) {
        cards[thisCardsID].x = mouseX;
        cards[thisCardsID].y = mouseY;
        if (cards[thisCardsID].id != undefined) {
            cards[thisCardsID].update(mouseY)
        }
        // var tmpPos = isCard(mouseX, mouseY, cards)
        // if (tmpPos != -1){
        //     //there is a card below the one you are draging
        // }
    }
    if (timeToSave > 1000) {        // var req1 = new XMLHttpRequest();
    	// req1.open("GET",url,true);
    	// req1.addEventListener("load",createCard);
    	// req1.addEventListener("error",erreur);
        // req1.send(null);
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

function fusionCard(card1,card2) {
    if (card1.id != undefined && card2.id != undefined) {
        var grFus = new Group(0,card1,card2);
        cards.splice(cards.indexOf(card1), 1);
        cards.splice(cards.indexOf(card2), 1);
        append(cards,gr);
        mouseIsPressedHere = false;
    } else if (card1.id == undefined && card2.id == undefined) {
        var gr = new Group(0,card1,card2);
        cards.splice(cards.indexOf(card1), 1);
        cards.splice(cards.indexOf(card2), 1);
        append(cards,gr);
        mouseIsPressedHere = false;
    } else if (card1.id != undefined) {
        card1.appendCard(card2)
        cards.splice(cards.indexOf(card2), 1);
        mouseIsPressedHere = false;
    } else if (card2.id != undefined) {
        card2.appendCard(card1)
        cards.splice(cards.indexOf(card1), 1);
        mouseIsPressedHere = false;
    }

}

function cardClose() {
    for (var i = 0; i < cards.length; i++) {
        for (var j = i+1; j < cards.length; j++) {
            if (Math.abs(cards[i].x-cards[j].x)<20 &&
                Math.abs(cards[i].y-cards[j].y)<20) {
                    fusionCard(cards[i],cards[j])                    
            }
        }
    }
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
        var tmpCard = new Card(w / 2,
                               h / 2,
                               bottomCards[cardID].imgURL,
                               bottomCards[cardID].desc)
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
    // console.log("Checking if (" + x + "," + y + ") is inside a card")
    //  console.log(list)

    for (var i = 0; i < list.length; i++) {
        if (x > list[i].x - (list[i].w/2) &&
            x < list[i].x + (list[i].w/2) &&
            y > list[i].y - (list[i].h/2) &&
            y < list[i].y + (list[i].h/2))
            return i;
    }
    return -1;
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