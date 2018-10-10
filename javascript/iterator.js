import {alignDynamicFieldContentOffset} from './message';
import {FieldTypes, FieldSizes} from './types';
import {getTextDecoder} from "./text";

export class Iterator {

  constructor(cursor, endCursor, fieldType, m) {
    this.cursor = cursor;
    this.endCursor = endCursor;
    this.fieldType = fieldType;
    this.m = m;
  }

  hasNext() {
    return this.cursor < this.endCursor;
  }

  nextUint8() {
    if (this.cursor+FieldSizes[FieldTypes.TypeUint8] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint8InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint8];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint8Array);
    return res;
  }

  nextUint16() {
    if (this.cursor+FieldSizes[FieldTypes.TypeUint16] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint16InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint16];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint16Array);
    return res;
  }

  nextUint32() {
    if (this.cursor+FieldSizes[FieldTypes.TypeUint32] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint32InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint32];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint32Array);
    return res;
  }

  nextUint64() {
    if (this.cursor+FieldSizes[FieldTypes.TypeUint64] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint64InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint64];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint64Array);
    return res;
  }

  nextMessage() {
    if (this.cursor+FieldSizes[FieldTypes.TypeMessage] > this.endCursor) {
      this.cursor = this.endCursor;
      return [new Uint8Array(), 0];
    }
    const resSize = this.m.getOffsetInOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeMessage];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeMessage);
    if (this.cursor+resSize > this.endCursor) {
      this.cursor = this.endCursor;
      return [new Uint8Array(), 0];
    }
    const resBuf = this.m.bytes.subarray(this.cursor, this.cursor+resSize);
    this.cursor += resSize;
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeMessageArray);
    return [resBuf, resSize];
  }

  nextBytes() {
    if (this.cursor+FieldSizes[FieldTypes.TypeBytes] > this.endCursor) {
      this.cursor = this.endCursor;
      return new Uint8Array();
    }
    const resSize = this.m.getOffsetInOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeBytes];
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeBytes);
    if (this.cursor+resSize > this.endCursor) {
      this.cursor = this.endCursor;
      return new Uint8Array();
    }
    const resBuf = this.m.bytes.subarray(this.cursor, this.cursor+resSize);
    this.cursor += resSize;
    this.cursor = alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeBytesArray);
    return resBuf;
  }

  nextString() {
    const b = this.nextBytes();
    return getTextDecoder().decode(b);
  }

  [Symbol.iterator]() {
    return {
      next: () => {
        if (this.hasNext()) {
          switch (this.fieldType) {
            case FieldTypes.TypeUint8:
              return {value: this.nextUint8(), done: false};
            case FieldTypes.TypeUint16:
              return {value: this.nextUint16(), done: false};
            case FieldTypes.TypeUint32:
              return {value: this.nextUint32(), done: false};
            case FieldTypes.TypeUint64:
              return {value: this.nextUint64(), done: false};
            case FieldTypes.TypeMessage:
              return {value: this.nextMessage(), done: false};
            case FieldTypes.TypeBytes:
              return {value: this.nextBytes(), done: false};
            case FieldTypes.TypeString:
              return {value: this.nextString(), done: false};
            default:
              throw new Error("unsupported array type");
          }
        } else {
          return {done: true};
        }
      }
    }
  }

}