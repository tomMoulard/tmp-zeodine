function Group(id,card1,card2) {
    this.id = id;
    this.x = card1.x;
    this.y = card1.y;
    this.w = 90;
    this.h = 180;
    card2.x = card1.x+10;
    card2.y= card1.y;
    this.theCards = [];
    
    append(this.theCards,card1);
    append(this.theCards,card2);

    this.display = function() {
        for (var i=0;i<this.theCards.length ;i++) {
            this.theCards[i].display()
        }
    }

    this.appendCard = function(card) {
        card.y=this.y
        append(this.theCards,card);
        var l = this.theCards.length

        for (var i=0;i<l ;i++) {
            this.theCards[i].x=this.x-15+((20/l)*i)
        }
    }

    this.update = function(newY) {
        var l = this.theCards.length

        for (var i=0;i<l ;i++) {
            this.theCards[i].x=this.x-15+((20/l)*i);
            this.theCards[i].y=newY;
        }
    }

    this.save = function(){
        return {
            "x": this.x,
            "y": this.y,
        }
    }
    
}