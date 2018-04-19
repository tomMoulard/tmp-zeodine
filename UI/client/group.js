function Group(id, card1) {
    this.id = id;
    this.x = card1.x;
    this.y = card1.y;
    this.theCards = [];
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