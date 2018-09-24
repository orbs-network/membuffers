import {FieldTypes, InternalMessage} from './membuffers';

test('TestMessageReadUint32', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]).buffer,
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]).buffer,
      scheme: [FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0x55667788,
    },
    {
      buf: new Uint8Array([0x01, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]).buffer,
      scheme: [FieldTypes.TypeUint8, FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x01,0x01,0x01, 0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]).buffer,
      scheme: [FieldTypes.TypeUint8, FieldTypes.TypeUint8, FieldTypes.TypeUint8, FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 3,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55]).buffer,
      scheme: [FieldTypes.TypeMessage, FieldTypes.TypeUint32, FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0x11223344,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]).buffer,
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 1,
      expected: 0,
    },
    {
      buf: new Uint8Array([0x44,0x33,0x22]).buffer,
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