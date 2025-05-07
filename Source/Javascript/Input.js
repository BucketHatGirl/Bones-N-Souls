let KEYPRESS = [];

function KeyDown(event) {
  const KEY = event.key;
  if (!KEYPRESS.includes(KEY)) {
     KEYPRESS.push(KEY);
  }
}

function KeyUp(event) {
  const KEY = event.key;
  const INDEX = KEYPRESS.indexOf(KEY);
  if (INDEX > -1) {
    KEYPRESS.splice(INDEX, 1);
  }
}

document.addEventListener('keydown', KeyDown);
document.addEventListener('keyup', KeyUp);

function FetchKeys() {
  return KEYPRESS;
}


