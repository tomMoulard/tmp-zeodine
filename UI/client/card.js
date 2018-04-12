function Card(x, y, imgURL, desc, show) {
    this.x = x;
    this.y = y;
    this.w = 90;
    this.h = 180;
    this.imgURL = imgURL;
    this.img = loadImage(this.imgURL);
    this.desc = desc;
    this.show = show
    // this.ID = ID;

    this.display = function() {
        noStroke()
        rectMode(CENTER)
        rect(this.x, this.y, this.w, this.h);
        image(this.img, //imgurl
            this.x - this.w / 2 + 4, // x pos
            this.y - this.h / 2 + 4, // y pos
            this.w - 8, //width of the pict
            this.h / 2 - 8) // height of the pict
        text(this.desc, // text to display
            this.x + 4, // x pos
            this.y + this.h / 2, // y pos
            this.w,
            this.h)
    }

    this.save = function(){
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