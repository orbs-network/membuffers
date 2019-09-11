/**
 * Copyright 2018 the membuffers authors
 * This file is part of the membuffers library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

let textEncoder = null;
function getTextEncoder() {
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
function getTextDecoder() {
  if (textDecoder === null) {
    if (typeof TextDecoder === "undefined") { // node.js does not support TextDecoder
      textDecoder = new (require('text-encoding').TextDecoder)("utf-8");
    } else {
      textDecoder = new TextDecoder("utf-8");
    }
  }
  return textDecoder;
}

function ch(char) {
  return char.charCodeAt(0);
}

module.exports = {getTextDecoder, getTextEncoder, ch};
