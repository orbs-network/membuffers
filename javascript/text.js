let textEncoder = null;
export function getTextEncoder() {
  if (textEncoder === null) {
    if (typeof TextEncoder === "undefined") { // node.js does not support TextEncoder
      textEncoder = new (require('text-encoding').TextEncoder)();
    } else {
      textEncoder = new TextEncoder();
    }
  }
  return textEncoder;
}

let textDecoder = null;
export function getTextDecoder() {
  if (textDecoder === null) {
    if (typeof TextDecoder === "undefined") { // node.js does not support TextDecoder
      textDecoder = new (require('text-encoding').TextDecoder)("utf-8");
    } else {
      textDecoder = new TextDecoder("utf-8");
    }
  }
  return textDecoder;
}

export function ch(char) {
  return char.charCodeAt(0);
}
