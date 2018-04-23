function Card(x, y, imgURL, desc) {
    this.x = x;
    this.y = y;
    this.w = 90;
    this.h = 180;
    this.imgURL = imgURL;
    this.img = loadImage(this.imgURL);
    this.desc = desc;
    // this.ID = ID;
    this.div = document.createElement("div");
    this.div.className = "cardInfo"


    this.display = function(color) {
        stroke(color[0], color[1], color[2]);
        if (color[0] != 0 && color[1] != 0 && color[2] != 0) {
            strokeWeight(3);
        } else {
            strokeWeight(1);
        }

        rectMode(CENTER)
        rect(this.x, this.y, this.w, this.h);
        image(this.img, //imgurl
                this.x - this.w / 2 + 4, // x pos
                this.y - this.h / 2 + 4, // y pos
                this.w - 8, //width of the pict
                this.h / 2 - 8) // height of the pict
        noStroke();
        text(this.desc, // text to display
            this.x + 4, // x pos
            this.y + this.h / 2, // y pos
            this.w,
            this.h)
    }

    this.globalView = function(index, color) {
        var min_wh = Math.min(w, h);
        stroke(color[0], color[1], color[2]);
        strokeWeight(6);
        rectMode(CORNER)
        rect((min_wh / 10), (min_wh / 10), min_wh / 3, (2 * min_wh) / 3);
        image(this.img, //imgurl
                (min_wh / 10) + 4, // x pos
                (min_wh / 10) + 4, // y pos
                min_wh / 3 - 8, //width of the pict
                min_wh / 3 - 8) // height of the pict
        noStroke();
        text(this.desc, // text to display
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
        cards.splice(cards.indexOf(this), 1);
    }

    this.check = function(xmouse, ymouse) {
        if (xmouse > this.x - (this.w / 2) &&
            xmouse < this.x + (this.w / 2) &&
            ymouse > this.y - (this.h / 2) &&
            ymouse < this.y + (this.h / 2)) {
            return [true, -1];
        }
        return [false, -1];
    }

    this.move = function(xnew, ynew, ind) {
        this.x = xnew;
        this.y = ynew;
    }

    this.save = function() {
        return {
            "x": this.x,
            "y": this.y,
            "w": this.w,
            "h": this.h,
            "img": this.imgURL,
            "text": this.desc
        }
    }
}