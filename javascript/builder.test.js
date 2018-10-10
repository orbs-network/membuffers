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