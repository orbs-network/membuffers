export const FieldTypes = Object.freeze({
  TypeMessage: 1,
  TypeBytes: 2,
  TypeString: 3,
  TypeUnion: 4,
  TypeUint8: 11,
  TypeUint16: 12,
  TypeUint32: 13,
  TypeUint64: 14,
  TypeUint8Array: 21,
  TypeUint16Array: 22,
  TypeUint32Array: 23,
  TypeUint64Array: 24,
  TypeMessageArray: 31,
  TypeBytesArray: 32,
  TypeStringArray: 33,
});

const FieldSizes = Object.freeze({
  [FieldTypes.TypeMessage]: 4,
  [FieldTypes.TypeBytes]: 4,
  [FieldTypes.TypeString]: 4,
  [FieldTypes.TypeUnion]: 2,
  [FieldTypes.TypeUint8]: 1,
  [FieldTypes.TypeUint16]: 2,
  [FieldTypes.TypeUint32]: 4,
  [FieldTypes.TypeUint64]: 8,
  [FieldTypes.TypeUint8Array]: 4,
  [FieldTypes.TypeUint16Array]: 4,
  [FieldTypes.TypeUint32Array]: 4,
  [FieldTypes.TypeUint64Array]: 4,
  [FieldTypes.TypeMessageArray]: 4,
  [FieldTypes.TypeBytesArray]: 4,
  [FieldTypes.TypeStringArray]: 4,
});

const FieldAlignment = Object.freeze({
  [FieldTypes.TypeMessage]: 4,
  [FieldTypes.TypeBytes]: 4,
  [FieldTypes.TypeString]: 4,
  [FieldTypes.TypeUnion]: 2,
  [FieldTypes.TypeUint8]: 1,
  [FieldTypes.TypeUint16]: 2,
  [FieldTypes.TypeUint32]: 4,
  [FieldTypes.TypeUint64]: 4,
  [FieldTypes.TypeUint8Array]: 4,
  [FieldTypes.TypeUint16Array]: 4,
  [FieldTypes.TypeUint32Array]: 4,
  [FieldTypes.TypeUint64Array]: 4,
  [FieldTypes.TypeMessageArray]: 4,
  [FieldTypes.TypeBytesArray]: 4,
  [FieldTypes.TypeStringArray]: 4,
});

export class InternalMessage {

  constructor(buf, size, scheme, unions) {
    this.bytes = buf;
    this.size = size;
    this.scheme = scheme;
    this.unions = unions;
    this.view = new DataView(buf);
    this.offsets = null; // map: fieldNum -> offset in bytes
  }

  alignOffsetToType(off, fieldType) {
    const fieldSize = FieldAlignment[fieldType];
    return Math.floor((off + fieldSize - 1) / fieldSize) * fieldSize;
  }

  lazyCalcOffsets() {
    if (this.offsets !== null) {
      return true;
    }
    const res = {};
    let off = 0;
    let unionNum = 0;
    for (let fieldNum = 0; fieldNum < this.scheme.length; fieldNum++) {
      const fieldType = this.scheme[fieldNum];
      off = this.alignOffsetToType(off, fieldType);
      if (off >= this.size) {
        return false;
      }
      res[fieldNum] = off;

      // skip over the content to the next field
      off += FieldSizes[fieldType];
    }
    if (off > this.size) {
      return false;
    }
    this.offsets = res;
    return true;
  }

  getUint32InOffset(off) {
    return this.view.getUint32(off, true);
  }

  getUint32(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint32InOffset(off);
  }

}

export class InternalBuilder {

}