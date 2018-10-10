import './matcher-extensions';
import {FieldTypes} from './types';
import {ch} from './text';
import {InternalMessage} from './message';

test('TestMessageRawBuffer', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      expected: new Uint8Array([0x00,0x00,0x00,0x00]),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c')]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      expected: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c')]),
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.rawBuffer();
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageRawBufferForField', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00, 0x00, 0x00, 0x00]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([]),
    },
    {
      buf: new Uint8Array([0x03, 0x00, 0x00, 0x00, ch('a'), ch('b'), ch('c'), 0x00, 0x44, 0x33, 0x22, 0x11]),
      scheme: [FieldTypes.TypeString, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([ch('a'), ch('b'), ch('c')]),
    },
    {
      buf: new Uint8Array([0x03, 0x00, 0x00, 0x00, ch('a'), ch('b'), ch('c'), 0x00, 0x44, 0x33, 0x22, 0x11]),
      scheme: [FieldTypes.TypeString, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      unionNum: 0,
      expected: new Uint8Array([0x44, 0x33, 0x22, 0x11]),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'), ch('b'), ch('c'), 0x00, 0x06,0x00,0x00,0x00, 0x55,0x66, 0x77,0x88, 0x99,0xaa, 0x00,0x00, 0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeString,FieldTypes.TypeUint16Array,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      unionNum: 0,
      expected: new Uint8Array([0x55,0x66,0x77,0x88,0x99,0xaa]),
    },
    {
      buf: new Uint8Array([0x03, 0x00, 0x00, 0x00, ch('a'), ch('b'), ch('c'), 0x00, 0x06, 0x00, 0x00, 0x00, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11]),
      scheme: [FieldTypes.TypeString, FieldTypes.TypeUint16Array, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 2,
      unionNum: 0,
      expected: new Uint8Array([0x44, 0x33, 0x22, 0x11]),
    },
    {
      buf: new Uint8Array([0x11, 0x12, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05]),
      scheme: [FieldTypes.TypeUint16, FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 1,
      unionNum: 0,
      expected: new Uint8Array([0x01, 0x02, 0x03, 0x04, 0x05]),
    },
    {
      buf: new Uint8Array([0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05]),
      scheme: [FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x01, 0x02, 0x03, 0x04, 0x05]),
    },
    {
      buf: new Uint8Array([0x01,0x00,0x33]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x33]),
    },
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00, 0x33,0x44,0x55,0x66]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x033,0x44,0x55,0x66]),
    },
    {
      buf: new Uint8Array([0x02,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x11,0x22,0x33]),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, ch('a'), ch('b'), ch('c')]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([ch('a'), ch('b'), ch('c')]),
    },
    {
      buf: new Uint8Array([0x04,0x00,0x00,0x00, 0x06,0x00,0x00,0x00, 0x11,0x11,0x22,0x22, 0x33,0x33]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString,FieldTypes.TypeUint16Array]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x11,0x11,0x22,0x22,0x33,0x33]),
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.rawBufferForField(tt.fieldNum, tt.unionNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageRawBufferWithHeaderForField', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x00,0x00,0x00,0x00]),
    },
    {
      buf: new Uint8Array([0x03, 0x00, 0x00, 0x00, ch('a'), ch('b'), ch('c'), 0x00, 0x44, 0x33, 0x22, 0x11]),
      scheme: [FieldTypes.TypeString, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x03, 0x00, 0x00, 0x00, ch('a'), ch('b'), ch('c')]),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c'), 0x00, 0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeString,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      unionNum: 0,
      expected: new Uint8Array([0x44,0x33,0x22,0x11]),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c'), 0x00, 0x06,0x00,0x00,0x00, 0x55,0x66, 0x77,0x88, 0x99,0xaa, 0x00,0x00, 0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeString,FieldTypes.TypeUint16Array,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      unionNum: 0,
      expected: new Uint8Array([0x06,0x00,0x00,0x00, 0x55,0x66,0x77,0x88,0x99,0xaa]),
    },
    {
      buf: new Uint8Array([0x03, 0x00, 0x00, 0x00, ch('a'), ch('b'), ch('c'), 0x00, 0x06, 0x00, 0x00, 0x00, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11]),
      scheme: [FieldTypes.TypeString, FieldTypes.TypeUint16Array, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 2,
      unionNum: 0,
      expected: new Uint8Array([0x44, 0x33, 0x22, 0x11]),
    },
    {
      buf: new Uint8Array([0x11,0x12, 0x00,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05]),
      scheme: [FieldTypes.TypeUint16,FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 1,
      unionNum: 0,
      expected: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05]),
    },
    {
      buf: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05]),
      scheme: [FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05]),
    },
    {
      buf: new Uint8Array([0x01,0x00,0x33]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x01,0x00,0x033]),
    },
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00, 0x33,0x44,0x55,0x66]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x00,0x00,0x00,0x00, 0x033,0x44,0x55,0x66]),
    },
    {
      buf: new Uint8Array([0x02,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x02,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33]),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c')]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x03,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c')]),
    },
    {
      buf: new Uint8Array([0x04,0x00,0x00,0x00, 0x06,0x00,0x00,0x00, 0x11,0x11,0x22,0x22, 0x33,0x33]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage,FieldTypes.TypeString,FieldTypes.TypeUint16Array]],
      fieldNum: 0,
      unionNum: 0,
      expected: new Uint8Array([0x04,0x00,0x00,0x00, 0x06,0x00,0x00,0x00, 0x11,0x11,0x22,0x22,0x33,0x33]),
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.rawBufferWithHeaderForField(tt.fieldNum, tt.unionNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageReadUint32', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0x55667788,
    },
    {
      buf: new Uint8Array([0x01, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8, FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x01,0x01,0x01, 0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8, FieldTypes.TypeUint8, FieldTypes.TypeUint8, FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 3,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeMessage, FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22]),
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      expected: 0,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getUint32(tt.fieldNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBe(tt.expected);
  }
});

