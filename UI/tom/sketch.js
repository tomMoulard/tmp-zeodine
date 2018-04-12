var cards;
function setup() {
    h      = windowHeight;
    w      = windowWidth;
    cards  = [];
    createCanvas(w, h);
}

function draw() {
    background(51);
    for (var i = 0; i < cards.length; i++) {
        cards[i].display()
    }
}

function mousePressed() {
    if (mouseX < w && mouseY < h) {
        console.log(mouseX, mouseY)
        var newCard = new Card(mouseX, mouseY)
        append(cards, newCard)
    }
}