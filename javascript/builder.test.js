import './matcher-extensions';
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