test('TestMessageReadUint8', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11,0x88,0x99]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeUint8],
      unions: [],
      fieldNum: 2,
      expected: 0x99,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getUint8(tt.fieldNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBe(tt.expected);
  }
});

test('TestMessageReadUint16', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11,0x88,0x00,0x99,0xaa]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeUint16],
      unions: [],
      fieldNum: 2,
      expected: 0xaa99,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getUint16(tt.fieldNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBe(tt.expected);
  }
});

test('TestMessageReadUint64', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55,0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUint64],
      unions: [],
      fieldNum: 1,
      expected: BigInt('0x1122334455667788'),
    },
    {
      buf: new Uint8Array([0x44,0x00,0x00,0x00, 0x88,0x77,0x66,0x55,0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeUint64],
      unions: [],
      fieldNum: 1,
      expected: BigInt('0x1122334455667788'),
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getUint64(tt.fieldNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBe(tt.expected);
  }
});

test('TestMessageReadBytes', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array(),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x22,0x23,0x24]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x22,0x23,0x24]),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x22,0x23,0x24, 0x00, 0x02,0x00,0x00,0x00, 0x77,0x88]),
      scheme: [FieldTypes.TypeBytes,FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x77,0x88]),
    },
    {
      buf: new Uint8Array([0x01, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeBytes,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x11,0x22,0x33]),
    },
    {
      buf: new Uint8Array([0x01,0x01,0x01, 0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeUint8,FieldTypes.TypeUint8,FieldTypes.TypeBytes,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 3,
      expected: new Uint8Array([0x11,0x22,0x33]),
    },
    {
      buf: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeMessage,FieldTypes.TypeBytes,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x11,0x22,0x33]),
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array(),
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array(),
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x11,0x22]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array(),
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getBytes(tt.fieldNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageReadString', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      expected: "",
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c')]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      expected: "abc",
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c'), 0x00, 0x02,0x00,0x00,0x00, ch('d'),ch('e')]),
      scheme: [FieldTypes.TypeString,FieldTypes.TypeString],
      unions: [],
      fieldNum: 1,
      expected: "de",
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      expected: "",
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      expected: "",
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b')]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      expected: "",
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getString(tt.fieldNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBe(tt.expected);
  }
});

test('TestMessageReadMessage', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 0,
      expectedBuf: new Uint8Array(),
      expectedSize: 0,
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x01,0x02,0x03]),
      scheme: [FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 0,
      expectedBuf: new Uint8Array([0x01,0x02,0x03]),
      expectedSize: 3,
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x01,0x02,0x03, 0x00, 0x02,0x00,0x00,0x00, 0x04,0x05]),
      scheme: [FieldTypes.TypeMessage,FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 1,
      expectedBuf: new Uint8Array([0x04,0x05]),
      expectedSize: 2,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 0,
      expectedBuf: new Uint8Array(),
      expectedSize: 0,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22]),
      scheme: [FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 0,
      expectedBuf: new Uint8Array(),
      expectedSize: 0,
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x01,0x02]),
      scheme: [FieldTypes.TypeMessage],
      unions: [],
      fieldNum: 0,
      expectedBuf: new Uint8Array(),
      expectedSize: 0,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getMessage(tt.fieldNum);
    // console.log(tt); // uncomment on failure to find out where
    expect(s).toBeEqualToUint8Array(tt.expectedBuf);
    expect(s.byteLength).toBe(tt.expectedSize);
  }
});

