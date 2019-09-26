/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

const fromCharCode = String.fromCharCode;
const Object_prototype_toString = ({}).toString;
const log = Math.log;
const LN2 = Math.LN2;
const clz32 = Math.clz32 || function(x) {return 31 - log(x >>> 0) / LN2 | 0};

function encoderReplacer(nonAsciiChars: string){
  // make the UTF string into a binary UTF-8 encoded string
  var point = nonAsciiChars.charCodeAt(0)|0;
  if (point >= 0xD800 && point <= 0xDBFF) {
    var nextcode = nonAsciiChars.charCodeAt(1)|0;
    if (nextcode !== nextcode) // NaN because string is 1 code point long
      return fromCharCode(0xef/*11101111*/, 0xbf/*10111111*/, 0xbd/*10111101*/);
    // https://mathiasbynens.be/notes/javascript-encoding#surrogate-formulae
    if (nextcode >= 0xDC00 && nextcode <= 0xDFFF) {
      point = ((point - 0xD800)<<10) + nextcode - 0xDC00 + 0x10000|0;
      if (point > 0xffff)
        return fromCharCode(
          (0x1e/*0b11110*/<<3) | (point>>>18),
          (0x2/*0b10*/<<6) | ((point>>>12)&0x3f/*0b00111111*/),
          (0x2/*0b10*/<<6) | ((point>>>6)&0x3f/*0b00111111*/),
          (0x2/*0b10*/<<6) | (point&0x3f/*0b00111111*/)
        );
    } else return fromCharCode(0xef, 0xbf, 0xbd);
  }
  if (point <= 0x007f) return nonAsciiChars;
  else if (point <= 0x07ff) {
    return fromCharCode((0x6<<5)|(point>>>6), (0x2<<6)|(point&0x3f));
  } else return fromCharCode(
    (0xe/*0b1110*/<<4) | (point>>>12),
    (0x2/*0b10*/<<6) | ((point>>>6)&0x3f/*0b00111111*/),
    (0x2/*0b10*/<<6) | (point&0x3f/*0b00111111*/)
  );
}
class TextEncoder {
  constructor() {

  }

  encode(inputString: string){
    // 0xc0 => 0b11000000; 0xff => 0b11111111; 0xc0-0xff => 0b11xxxxxx
    // 0x80 => 0b10000000; 0xbf => 0b10111111; 0x80-0xbf => 0b10xxxxxx
    var encodedString = inputString === void 0 ?  "" : ("" + inputString).replace(/[\x80-\uD7ff\uDC00-\uFFFF]|[\uD800-\uDBFF][\uDC00-\uDFFF]?/g, encoderReplacer);
    var len=encodedString.length|0, result = new Uint8Array(len);
    for (var i=0; i<len; i=i+1|0)
      result[i] = encodedString.charCodeAt(i)|0;
    return result;
  };
}

function decoderReplacer(encoded: string) {
  var codePoint = encoded.charCodeAt(0) << 24;
  var leadingOnes = clz32(~codePoint)|0;
  var endPos = 0, stringLen = encoded.length|0;
  var result = "";
  if (leadingOnes < 5 && stringLen >= leadingOnes) {
    codePoint = (codePoint<<leadingOnes)>>>(24+leadingOnes);
    for (endPos = 1; endPos < leadingOnes; endPos=endPos+1|0)
      codePoint = (codePoint<<6) | (encoded.charCodeAt(endPos)&0x3f/*0b00111111*/);
    if (codePoint <= 0xFFFF) { // BMP code point
      result += fromCharCode(codePoint);
    } else if (codePoint <= 0x10FFFF) {
      // https://mathiasbynens.be/notes/javascript-encoding#surrogate-formulae
      codePoint = codePoint - 0x10000|0;
      result += fromCharCode(
        (codePoint >> 10) + 0xD800|0,  // highSurrogate
        (codePoint & 0x3ff) + 0xDC00|0 // lowSurrogate
      );
    } else endPos = 0; // to fill it in with INVALIDs
  }
  for (; endPos < stringLen; endPos=endPos+1|0) result += "\ufffd"; // replacement character
  return result;
}

class TextDecoder {
  constructor(encoding: string) {

  }

  decode(inputArrayOrBuffer: Uint8Array){
    var buffer = inputArrayOrBuffer.buffer;
    var asString = Object_prototype_toString.call(buffer);
    var inputAs8 = inputArrayOrBuffer;
    var resultingString = "";
    for (var index=0,len=inputAs8.length|0; index<len; index=index+32768|0)
      resultingString += fromCharCode.apply(0, inputAs8["subarray"](index,index+32768|0));

    return resultingString.replace(/[\xc0-\xff][\x80-\xbf]*/g, decoderReplacer);
  }
}

let textEncoder: any = null;
export function getTextEncoder(): any {
  if (textEncoder === null) {
    textEncoder = new TextEncoder();
  }
  return textEncoder;
}

let textDecoder: any = null;
export function getTextDecoder(): any {
  if (textDecoder === null) {
    textDecoder = new TextDecoder("utf-8");
  }
  return textDecoder;
}

export function ch(char: string): number {
  return char.charCodeAt(0);
}
