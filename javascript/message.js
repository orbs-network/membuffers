import {FieldTypes, FieldSizes, FieldAlignment, FieldDynamic, FieldDynamicContentAlignment} from './types';
import {Iterator} from "./iterator";
import {getTextEncoder, getTextDecoder} from './text';

export function alignOffsetToType(off, fieldType) {
  const fieldSize = FieldAlignment[fieldType];
  return Math.floor((off + fieldSize - 1) / fieldSize) * fieldSize;
}

export function alignDynamicFieldContentOffset(off, fieldType) {
  const contentAlignment = FieldDynamicContentAlignment[fieldType];
  return Math.floor((off + contentAlignment - 1) / contentAlignment) * contentAlignment;
}

export class InternalMessage {

  constructor(buf, size, scheme, unions) {
    this.bytes = buf; // buf should be Uint8Array (a view over an ArrayBuffer)
    this.size = size;
    this.scheme = scheme;
    this.unions = unions;
    this.dataView = new DataView(buf.buffer, buf.byteOffset);
    this.offsets = null; // map: fieldNum -> offset in bytes
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
      off = alignOffsetToType(off, fieldType);
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
        off = alignOffsetToType(off, fieldType);
      }
      if (FieldDynamic[fieldType]) {
        if (off + FieldSizes[fieldType] > this.size) {
          return false;
        }
        const contentSize = this.dataView.getUint32(off, true);
        off += FieldSizes[fieldType];
        off = alignDynamicFieldContentOffset(off, fieldType);
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
      off = alignOffsetToType(off, fieldType);
    }
    if (FieldDynamic[fieldType]) {
      const contentSize = this.dataView.getUint32(off, true);
      off += FieldSizes[fieldType];
      off = alignDynamicFieldContentOffset(off, fieldType);
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
      off = alignOffsetToType(off, fieldType);
    }
    if (FieldDynamic[fieldType]) {
      const contentSize = this.dataView.getUint32(off, true);
      off += FieldSizes[fieldType];
      off = alignDynamicFieldContentOffset(off, fieldType);
      return this.bytes.subarray(fieldHeaderOff, off+contentSize);
    } else {
      return this.bytes.subarray(fieldHeaderOff, off+FieldSizes[fieldType]);
    }
  }

  getOffsetInOffset(off) {
    return this.dataView.getUint32(off, true);
  }

  getUint8InOffset(off) {
    return this.dataView.getUint8(off, true);
  }

  setUint8InOffset(off, v) {
    return this.dataView.setUint8(off, v, true);
  }

  getUint8(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint8InOffset(off);
  }

  setUint8(fieldNum, v) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint8InOffset(off, v);
  }

  getUint16InOffset(off) {
    return this.dataView.getUint16(off, true);
  }

  setUint16InOffset(off, v) {
    return this.dataView.setUint16(off, v, true);
  }

  getUint16(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint16InOffset(off);
  }

  setUint16(fieldNum, v) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint16InOffset(off, v);
  }

  getUint32InOffset(off) {
    return this.dataView.getUint32(off, true);
  }

  setUint32InOffset(off, v) {
    return this.dataView.setUint32(off, v, true);
  }

  getUint32(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint32InOffset(off);
  }

  setUint32(fieldNum, v) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint32InOffset(off, v);
  }

  getUint64InOffset(off) {
    return this.dataView.getBigUint64(off, true);
  }

  setUint64InOffset(off, v) {
    return this.dataView.setBigUint64(off, v, true);
  }

  getUint64(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return BigInt(0);
    }
    const off = this.offsets[fieldNum];
    return this.getUint64InOffset(off);
  }

  setUint64(fieldNum, v) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint64InOffset(off, v);
  }

  getMessageInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeMessage];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeMessage);
    return this.bytes.subarray(off, off+contentSize);
  }

  getMessage(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Uint8Array();
    }
    const off = this.offsets[fieldNum];
    return this.getMessageInOffset(off);
  }

  getBytesInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeBytes];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeBytes);
    if (off+contentSize > this.bytes.byteLength) {
      return new Uint8Array();
    }
    return this.bytes.subarray(off, off+contentSize);
  }

  setBytesInOffset(off, v) {
    const contentSize = this.dataView.getUint32(off, true);
    if (contentSize != v.byteLength) {
      throw new Error("size mismatch");
    }
    off += FieldSizes[FieldTypes.TypeBytes];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeBytes);
    return this.bytes.set(v, off);
  }

  getBytes(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Uint8Array();
    }
    const off = this.offsets[fieldNum];
    return this.getBytesInOffset(off);
  }

  setBytes(fieldNum, v) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setBytesInOffset(off, v);
  }

  getStringInOffset(off) {
    const b = this.getBytesInOffset(off);
    return getTextDecoder().decode(b);
  }

  setStringInOffset(off, v) {
    return this.setBytesInOffset(off, getTextEncoder().encode(v));
  }

  getString(fieldNum) {
    const b = this.getBytes(fieldNum);
    return getTextDecoder().decode(b);
  }

  setString(fieldNum, v) {
    return this.setBytes(fieldNum, getTextEncoder().encode(v));
  }

  getUnionIndex(fieldNum, unionNum) {
    const invalidUnionIndex = 0xffff;
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return invalidUnionIndex;
    }
    let off = this.offsets[fieldNum];
    const unionType = this.dataView.getUint16(off, true);
    off += FieldSizes[FieldTypes.TypeUnion];
    if (unionNum >= this.unions.length || unionType >= this.unions[unionNum].length) {
      return invalidUnionIndex;
    }
    const fieldType = this.unions[unionNum][unionType];
    off = alignOffsetToType(off, fieldType);
    return unionType;
  }

  isUnionIndex(fieldNum, unionNum, unionIndex) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return [false, 0];
    }
    let off = this.offsets[fieldNum];
    const unionType = this.dataView.getUint16(off, true);
    off += FieldSizes[FieldTypes.TypeUnion];
    if (unionNum >= this.unions.length || unionType >= this.unions[unionNum].length) {
      return [false, 0];
    }
    const fieldType = this.unions[unionNum][unionType];
    off = alignOffsetToType(off, fieldType);
    return [unionType == unionIndex, off];
  }

  getUint8ArrayIteratorInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint8Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint8Array);
    return new Iterator(off, off+contentSize, FieldTypes.TypeUint8, this);
  }

  getUint8ArrayIterator(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Iterator(0, 0, FieldTypes.TypeUint8, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint8ArrayIteratorInOffset(off);
  }

  getUint16ArrayIteratorInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint32Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint16Array);
    return new Iterator(off, off+contentSize, FieldTypes.TypeUint16, this);
  }

  getUint16ArrayIterator(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Iterator(0, 0, FieldTypes.TypeUint16, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint16ArrayIteratorInOffset(off);
  }

  getUint32ArrayIteratorInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint32Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint32Array);
    return new Iterator(off, off+contentSize, FieldTypes.TypeUint32, this);
  }

  getUint32ArrayIterator(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Iterator(0, 0, FieldTypes.TypeUint32, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint32ArrayIteratorInOffset(off);
  }

  getUint64ArrayIteratorInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint64Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint64Array);
    return new Iterator(off, off+contentSize, FieldTypes.TypeUint64, this);
  }

  getUint64ArrayIterator(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Iterator(0, 0, FieldTypes.TypeUint64, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint64ArrayIteratorInOffset(off);
  }

  getMessageArrayIteratorInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeMessageArray];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeMessageArray);
    return new Iterator(off, off+contentSize, FieldTypes.TypeMessage, this);
  }

  getMessageArrayIterator(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Iterator(0, 0, FieldTypes.TypeMessage, this);
    }
    const off = this.offsets[fieldNum];
    return this.getMessageArrayIteratorInOffset(off);
  }

  getBytesArrayIteratorInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeMessageArray];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeBytesArray);
    return new Iterator(off, off+contentSize, FieldTypes.TypeBytes, this);
  }

  getBytesArrayIterator(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Iterator(0, 0, FieldTypes.TypeBytes, this);
    }
    const off = this.offsets[fieldNum];
    return this.getBytesArrayIteratorInOffset(off);
  }

  getStringArrayIteratorInOffset(off) {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeStringArray];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeStringArray);
    return new Iterator(off, off+contentSize, FieldTypes.TypeString, this);
  }

  getStringArrayIterator(fieldNum) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Iterator(0, 0, FieldTypes.TypeString, this);
    }
    const off = this.offsets[fieldNum];
    return this.getStringArrayIteratorInOffset(off);
  }

}