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

const FieldDynamic = Object.freeze({
  [FieldTypes.TypeMessage]: true,
  [FieldTypes.TypeBytes]: true,
  [FieldTypes.TypeString]: true,
  [FieldTypes.TypeUnion]: true,
  [FieldTypes.TypeUint8]: false,
  [FieldTypes.TypeUint16]: false,
  [FieldTypes.TypeUint32]: false,
  [FieldTypes.TypeUint64]: false,
  [FieldTypes.TypeUint8Array]: true,
  [FieldTypes.TypeUint16Array]: true,
  [FieldTypes.TypeUint32Array]: true,
  [FieldTypes.TypeUint64Array]: true,
  [FieldTypes.TypeMessageArray]: true,
  [FieldTypes.TypeBytesArray]: true,
  [FieldTypes.TypeStringArray]: true,
});

const FieldDynamicContentAlignment = Object.freeze({
  [FieldTypes.TypeMessage]: 4,
  [FieldTypes.TypeBytes]: 1,
  [FieldTypes.TypeString]: 1,
  [FieldTypes.TypeUnion]: 0,
  [FieldTypes.TypeUint8]: 0,
  [FieldTypes.TypeUint16]: 0,
  [FieldTypes.TypeUint32]: 0,
  [FieldTypes.TypeUint64]: 0,
  [FieldTypes.TypeUint8Array]: 1,
  [FieldTypes.TypeUint16Array]: 2,
  [FieldTypes.TypeUint32Array]: 4,
  [FieldTypes.TypeUint64Array]: 4,
  [FieldTypes.TypeMessageArray]: 4,
  [FieldTypes.TypeBytesArray]: 4,
  [FieldTypes.TypeStringArray]: 4,
});

export class InternalMessage {

  constructor(buf, size, scheme, unions) {
    this.bytes = buf; // buf should be Uint8Array (a view over an ArrayBuffer)
    this.size = size;
    this.scheme = scheme;
    this.unions = unions;
    this.dataView = new DataView(buf.buffer);
    this.offsets = null; // map: fieldNum -> offset in bytes
  }

  alignOffsetToType(off, fieldType) {
    const fieldSize = FieldAlignment[fieldType];
    return Math.floor((off + fieldSize - 1) / fieldSize) * fieldSize;
  }

  alignDynamicFieldContentOffset(off, fieldType) {
    const contentAlignment = FieldDynamicContentAlignment[fieldType];
    return Math.floor((off + contentAlignment - 1) / contentAlignment) * contentAlignment;
  }

  lazyCalcOffsets() {
    if (this.offsets !== null) {
      return true;
    }
    const res = {};
    let off = 0;
    let unionNum = 0;
    for (let fieldNum = 0; fieldNum < this.scheme.length; fieldNum++) {
      let fieldType = this.scheme[fieldNum];

      // write the current offset
      off = this.alignOffsetToType(off, fieldType);
      if (off >= this.size) {
        return false;
      }
      res[fieldNum] = off;

      // skip over the content to the next field
      if (fieldType == FieldTypes.TypeUnion) {
        if (off + FieldSizes[FieldTypes.TypeUnion] > this.size) {
          return false;
        }
        const unionType = this.dataView.getUint16(off, true);
        off += FieldSizes[FieldTypes.TypeUnion];
        if (unionNum >= this.unions.length || unionType >= this.unions[unionNum].length) {
          return false;
        }
        fieldType = this.unions[unionNum][unionType];
        unionNum += 1;
        off = this.alignOffsetToType(off, fieldType);
      }
      if (FieldDynamic[fieldType]) {
        if (off + FieldSizes[fieldType] > this.size) {
          return false;
        }
        const contentSize = this.dataView.getUint32(off, true);
        off += FieldSizes[fieldType];
        off = this.alignDynamicFieldContentOffset(off, fieldType);
        off += contentSize;
      } else {
        off += FieldSizes[fieldType];
      }
    }
    if (off > this.size) {
      return false;
    }
    this.offsets = res;
    return true;
  }

  isValid() {
    if (this.bytes === undefined) {
      throw `uninitialized membuffer, did you create it directly without a Builder or a Reader?`;
    }
    return this.lazyCalcOffsets();
  }

  rawBuffer() {
    return this.bytes.subarray(0, this.size);
  }

  rawBufferForField(fieldNum, unionNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length || fieldNum >= this.scheme.length) {
      return new Uint8Array();
    }
    let fieldType = this.scheme[fieldNum];
    let off = this.offsets[fieldNum];
    if (fieldType == FieldTypes.TypeUnion) {
      const unionType = this.dataView.getUint16(off, true);
      off += FieldSizes[FieldTypes.TypeUnion];
      if (unionNum >= this.unions.length || unionType >= this.unions[unionNum].length) {
        return new Uint8Array();
      }
      fieldType = this.unions[unionNum][unionType];
      off = this.alignOffsetToType(off, fieldType);
    }
    if (FieldDynamic[fieldType]) {
      const contentSize = this.dataView.getUint32(off, true);
      off += FieldSizes[fieldType];
      off = this.alignDynamicFieldContentOffset(off, fieldType);
      return this.bytes.subarray(off, off+contentSize);
    } else {
      return this.bytes.subarray(off, off+FieldSizes[fieldType]);
    }
  }

  rawBufferWithHeaderForField(fieldNum, unionNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length || fieldNum >= this.scheme.length) {
      return new Uint8Array();
    }
    let fieldType = this.scheme[fieldNum];
    let off = this.offsets[fieldNum];
    const fieldHeaderOff = off;
    if (fieldType == FieldTypes.TypeUnion) {
      const unionType = this.dataView.getUint16(off, true);
      off += FieldSizes[FieldTypes.TypeUnion];
      if (unionNum >= this.unions.length || unionType >= this.unions[unionNum].length) {
        return new Uint8Array();
      }
      fieldType = this.unions[unionNum][unionType];
      off = this.alignOffsetToType(off, fieldType);
    }
    if (FieldDynamic[fieldType]) {
      const contentSize = this.dataView.getUint32(off, true);
      off += FieldSizes[fieldType];
      off = this.alignDynamicFieldContentOffset(off, fieldType);
      return this.bytes.subarray(fieldHeaderOff, off+contentSize);
    } else {
      return this.bytes.subarray(fieldHeaderOff, off+FieldSizes[fieldType]);
    }
  }

  getUint8InOffset(off) {
    return this.dataView.getUint8(off, true);
  }

  getUint8(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint8InOffset(off);
  }

  getUint16InOffset(off) {
    return this.dataView.getUint16(off, true);
  }

  getUint16(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint16InOffset(off);
  }

  getUint32InOffset(off) {
    return this.dataView.getUint32(off, true);
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