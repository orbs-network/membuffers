/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import { alignDynamicFieldContentOffset, InternalMessage } from "./message";
import { FieldTypes, FieldSizes, FieldType } from "./types";
import { getTextDecoder } from "./text";

export class ArrayIterator {
  constructor(private cursor: number, private endCursor: number, private fieldType: FieldType, private m: InternalMessage) {}

  hasNext(): boolean {
    return this.cursor < this.endCursor;
  }

  nextUint8(): number {
    if (this.cursor + FieldSizes[FieldTypes.TypeUint8] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint8InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint8];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint8Array);
    return res;
  }

  nextUint16(): number {
    if (this.cursor + FieldSizes[FieldTypes.TypeUint16] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint16InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint16];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint16Array);
    return res;
  }

  nextUint32(): number {
    if (this.cursor + FieldSizes[FieldTypes.TypeUint32] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint32InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint32];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint32Array);
    return res;
  }

  nextUint64(): bigint {
    if (this.cursor + FieldSizes[FieldTypes.TypeUint64] > this.endCursor) {
      this.cursor = this.endCursor;
      return BigInt(0);
    }
    const res = this.m.getUint64InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint64];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint64Array);
    return res;
  }

  nextMessage(): [Uint8Array, number] {
    if (this.cursor + FieldSizes[FieldTypes.TypeMessage] > this.endCursor) {
      this.cursor = this.endCursor;
      return [new Uint8Array(), 0];
    }
    const resSize = this.m.getOffsetInOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeMessage];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeMessage);
    if (this.cursor + resSize > this.endCursor) {
      this.cursor = this.endCursor;
      return [new Uint8Array(), 0];
    }
    const resBuf = this.m.bytes.subarray(this.cursor, this.cursor + resSize);
    this.cursor += resSize;
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeMessageArray);
    return [resBuf, resSize];
  }

  nextBytes(): Uint8Array {
    if (this.cursor + FieldSizes[FieldTypes.TypeBytes] > this.endCursor) {
      this.cursor = this.endCursor;
      return new Uint8Array();
    }
    const resSize = this.m.getOffsetInOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeBytes];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeBytes);
    if (this.cursor + resSize > this.endCursor) {
      this.cursor = this.endCursor;
      return new Uint8Array();
    }
    const resBuf = this.m.bytes.subarray(this.cursor, this.cursor + resSize);
    this.cursor += resSize;
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeBytesArray);
    return resBuf;
  }

  nextString(): string {
    const b = this.nextBytes();
    return getTextDecoder().decode(b);
  }

  [Symbol.iterator](): Iterator<number | bigint | [Uint8Array, number] | Uint8Array | string> {
    return {
      next: () => {
        if (this.hasNext()) {
          switch (this.fieldType) {
            case FieldTypes.TypeUint8:
              return { value: this.nextUint8(), done: false };
            case FieldTypes.TypeUint16:
              return { value: this.nextUint16(), done: false };
            case FieldTypes.TypeUint32:
              return { value: this.nextUint32(), done: false };
            case FieldTypes.TypeUint64:
              return { value: this.nextUint64(), done: false };
            case FieldTypes.TypeMessage:
              return { value: this.nextMessage(), done: false };
            case FieldTypes.TypeBytes:
              return { value: this.nextBytes(), done: false };
            case FieldTypes.TypeString:
              return { value: this.nextString(), done: false };
            default:
              throw new Error("unsupported array type");
          }
        } else {
          return { done: true, value: undefined };
        }
      },
    };
  }
}
