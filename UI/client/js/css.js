var base_css;
window.addEventListener("load", cssUpdater);

function cssUpdater() {
    var rawFile = new XMLHttpRequest();
    rawFile.open("GET", "http://147.135.194.248/css/base.css", true);
    rawFile.onreadystatechange = function() { if (rawFile.readyState === 4) { addcss(rawFile.responseText) } }
    rawFile.send();
}

function addcss(css) {
    var head = document.getElementsByTagName('head')[0];
    var s = document.createElement('style');
    s.setAttribute('type', 'text/css');
    if (s.styleSheet) {
        s.styleSheet.cssText = css;
    } else {
        s.appendChild(document.createTextNode(css));
    }
    head.appendChild(s);
}