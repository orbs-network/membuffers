import './matcher-extensions';
import {ch} from './text';
import {InternalBuilder} from "./builder";

test('TestBuilderUint8', () => {
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

test('TestBuilderUint16', () => {
  const w = new InternalBuilder();
  const v = 0x17;
  w.writeUint16(null, v);
  expect(w.size).toBe(2);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint16(buf, v);
  const expected = new Uint8Array([0x17,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderUint32', () => {
  const w = new InternalBuilder();
  const v = 0x17;
  w.writeUint32(null, v);
  expect(w.size).toBe(4);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint32(buf, v);
  const expected = new Uint8Array([0x17,0x00,0x00,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderUint64', () => {
  const w = new InternalBuilder();
  const v = BigInt(0x17);
  w.writeUint64(null, v);
  expect(w.size).toBe(8);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint64(buf, v);
  const expected = new Uint8Array([0x17,0x00,0x00,0x00,0x00,0x00,0x00,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderBytes', () => {
  const w = new InternalBuilder();
  const v = new Uint8Array([0x01,0x02,0x03]);
  w.writeBytes(null, v);
  expect(w.size).toBe(7);
  w.writeBytes(null, v);
  expect(w.size).toBe(15);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeBytes(buf, v);
  const expected = new Uint8Array([0x03,0x00,0x00,0x00, 0x01,0x02,0x03]);
  expect(buf.subarray(0,7)).toBeEqualToUint8Array(expected);
  w.writeBytes(buf, v);
  const expected2 = new Uint8Array([0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,]);
  expect(buf).toBeEqualToUint8Array(expected2);
});

test('TestBuilderString', () => {
  const w = new InternalBuilder();
  const v = "hello";
  w.writeString(null, v);
  expect(w.size).toBe(9);
  w.writeString(null, v);
  expect(w.size).toBe(21);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeString(buf, v);
  const expected = new Uint8Array([0x05,0x00,0x00,0x00, ch('h'),ch('e'),ch('l'),ch('l'),ch('o')]);
  expect(buf.subarray(0,9)).toBeEqualToUint8Array(expected);
  w.writeString(buf, v);
  const expected2 = new Uint8Array([0x05,0x00,0x00,0x00, ch('h'),ch('e'),ch('l'),ch('l'),ch('o'),0x00,0x00,0x00, 0x05,0x00,0x00,0x00, ch('h'),ch('e'),ch('l'),ch('l'),ch('o')]);
  expect(buf).toBeEqualToUint8Array(expected2);
});

test('TestBuilderUnionIndex', () => {
  const w = new InternalBuilder();
  const v = 0x01;
  w.writeUnionIndex(null, v);
  expect(w.size).toBe(2);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUnionIndex(buf, v);
  const expected = new Uint8Array([0x01,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderUint8Array', () => {
  const w = new InternalBuilder();
  const v = [0x01,0x02,0x03];
  w.writeUint8Array(null, v);
  expect(w.size).toBe(7);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint8Array(buf, v);
  const expected = new Uint8Array([0x03,0x00,0x00,0x00, 0x01,0x02,0x03]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderUint16Array', () => {
  const w = new InternalBuilder();
  const v = [0x01,0x02,0x03];
  w.writeUint16Array(null, v);
  expect(w.size).toBe(10);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint16Array(buf, v);
  const expected = new Uint8Array([0x06,0x00,0x00,0x00, 0x01,0x00, 0x02,0x00, 0x03,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderUint32Array', () => {
  const w = new InternalBuilder();
  const v = [0x01,0x02,0x03];
  w.writeUint32Array(null, v);
  expect(w.size).toBe(16);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint32Array(buf, v);
  const expected = new Uint8Array([0x0c,0x00,0x00,0x00, 0x01,0x00,0x00,0x00, 0x02,0x00,0x00,0x00, 0x03,0x00,0x00,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderUint64Array', () => {
  const w = new InternalBuilder();
  const v = [BigInt(0x01),BigInt(0x02),BigInt(0x03)];
  w.writeUint64Array(null, v);
  expect(w.size).toBe(28);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeUint64Array(buf, v);
  const expected = new Uint8Array([0x18,0x00,0x00,0x00, 0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x02,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x03,0x00,0x00,0x00,0x00,0x00,0x00,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderBytesArray', () => {
  const w = new InternalBuilder();
  const v = [new Uint8Array([0x01,0x02,0x03]),new Uint8Array([0x04,0x05])];
  w.writeBytesArray(null, v);
  expect(w.size).toBe(18);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeBytesArray(buf, v);
  const expected = new Uint8Array([0x0e,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x02,0x00,0x00,0x00, 0x04,0x05]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderStringArray', () => {
  const w = new InternalBuilder();
  const v = ["jay","lo"];
  w.writeStringArray(null, v);
  expect(w.size).toBe(18);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeStringArray(buf, v);
  const expected = new Uint8Array([0x0e,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, ch('j'),ch('a'),ch('y'),0x00, 0x02,0x00,0x00,0x00, ch('l'),ch('o')]);
  expect(buf).toBeEqualToUint8Array(expected);
});

class ExampleMessageBuilder {
  constructor() {
    this.builder = new InternalBuilder();
  }
  write(buf) {
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

test('TestBuilderMessage', () => {
  const w = new InternalBuilder();
  const v = new ExampleMessageBuilder();
  w.writeMessage(null, v);
  expect(w.size).toBe(12);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeMessage(buf, v);
  const expected = new Uint8Array([0x08,0x00,0x00,0x00, 0x17,0x00,0x00,0x00, 0x33,0x00,0x00,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});

test('TestBuilderMessageArray', () => {
  const w = new InternalBuilder();
  const v = [new ExampleMessageBuilder(), new ExampleMessageBuilder()];
  w.writeMessageArray(null, v);
  expect(w.size).toBe(28);
  const buf = new Uint8Array(w.size);
  w.reset();
  w.writeMessageArray(buf, v);
  const expected = new Uint8Array([0x18,0x00,0x00,0x00, 0x08,0x00,0x00,0x00, 0x17,0x00,0x00,0x00, 0x33,0x00,0x00,0x00, 0x08,0x00,0x00,0x00, 0x17,0x00,0x00,0x00, 0x33,0x00,0x00,0x00]);
  expect(buf).toBeEqualToUint8Array(expected);
});