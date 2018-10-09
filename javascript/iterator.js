import {FieldTypes, FieldSizes, FieldAlignment, FieldDynamic, FieldDynamicContentAlignment} from './types';

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
    this.cursor = this.m.alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint8Array);
    return res;
  }

  nextUint16() {
    if (this.cursor+FieldSizes[FieldTypes.TypeUint16] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint16InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint16];
    this.cursor = this.m.alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint16Array);
    return res;
  }

  nextUint32() {
    if (this.cursor+FieldSizes[FieldTypes.TypeUint32] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint32InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint32];
    this.cursor = this.m.alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint32Array);
    return res;
  }

  nextUint64() {
    if (this.cursor+FieldSizes[FieldTypes.TypeUint64] > this.endCursor) {
      this.cursor = this.endCursor;
      return 0;
    }
    const res = this.m.getUint64InOffset(this.cursor);
    this.cursor += FieldSizes[FieldTypes.TypeUint64];
    this.cursor = this.m.alignDynamicFieldContentOffset(this.cursor, FieldTypes.TypeUint64Array);
    return res;
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