/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

let textEncoder: TextEncoder = null;
export function getTextEncoder(): TextEncoder {
  if (textEncoder === null) {
    if (typeof TextEncoder === "undefined") {
      // node.js does not support TextEncoder
      require("fast-text-encoding");
    }
    textEncoder = new TextEncoder();
  }
  return textEncoder;
}

let textDecoder: TextDecoder = null;
export function getTextDecoder(): TextDecoder {
  if (textDecoder === null) {
    if (typeof TextDecoder === "undefined") {
      // node.js does not support TextDecoder
      require("fast-text-encoding");
    }
    textDecoder = new TextDecoder("utf-8");
  }
  return textDecoder;
}

export function ch(char: string): number {
  return char.charCodeAt(0);
}
