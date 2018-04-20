function Group(id, card1) {
    this.id = id;
    this.x = card1.x;
    this.y = card1.y;
    this.theCards = [];
    this.div = document.createElement("div");
    this.div.className = "cardInfo"


    if (card1.theCards != undefined) {
        for (var i = 0; i < card1.theCards.length; i++) {
            append(this.theCards, card1.theCards[i]);
        }
        this.color = card1.color;
    } else {
        append(this.theCards, card1);
        this.color = [Math.floor(Math.random() * 255), Math.floor(Math.random() * 255), Math.floor(Math.random() * 255)];
    }

    this.display = function(color) {
        for (var i = 0; i < this.theCards.length; i++) {
            if (i < this.theCards.length - 1) {
                stroke(this.color[0], this.color[1], this.color[2]);
                strokeWeight(2);
                line(this.theCards[i].x, this.theCards[i].y, this.theCards[i + 1].x, this.theCards[i + 1].y);
            }
            this.theCards[i].display(this.color)
        }
    }

    this.globalView = function(index, color) {
        var min_wh = Math.min(w, h);
        stroke(this.color[0], this.color[1], this.color[2]);
        strokeWeight(10);
        rectMode(CORNER)
        rect((min_wh / 10), (min_wh / 10), min_wh / 3, (2 * min_wh) / 3);
        image(this.theCards[index].img, //imgurl
                (min_wh / 10) + 4, // x pos
                (min_wh / 10) + 4, // y pos
                min_wh / 3 - 8, //width of the pict
                min_wh / 3 - 8) // height of the pict
        noStroke();
        text(this.theCards[index].desc, // text to display
            (min_wh / 10) + 4, // x pos
            (min_wh / 10) + min_wh / 3, // y pos
            min_wh / 3,
            min_wh)

        document.body.appendChild(this.div)
    }

    this.rmvChild = function() {
        document.body.removeChild(this.div)
    }

    this.removeCard = function(index) {
        var card = this.theCards[index]
        if (this.theCards.length == 1) {
            cards.splice(cards.indexOf(this), 1);
            return;
        }
        append(cards, card);
        this.theCards.splice(this.theCards.indexOf(card), 1);
    }

    this.appendCard = function(card) {
        if (card.theCards != undefined) {
            for (var i = 0; i < card.theCards.length; i++) {
                append(this.theCards, card.theCards[i]);
            }
        } else {
            append(this.theCards, card);
        }
    }

    this.check = function(xmouse, ymouse) {
        var res = [false, -1];
        for (var i = 0; i < this.theCards.length; i++) {
            res = this.theCards[i].check(xmouse, ymouse)
            if (res[0]) {
                return [true, i];
            }
        }
        return [false, -1];
    }

    this.move = function(xnew, ynew, ind) {
        this.theCards[ind].move(xnew, ynew, -1);
        this.x = this.theCards[0].x;
        this.y = this.theCards[0].y;
    }

    this.save = function() {
        return {
            "x": this.x,
            "y": this.y,
        }
    }

}