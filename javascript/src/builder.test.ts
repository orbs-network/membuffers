/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import "./matcher-extensions";
import { ch } from "./text";
import { InternalBuilder } from "./builder";

test("TestBuilderUint8", () => {
  const w = new InternalBuilder();
  const v = 0x17;
  w.writeUint8(null, v);
  expect(w.size).toBe(1);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint8(buf, v);
  const expected = new Uint8Array([0x17]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint16", () => {
  const w = new InternalBuilder();
  const v = 0x17;
  w.writeUint16(null, v);
  expect(w.size).toBe(2);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint16(buf, v);
  const expected = new Uint8Array([0x17, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint32", () => {
  const w = new InternalBuilder();
  const v = 0x17;
  w.writeUint32(null, v);
  expect(w.size).toBe(4);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint32(buf, v);
  const expected = new Uint8Array([0x17, 0x00, 0x00, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint64", () => {
  const w = new InternalBuilder();
  const v = BigInt(0x17);
  w.writeUint64(null, v);
  expect(w.size).toBe(8);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint64(buf, v);
  const expected = new Uint8Array([0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint256", () => {
    const w = new InternalBuilder();
    const v = BigInt("0x2030405060708090a0b0c0d0e0e0f0102030405060708090a0b0c0d0e0e0f");
    w.writeUint256(null, v);
    expect(w.size).toBe(32);
    const buf = new Uint8Array(w.size);
    w.reset();
    w.writeUint256(buf, v);
    const expected = new Uint8Array([0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf]);
    expect(buf.subarray(0, 32)).toBeEqualToUint8Array(expected);

    // test to write 2 bytes32
    const buf2 = new Uint8Array(64);
    w.reset();
    w.writeUint256(buf2, v);
    w.writeUint256(buf2, v);
    const expected2 = new Uint8Array([
        0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf]);
    expect(buf2).toBeEqualToUint8Array(expected2);

    const vTooBig = BigInt("0x10012030405060708090a0b0c0d0e0e0f0102030405060708090a0b0c0d0e0e0f");
    w.reset();
    expect(() => {
        w.writeUint256(buf, vTooBig);
    }).toThrow();
});

test("TestBuilderBytes", () => {
  const w = new InternalBuilder();
  const v = new Uint8Array([0x01, 0x02, 0x03]);
  w.writeBytes(null, v);
  expect(w.size).toBe(7);
  w.writeBytes(null, v);
  expect(w.size).toBe(15);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeBytes(buf, v);
  const expected = new Uint8Array([0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03]);
  expect(buf.subarray(0, 7)).toBeEqualToUint8Array(expected);
  w.writeBytes(buf, v);
  const expected2 = new Uint8Array([0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03]);
  expect(buf).toBeEqualToUint8Array(expected2);
});

test("TestBuilderBytes20", () => {
    const w = new InternalBuilder();
    const v = new Uint8Array([0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13]);
    w.writeBytes20(null, v);
    expect(w.size).toBe(20);
    const buf = new Uint8Array(w.size);
    w.reset();
    w.writeBytes20(buf, v);
    const expected = new Uint8Array([0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13]);
    expect(buf.subarray(0, 20)).toBeEqualToUint8Array(expected);

    // test to write 2 bytes20
    const buf2 = new Uint8Array(40);
    w.reset();
    w.writeBytes20(buf2, v);
    w.writeBytes20(buf2, v);
    const expected2 = new Uint8Array([0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13]);
    expect(buf2).toBeEqualToUint8Array(expected2);
});

test("TestBuilderBytes32", () => {
    const w = new InternalBuilder();
    const v = new Uint8Array([0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf]);
    w.writeBytes32(null, v);
    expect(w.size).toBe(32);
    const buf = new Uint8Array(w.size);
    w.reset();
    w.writeBytes32(buf, v);
    const expected = new Uint8Array([0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf]);
    expect(buf.subarray(0, 32)).toBeEqualToUint8Array(expected);

    // test to write 2 bytes32
    const buf2 = new Uint8Array(64);
    w.reset();
    w.writeBytes32(buf2, v);
    w.writeBytes32(buf2, v);
    const expected2 = new Uint8Array([
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf]);
    expect(buf2).toBeEqualToUint8Array(expected2);
});

test("TestBuilderString", () => {
  const w = new InternalBuilder();
  const v = "hello";
  w.writeString(null, v);
  expect(w.size).toBe(9);
  w.writeString(null, v);
  expect(w.size).toBe(21);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeString(buf, v);
  const expected = new Uint8Array([0x05, 0x00, 0x00, 0x00, ch("h"), ch("e"), ch("l"), ch("l"), ch("o")]);
  expect(buf.subarray(0, 9)).toBeEqualToUint8Array(expected);
  w.writeString(buf, v);
  const expected2 = new Uint8Array([0x05, 0x00, 0x00, 0x00, ch("h"), ch("e"), ch("l"), ch("l"), ch("o"), 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, ch("h"), ch("e"), ch("l"), ch("l"), ch("o")]);
  expect(buf).toBeEqualToUint8Array(expected2);
});

test("TestBuilderUnionIndex", () => {
  const w = new InternalBuilder();
  const v = 0x01;
  w.writeUnionIndex(null, v);
  expect(w.size).toBe(2);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUnionIndex(buf, v);
  const expected = new Uint8Array([0x01, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint8Array", () => {
  const w = new InternalBuilder();
  const v = [0x01, 0x02, 0x03];
  w.writeUint8Array(null, v);
  expect(w.size).toBe(7);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint8Array(buf, v);
  const expected = new Uint8Array([0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint16Array", () => {
  const w = new InternalBuilder();
  const v = [0x01, 0x02, 0x03];
  w.writeUint16Array(null, v);
  expect(w.size).toBe(10);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint16Array(buf, v);
  const expected = new Uint8Array([0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x00, 0x03, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint32Array", () => {
  const w = new InternalBuilder();
  const v = [0x01, 0x02, 0x03];
  w.writeUint32Array(null, v);
  expect(w.size).toBe(16);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint32Array(buf, v);
  const expected = new Uint8Array([0x0c, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint64Array", () => {
  const w = new InternalBuilder();
  const v = [BigInt(0x01), BigInt(0x02), BigInt(0x03)];
  w.writeUint64Array(null, v);
  expect(w.size).toBe(28);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint64Array(buf, v);
  const expected = new Uint8Array([0x18, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderUint256Array", () => {
    const w = new InternalBuilder();
    const v = [BigInt("0x2030405060708090a0b0c0d0e0e0f0102030405060708090a0b0c0d0e0e0f"), BigInt("0x102030405060708090a0b0c0d0e0e0f0102030405060708090a0b0c0d0e0eff"), BigInt("0x1002030405060708090a0b0c0d0e0e0f0102030405060708090a0b0c0d0e0eee")];
    w.writeUint256Array(null, v);
    expect(w.size).toBe(4+32*3);
    const buf = new Uint8Array(w.size);
    w.reset();
    w.writeUint256Array(buf, v);
    const expected = new Uint8Array([0x60, 0x00, 0x00, 0x00,
        0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xff,
        0x10, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xee]);
    expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderBytesArray", () => {
  const w = new InternalBuilder();
  const v = [new Uint8Array([0x01, 0x02, 0x03]), new Uint8Array([0x04, 0x05])];
  w.writeBytesArray(null, v);
  expect(w.size).toBe(18);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeBytesArray(buf, v);
  const expected = new Uint8Array([0x0e, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x00, 0x02, 0x00, 0x00, 0x00, 0x04, 0x05]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderBytes20Array", () => {
  const w = new InternalBuilder();
  const v = [
      new Uint8Array([0xcc, 0xdd, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13]),
      new Uint8Array([0xaa, 0xdd, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0xcc])];
  w.writeBytes20Array(null, v);
  expect(w.size).toBe(44);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeBytes20Array(buf, v);
  const expected = new Uint8Array([0x28, 0x00, 0x00, 0x00,
      0xcc, 0xdd, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13,
      0xaa, 0xdd, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf, 0x10, 0x11, 0x12, 0xcc]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderBytes32Array", () => {
    const w = new InternalBuilder();
    const v = [
        new Uint8Array([0xcc, 0xdd, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
            0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xcc, 0xdd]),
        new Uint8Array([0xaa, 0xbb, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
            0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xbb, 0xaa])];
    w.writeBytes32Array(null, v);
    expect(w.size).toBe(68);
    const buf = new Uint8Array(w.size);
    w.reset();
    w.writeBytes32Array(buf, v);
    const expected = new Uint8Array([0x40, 0x00, 0x00, 0x00,
        0xcc, 0xdd, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xcc, 0xdd,
        0xaa, 0xbb, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
        0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xbb, 0xaa]);
    expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderStringArray", () => {
  const w = new InternalBuilder();
  const v = ["jay", "lo"];
  w.writeStringArray(null, v);
  expect(w.size).toBe(18);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeStringArray(buf, v);
  const expected = new Uint8Array([0x0e, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, ch("j"), ch("a"), ch("y"), 0x00, 0x02, 0x00, 0x00, 0x00, ch("l"), ch("o")]);
  expect(buf).toBeEqualToUint8Array(expected);
});

class ExampleMessageBuilder {
  private builder: InternalBuilder;

  constructor() {
    this.builder = new InternalBuilder();
  }
  write(buf: Uint8Array) {
    this.builder.reset();
    this.builder.writeUint8(buf, 0x17);
    this.builder.writeUint32(buf, 0x033);
  }
  getSize() {
    return this.builder.getSize();
  }
  calcRequiredSize() {
    this.write(null);
    return this.builder.getSize();
  }
}

test("TestBuilderMessage", () => {
  const w = new InternalBuilder();
  const v = new ExampleMessageBuilder();
  w.writeMessage(null, v);
  expect(w.size).toBe(12);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeMessage(buf, v);
  const expected = new Uint8Array([0x08, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test("TestBuilderMessageArray", () => {
  const w = new InternalBuilder();
  const v = [new ExampleMessageBuilder(), new ExampleMessageBuilder()];
  w.writeMessageArray(null, v);
  expect(w.size).toBe(28);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeMessageArray(buf, v);
  const expected = new Uint8Array([0x18, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});
