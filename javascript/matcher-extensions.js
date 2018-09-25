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