/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import "./matcher-extensions";
import { FieldTypes, FieldType } from "./types";
import { InternalMessage } from "./message";
import { InternalBuilder } from "./builder";

test("TestCreateAndReadComplexMessage", () => {
    const scheme = [FieldTypes.TypeUint32, FieldTypes.TypeBytesArray, FieldTypes.TypeBytes32Array, FieldTypes.TypeUint32];
    const firstUint = 24041977;
    const secondUint = 64;
    const bytesArray = [new Uint8Array([0x33, 0x11, 0x22]), new Uint8Array([0x04, 0x01, 0x02]), new Uint8Array([0xbb, 0xdd, 0xcc])];
    const bytes32Array = [
        new Uint8Array([0xee, 0xdd, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
            0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xcc, 0xdd]),
        new Uint8Array([0xaa, 0xbb, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
            0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xbb, 0xaa])];
    const w = new InternalBuilder();
    const buf = new Uint8Array(104);
    w.writeUint32(buf, firstUint);
    w.writeBytesArray(buf, bytesArray);
    w.writeBytes32Array(buf, bytes32Array);
    w.writeUint32(buf, secondUint);

    const m = new InternalMessage(buf, buf.byteLength, scheme, []);

    expect(m.getUint32(0)).toBe(firstUint);
    const iteratorBytes = m.getBytesArrayIterator(1);
    let i = 0;
    for (const v of iteratorBytes) {
        expect(v).toBeEqualToUint8Array(bytesArray[i]);
        i++
    }
    const iteratorBytes32 = m.getBytes32ArrayIterator(2);
    i = 0;
    for (const v of iteratorBytes32) {
        expect(v).toBeEqualToUint8Array(bytes32Array[i]);
        i++
    }
    expect(m.getUint32(3)).toBe(secondUint);
});
