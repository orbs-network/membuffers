import {FieldTypes, InternalMessage} from './membuffers';

test('TestMessageReadUint32', () => {
  const tests = [
    {
      buf: new Uint8Array([0x44,0x33,0x22,0x11]).buffer,
      scheme: [FieldTypes.TypeUint32],
      unions: [],
      fieldNum: 0,
      expected: 0x11223344,
    }
  ];

  for (const tt of tests) {
    const m = new InternalMessage(tt.buf, tt.buf.byteLength, tt.scheme, tt.unions);
    const s = m.getUint32(tt.fieldNum);
    expect(s).toBe(tt.expected);
  }
});