test('TestMessageReadUnion', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00, 0x01,0x02,0x03,0x04]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8]],
      fieldNum: 0,
      unionNum: 0,
      unionIndex: 0,
      expectedIs: true,
      expectedOff: 4,
    },
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00, 0x01,0x02,0x03,0x04]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8]],
      fieldNum: 0,
      unionNum: 0,
      unionIndex: 1,
      expectedIs: false,
      expectedOff: 4,
    },
    {
      buf: new Uint8Array([0x01,0x00,0x00,0x00, 0x01,0x02,0x03,0x04]),
      scheme: [FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8]],
      fieldNum: 0,
      unionNum: 0,
      unionIndex: 1,
      expectedIs: true,
      expectedOff: 2,
    },
    {
      buf: new Uint8Array([0x01,0x00,0x11,0x00, 0x00,0x00,0x22,0x23]),
      scheme: [FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 1,
      unionNum: 1,
      unionIndex: 0,
      expectedIs: true,
      expectedOff: 6,
    },
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00, 0x11,0x22,0x33,0x44, 0x00,0x00,0x22,0x23]),
      scheme: [FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 1,
      unionNum: 1,
      unionIndex: 0,
      expectedIs: true,
      expectedOff: 10,
    },
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00, 0x11,0x22,0x33,0x44, 0x02,0x00,0x00,0x00, 0x01,0x02,0x03,0x04]),
      scheme: [FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 1,
      unionNum: 1,
      unionIndex: 2,
      expectedIs: true,
      expectedOff: 12,
    },
    {
      buf: new Uint8Array([0x22,0x22,0x22,0x22, 0x02,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x01,0x00, 0x17]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 2,
      unionNum: 1,
      unionIndex: 1,
      expectedIs: true,
      expectedOff: 20,
    },
    {
      buf: new Uint8Array([0x22,0x22,0x22,0x22, 0x07,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x01,0x00, 0x17]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 2,
      unionNum: 1,
      unionIndex: 1,
      expectedIs: false,
      expectedOff: 0,
    },
    {
      buf: new Uint8Array([0x22,0x22,0x22,0x22, 0x02,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x01]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 2,
      unionNum: 1,
      unionIndex: 1,
      expectedIs: false,
      expectedOff: 0,
    },
    {
      buf: new Uint8Array([0x22,0x22,0x22,0x22, 0x02,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x01,0x00]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 2,
      unionNum: 1,
      unionIndex: 1,
      expectedIs: false,
      expectedOff: 0,
    },
    {
      buf: new Uint8Array([0x22,0x22,0x22,0x22, 0x02,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x00,0x00, 0x11]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUnion,FieldTypes.TypeUnion],
      unions: [[FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeMessage],[FieldTypes.TypeUint16,FieldTypes.TypeUint8,FieldTypes.TypeUint32]],
      fieldNum: 2,
      unionNum: 1,
      unionIndex: 1,
      expectedIs: false,
      expectedOff: 0,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const [is, off] = m.isUnionIndex(tt.fieldNum, tt.unionNum, tt.unionIndex);
    // console.log(tt); // uncomment on failure to find out where
    expect(is).toBe(tt.expectedIs);
    expect(off).toBe(tt.expectedOff);
  }
});

test('TestMessageMutateUint32', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x55,0x55,0x55,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x44,0x33,0x22,0x11, 0x55,0x55,0x55,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x01, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeUint32,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x01, 0x00,0x00,0x00, 0x55,0x55,0x55,0x55, 0x88,0x77,0x66,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x01,0x01,0x01, 0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeUint8,FieldTypes.TypeUint8,FieldTypes.TypeUint32,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 3,
      expected: new Uint8Array([0x01,0x01,0x01, 0x00, 0x55,0x55,0x55,0x55, 0x88,0x77,0x66,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeMessage,FieldTypes.TypeUint32,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x55,0x55,0x55,0x55, 0x88,0x77,0x66,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x44,0x33,0x22,0x11]),
      err: true,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22]),
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x44,0x33,0x22]),
      err: true,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    // console.log(tt); // uncomment on failure to find out where
    if (tt.err) {
      expect(() => {
        m.setUint32(tt.fieldNum, 0x55555555);
      }).toThrow();
    } else {
      expect(() => {
        m.setUint32(tt.fieldNum, 0x55555555);
      }).not.toThrow();
    }
    expect(tt.buf).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageMutateUint8', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11,0x88,0x99]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeUint8],
      unions: [],
      fieldNum: 2,
      expected: new Uint8Array([0x44,0x33,0x22,0x11,0x88,0x55]),
      err: false,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    // console.log(tt); // uncomment on failure to find out where
    if (tt.err) {
      expect(() => {
        m.setUint8(tt.fieldNum, 0x55);
      }).toThrow();
    } else {
      expect(() => {
        m.setUint8(tt.fieldNum, 0x55);
      }).not.toThrow();
    }
    expect(tt.buf).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageMutateUint16', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11,0x88,0x00,0x99,0xaa]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUint8,FieldTypes.TypeUint16],
      unions: [],
      fieldNum: 2,
      expected: new Uint8Array([0x44,0x33,0x22,0x11,0x88,0x00,0x55,0x55]),
      err: false,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    // console.log(tt); // uncomment on failure to find out where
    if (tt.err) {
      expect(() => {
        m.setUint16(tt.fieldNum, 0x5555);
      }).toThrow();
    } else {
      expect(() => {
        m.setUint16(tt.fieldNum, 0x5555);
      }).not.toThrow();
    }
    expect(tt.buf).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageMutateUint64', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55,0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint32,FieldTypes.TypeUint64],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x44,0x33,0x22,0x11, 0x55,0x55,0x55,0x55,0x55,0x55,0x55,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x44,0x00,0x00,0x00, 0x88,0x77,0x66,0x55,0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeUint64],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x44,0x00,0x00,0x00, 0x55,0x55,0x55,0x55,0x55,0x55,0x55,0x55]),
      err: false,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    // console.log(tt); // uncomment on failure to find out where
    if (tt.err) {
      expect(() => {
        m.setUint64(tt.fieldNum, BigInt('0x5555555555555555'));
      }).toThrow();
    } else {
      expect(() => {
        m.setUint64(tt.fieldNum, BigInt('0x5555555555555555'));
      }).not.toThrow();
    }
    expect(tt.buf).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageMutateBytes', () => {
  const tests = [
    {
      buf: new Uint8Array([0x00,0x00,0x00,0x00]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x00,0x00,0x00,0x00]),
      err: true,
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x22,0x23,0x24]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x03,0x00,0x00,0x00, 0x55,0x55,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x22,0x23,0x24, 0x00, 0x02,0x00,0x00,0x00, 0x77,0x88]),
      scheme: [FieldTypes.TypeBytes,FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x03,0x00,0x00,0x00, 0x22,0x23,0x24, 0x00, 0x02,0x00,0x00,0x00, 0x77,0x88]),
      err: true,
    },
    {
      buf: new Uint8Array([0x01, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeBytes,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x01, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x55,0x55,0x55, 0x00, 0x88,0x77,0x66,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x01,0x01,0x01, 0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeUint8,FieldTypes.TypeUint8,FieldTypes.TypeUint8,FieldTypes.TypeBytes,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 3,
      expected: new Uint8Array([0x01,0x01,0x01, 0x00, 0x03,0x00,0x00,0x00, 0x55,0x55,0x55, 0x00, 0x88,0x77,0x66,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55]),
      scheme: [FieldTypes.TypeMessage,FieldTypes.TypeBytes,FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x55,0x55,0x55, 0x00, 0x88,0x77,0x66,0x55]),
      err: false,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x44,0x33,0x22,0x11]),
      err: true,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x44,0x33,0x22]),
      err: true,
    },
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, 0x11,0x22]),
      scheme: [FieldTypes.TypeBytes],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x03,0x00,0x00,0x00, 0x11,0x22]),
      err: true,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    // console.log(tt); // uncomment on failure to find out where
    if (tt.err) {
      expect(() => {
        m.setBytes(tt.fieldNum, new Uint8Array([0x55,0x55,0x55]));
      }).toThrow();
    } else {
      expect(() => {
        m.setBytes(tt.fieldNum, new Uint8Array([0x55,0x55,0x55]));
      }).not.toThrow();
    }
    expect(tt.buf).toBeEqualToUint8Array(tt.expected);
  }
});

test('TestMessageMutateString', () => {
  const tests = [
    {
      buf: new Uint8Array([0x03,0x00,0x00,0x00, ch('a'),ch('b'),ch('c')]),
      scheme: [FieldTypes.TypeString],
      unions: [],
      fieldNum: 0,
      expected: new Uint8Array([0x03,0x00,0x00,0x00, ch('z'),ch('z'),ch('z')]),
      err: false,
    },
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    // console.log(tt); // uncomment on failure to find out where
    if (tt.err) {
      expect(() => {
        m.setString(tt.fieldNum, "zzz");
      }).toThrow();
    } else {
      expect(() => {
        m.setString(tt.fieldNum, "zzz");
      }).not.toThrow();
    }
    expect(tt.buf).toBeEqualToUint8Array(tt.expected);
  }
});