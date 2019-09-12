declare namespace jest {
  interface Matchers<R> {
    toBeEqualToUint8Array(expected: Uint8Array): R;
  }
}
