import {} from './matcher-extensions';
import {FieldTypes, InternalMessage} from './membuffers';

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

function ch(char) {
  return char.charCodeAt(0);
}
