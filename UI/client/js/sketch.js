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
var openPanel;
var json;
var ws;

function preload() {
    openPanel = false;
    h = windowHeight;
    w = windowWidth;
    cards = []; //Personal cards
    loadedCards = []; //all cards provided
    bottomCards = []; //cards at the bottom of the page
    mouseIsPressedHere = false;
    doubleClick = false;
    nbCards = floor(w / 100); // max card in a row
    nbCardsMax = 0; //total cards provided
    cardPos = 0; //Bottom first card
    timeToSave = 0; //timer to save cards
}

function setup() {
    console.log(userid, wsid)
    createCanvas(w, h);
    var postData = { user_id: userid, ws_id: wsid };
    var req = new XMLHttpRequest();
    req.open("POST", "http://147.135.194.248/load/", true);

    req.addEventListener("load", loadCard);
    req.addEventListener("error", errorFunction);
    req.setRequestHeader("Content-Type", "application/json");
    req.send(JSON.stringify(postData));
}


function windowResized() {
    resizeCanvas(windowWidth, windowHeight);
    h = windowHeight;
    w = windowWidth;
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
    if (timeToSave > 1000) {
        var crd = { card_content: "too" };
        var crds = { card_pub: crdp, card_id: crd_id, card: crd };
        var grs = { groupe_id: gr_id, cards: crds };
        var grs = [];
        var postData = { user_id: userid, ws_id: wsid, groupes: grs };

        for (var i = 0; i < cards.length; ++i) {
            append(grs, cards[i].save());
        }

        var req1 = new XMLHttpRequest();
        req1.open("GET", "http://147.135.194.248/load/", true);
        req1.addEventListener("load", saveCard);
        req1.addEventListener("error", errorFunction);
        req.send(JSON.stringify(postData));

        `{"user_id": 1,"ws_id": 1,"groupes": 
        [{"groupe_id": 12,"cards": 
            [{"card_id": 1,
                "card": {"card_content": "{}"}},{"card_id": 2,"card": {"card_content": "{}"}},{"card_id": 3,"card": {"card_content": "{}"}},{"card_id": 4,"card": {"card_content": "{}"}}]},{"groupe_id": 21,"cards": [{"card_id": 5,"card": {"card_content": "{\"card_pos\":12}"}}
            ]}
        ]}`

        console.log("Saving ...");
        console.log({ grs });
        timeToSave = -1;
    }
    stroke(255, 0, 0);
    rectMode(CENTER);
    rect(50, 50, 20, 30);
    cardClose();
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

    if (mouseX > 50 - (10) &&
        mouseX < 50 + (10) &&
        mouseY > 50 - (15) &&
        mouseY < 50 + (15)) {

        if (openPanel) {
            resizeCanvas(windowWidth * 0.8, windowHeight);
            openPanel = false;
        } else {
            resizeCanvas(windowWidth, windowHeight);
            openPanel = true;
        }


        return;
    }

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