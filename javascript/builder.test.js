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