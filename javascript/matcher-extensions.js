/**
 * Copyright 2018 the membuffers authors
 * This file is part of the membuffers library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

function arrayBufferToHex(buff) {
  return Array.prototype.map.call(new Uint8Array(buff), x => ('00' + x.toString(16)).slice(-2)).join('');
}

function uint8ArrayToHex(arr) {
  return Array.prototype.map.call(arr, x => ('00' + x.toString(16)).slice(-2)).join('');
}

expect.extend({
  toBeEqualToArrayBuffer(received, other) {
    if (received.byteLength != other.byteLength) {
      return {
        message: () => `expected arrayBuffer length ${received.byteLength} to be equal ${other.byteLength}`,
        pass: false,
      };
    }
    received = new Uint8Array(received);
    other = new Uint8Array(other);
    for (let i = 0; i < received.byteLength; i++) {
      if (received[i] != other[i]) {
        return {
          message: () => `expected arrayBuffer ${arrayBufferToHex(received)} to equal ${arrayBufferToHex(other)}`,
          pass: false,
        };
      }
    }
    return {
      message: () => `expected arrayBuffer ${arrayBufferToHex(received)} not to equal ${arrayBufferToHex(other)}`,
      pass: true,
    };
  },
  toBeEqualToUint8Array(received, other) {
    if (received.byteLength != other.byteLength) {
      return {
        message: () => `expected arrayBuffer length ${received.byteLength} to be equal ${other.byteLength}`,
        pass: false,
      };
    }
    for (let i = 0; i < received.byteLength; i++) {
      if (received[i] != other[i]) {
        return {
          message: () => `expected uint8Array ${uint8ArrayToHex(received)} to equal ${uint8ArrayToHex(other)}`,
          pass: false,
        };
      }
    }
    return {
      message: () => `expected uint8Array ${uint8ArrayToHex(received)} not to equal ${uint8ArrayToHex(other)}`,
      pass: true,
    };
  },
});