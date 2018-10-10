import './matcher-extensions';
import {FieldTypes} from './types';
import {ch} from './text';
import {InternalMessage} from "./message";

test('TestIteratorUint32', () => {
  const tests = [
    {
      buf: new Uint8Array([0x0c,0x00,0x00,0x00, 0x13,0x00,0x00,0x00, 0x14,0x00,0x00,0x00, 0x15,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeUint32Array],
      unions: [],
      expected: [0x13,0x14,0x15],
    },
    {
      buf: new Uint8Array([0x08,0x00,0x00,0x00, 0x88,0x00,0x00,0x00, 0x11,0x22,0x33]),
      scheme: [FieldTypes.TypeUint32Array],
      unions: [],
      expected: [],
    },
    {
      buf: new Uint8Array([0x07,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x11,0x22,0x33]),
      scheme: [FieldTypes.TypeUint32Array],
      unions: [],
      expected: [0x05, 0x00],
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const iterator = m.getUint32ArrayIterator(0);
    const res = [];
    for (const v of iterator) {
      res.push(v);
    }
    // console.log(tt); // uncomment on failure to find out where
    expect(res).toEqual(tt.expected);
  }
});

test('TestIteratorUint8', () => {
  const tests = [
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x13, 0x14, 0x15]),
      scheme: [FieldTypes.TypeUint8Array],
      unions: [],
      expected: [0x13,0x14,0x15],
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const iterator = m.getUint8ArrayIterator(0);
    const res = [];
    for (const v of iterator) {
      res.push(v);
    }
    // console.log(tt); // uncomment on failure to find out where
    expect(res).toEqual(tt.expected);
  }
});

test('TestIteratorUint16', () => {
  const tests = [
    {
      buf: new Uint8Array([0x06,0x00,0x00,0x00, 0x13,0x00, 0x14,0x00, 0x15,0x00]),
      scheme: [FieldTypes.TypeUint16Array],
      unions: [],
      expected: [0x13,0x14,0x15],
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const iterator = m.getUint16ArrayIterator(0);
    const res = [];
    for (const v of iterator) {
      res.push(v);
    }
    // console.log(tt); // uncomment on failure to find out where
    expect(res).toEqual(tt.expected);
  }
});

test('TestIteratorUint64', () => {
  const tests = [
    {
      buf: new Uint8Array([0x18,0x00,0x00,0x00, 0x13,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x14,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x15,0x00,0x00,0x00,0x00,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeUint64Array],
      unions: [],
      expected: [BigInt(0x13),BigInt(0x14),BigInt(0x15)],
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const iterator = m.getUint64ArrayIterator(0);
    const res = [];
    for (const v of iterator) {
      res.push(v);
    }
    // console.log(tt); // uncomment on failure to find out where
    expect(res).toEqual(tt.expected);
  }
});

test('TestIteratorMessage', () => {
  const tests = [
    {
      buf: new Uint8Array([0x1b,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03]),
      scheme: [FieldTypes.TypeMessageArray],
      unions: [],
      expectedSizes: [0x03,0x05,0x03],
    },
    {
      buf: new Uint8Array([0x08,0x00,0x00,0x00, 0x88,0x00,0x00,0x00, 0x11,0x22,0x33,0x44]),
      scheme: [FieldTypes.TypeMessageArray],
      unions: [],
      expectedSizes: [0],
    },
    {
      buf: new Uint8Array([0x08,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x11,0x22,0x33,0x44]),
      scheme: [FieldTypes.TypeMessageArray],
      unions: [],
      expectedSizes: [0],
    },
    {
      buf: new Uint8Array([0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44]),
      scheme: [FieldTypes.TypeMessageArray],
      unions: [],
      expectedSizes: [],
    },
    {
      buf: new Uint8Array([0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44,0x55]),
      scheme: [FieldTypes.TypeMessageArray],
      unions: [],
      expectedSizes: [4,0],
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const iterator = m.getMessageArrayIterator(0);
    const res = [];
    for (const [_, size] of iterator) {
      res.push(size);
    }
    // console.log(tt); // uncomment on failure to find out where
    expect(res).toEqual(tt.expectedSizes);
  }
});

test('TestIteratorBytes', () => {
  const tests = [
    {
      buf: new Uint8Array([0x1b,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03]),
      scheme: [FieldTypes.TypeBytesArray],
      unions: [],
      expected: [new Uint8Array([0x01,0x02,0x03]),new Uint8Array([0x01,0x02,0x03,0x04,0x05]),new Uint8Array([0x01,0x02,0x03])],
    },
    {
      buf: new Uint8Array([0x08,0x00,0x00,0x00, 0x88,0x00,0x00,0x00, 0x11,0x22,0x33,0x44]),
      scheme: [FieldTypes.TypeBytesArray],
      unions: [],
      expected: [new Uint8Array()],
    },
    {
      buf: new Uint8Array([0x08,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x11,0x22,0x33,0x44]),
      scheme: [FieldTypes.TypeBytesArray],
      unions: [],
      expected: [new Uint8Array()],
    },
    {
      buf: new Uint8Array([0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44]),
      scheme: [FieldTypes.TypeBytesArray],
      unions: [],
      expected: [],
    },
    {
      buf: new Uint8Array([0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44,0x55]),
      scheme: [FieldTypes.TypeBytesArray],
      unions: [],
      expected: [new Uint8Array([0x11,0x22,0x33,0x44]),new Uint8Array()],
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const iterator = m.getBytesArrayIterator(0);
    const res = [];
    for (const v of iterator) {
      res.push(v);
    }
    // console.log(tt); // uncomment on failure to find out where
    expect(res).toEqual(tt.expected);
  }
});

test('TestIteratorString', () => {
  const tests = [
    {
      buf: new Uint8Array([0x1b,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c'),0x00, 0x05,0x00,0x00,0x00, ch('h'),ch('e'),ch('l'),ch('l'), ch('o'),0x00,0x00,0x00, 0x03,0x00,0x00,0x00, ch('d'),ch('e'),ch('f')]),
      scheme: [FieldTypes.TypeStringArray],
      unions: [],
      expected: ["abc","hello","def"],
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const iterator = m.getStringArrayIterator(0);
    const res = [];
    for (const v of iterator) {
      res.push(v);
    }
    // console.log(tt); // uncomment on failure to find out where
    expect(res).toEqual(tt.expected);
  }
});