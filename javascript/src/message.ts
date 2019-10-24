/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import { FieldTypes, FieldSizes, FieldAlignment, FieldDynamic, FieldDynamicContentAlignment, FieldType } from "./types";
import { ArrayIterator } from "./iterator";
import { getTextEncoder, getTextDecoder } from "./text";
import { DataViewWrapper } from "./data-view-wrapper";

export function alignOffsetToType(off: number, fieldType: FieldType) {
  const fieldSize = FieldAlignment[fieldType];
  return Math.floor((off + fieldSize - 1) / fieldSize) * fieldSize;
}

export function alignDynamicFieldContentOffset(off: number, fieldType: FieldType) {
  const contentAlignment = FieldDynamicContentAlignment[fieldType];
  return Math.floor((off + contentAlignment - 1) / contentAlignment) * contentAlignment;
}

export class InternalMessage {
  public bytes: Uint8Array;

  private dataView: DataView;
  private offsets: { [fieldNum: number]: number };

  constructor(buf: Uint8Array, private size: number, private scheme: FieldType[], private unions: FieldType[][]) {
    this.bytes = buf; // buf should be Uint8Array (a view over an ArrayBuffer)
    this.dataView = new DataViewWrapper(buf.buffer, buf.byteOffset);
    this.offsets = null; // map: fieldNum -> offset in bytes
  }

  private lazyCalcOffsets() {
    if (this.offsets !== null) {
      return true;
    }
    const res: any = {};
    let off = 0;
    let unionNum = 0;
    for (let fieldNum = 0; fieldNum < this.scheme.length; fieldNum++) {
      let fieldType = this.scheme[fieldNum];

      if (off === this.size) { // This means we are at end of field (but may be postfix newer fields we ignore) stop parsing
        break;
      }
      // write the current offset
      off = alignOffsetToType(off, fieldType);
      if (off > this.size) {
        return false;
      }
      res[fieldNum] = off;

      // skip over the content to the next field
      if (fieldType === FieldTypes.TypeUnion) {
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
    if (off > this.size || off === 0) { // past end of buffer or empty buffer fail
      return false;
    }
    this.offsets = res;
    return true;
  }

  isValid(): boolean {
    if (this.bytes === undefined) {
      throw `uninitialized membuffer, did you create it directly without a Builder or a Reader?`;
    }
    return this.lazyCalcOffsets();
  }

  rawBuffer(): Uint8Array {
    return this.bytes.subarray(0, this.size);
  }

  rawBufferForField(fieldNum: number, unionNum: number): Uint8Array {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length || fieldNum >= this.scheme.length) {
      return new Uint8Array();
    }
    let fieldType = this.scheme[fieldNum];
    let off = this.offsets[fieldNum];
    if (fieldType === FieldTypes.TypeUnion) {
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
      return this.bytes.subarray(off, off + contentSize);
    } else {
      return this.bytes.subarray(off, off + FieldSizes[fieldType]);
    }
  }

  rawBufferWithHeaderForField(fieldNum: number, unionNum: number): Uint8Array {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length || fieldNum >= this.scheme.length) {
      return new Uint8Array();
    }
    let fieldType = this.scheme[fieldNum];
    let off = this.offsets[fieldNum];
    const fieldHeaderOff = off;
    if (fieldType === FieldTypes.TypeUnion) {
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
      return this.bytes.subarray(fieldHeaderOff, off + contentSize);
    } else {
      return this.bytes.subarray(fieldHeaderOff, off + FieldSizes[fieldType]);
    }
  }

  getOffsetInOffset(off: number): number {
    return this.dataView.getUint32(off, true);
  }

  getUint8InOffset(off: number): number {
    return this.dataView.getUint8(off);
  }

  setUint8InOffset(off: number, v: number): void {
    return this.dataView.setUint8(off, v);
  }

  getUint8(fieldNum: number): number {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint8InOffset(off);
  }

  setUint8(fieldNum: number, v: number): void {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint8InOffset(off, v);
  }

  getUint16InOffset(off: number): number {
    return this.dataView.getUint16(off, true);
  }

  setUint16InOffset(off: number, v: number): void {
    return this.dataView.setUint16(off, v, true);
  }

  getUint16(fieldNum: number): number {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint16InOffset(off);
  }

  setUint16(fieldNum: number, v: number): void {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint16InOffset(off, v);
  }

  getUint32InOffset(off: number): number {
    return this.dataView.getUint32(off, true);
  }

  setUint32InOffset(off: number, v: number): void {
    return this.dataView.setUint32(off, v, true);
  }

  getUint32(fieldNum: number): number {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return 0;
    }
    const off = this.offsets[fieldNum];
    return this.getUint32InOffset(off);
  }

  setUint32(fieldNum: number, v: number): void {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint32InOffset(off, v);
  }

  getUint64InOffset(off: number): bigint {
    return this.dataView.getBigUint64(off, true);
  }

  setUint64InOffset(off: number, v: bigint): void {
    return this.dataView.setBigUint64(off, v, true);
  }

  getUint64(fieldNum: number): bigint {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return BigInt(0);
    }
    const off = this.offsets[fieldNum];
    return this.getUint64InOffset(off);
  }

  setUint64(fieldNum: number, v: bigint): void {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setUint64InOffset(off, v);
  }

  getMessageInOffset(off: number): Uint8Array {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeMessage];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeMessage);
    return this.bytes.subarray(off, off + contentSize);
  }

  getMessage(fieldNum: number): Uint8Array {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Uint8Array();
    }
    const off = this.offsets[fieldNum];
    return this.getMessageInOffset(off);
  }

  getBytesInOffset(off: number): Uint8Array {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeBytes];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeBytes);
    if (off + contentSize > this.bytes.byteLength) {
      return new Uint8Array();
    }
    return this.bytes.subarray(off, off + contentSize);
  }

  setBytesInOffset(off: number, v: Uint8Array) {
    const contentSize = this.dataView.getUint32(off, true);
    if (contentSize !== v.byteLength) {
      throw new Error("size mismatch");
    }
    off += FieldSizes[FieldTypes.TypeBytes];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeBytes);
    return this.bytes.set(v, off);
  }

  getBytes(fieldNum: number): Uint8Array {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Uint8Array();
    }
    const off = this.offsets[fieldNum];
    return this.getBytesInOffset(off);
  }

  setBytes(fieldNum: number, v: Uint8Array) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    const off = this.offsets[fieldNum];
    return this.setBytesInOffset(off, v);
  }

  getBytes20(fieldNum: number): Uint8Array {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Uint8Array();
    }
    const off = this.offsets[fieldNum];
    return this.bytes.subarray(off, off + 20);
  }

  setBytes20(fieldNum: number, v: Uint8Array) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    if (20 !== v.byteLength) {
      throw new Error("size mismatch");
    }
    const off = this.offsets[fieldNum];
    return this.bytes.set(v, off);
  }

  getBytes32(fieldNum: number): Uint8Array {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new Uint8Array();
    }
    const off = this.offsets[fieldNum];
    return this.bytes.subarray(off, off + 32);
  }

  setBytes32(fieldNum: number, v: Uint8Array) {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      throw new Error("invalid field");
    }
    if (32 !== v.byteLength) {
      throw new Error("size mismatch");
    }
    const off = this.offsets[fieldNum];
    return this.bytes.set(v, off);
  }

  getStringInOffset(off: number): string {
    const b = this.getBytesInOffset(off);
    return getTextDecoder().decode(b);
  }

  setStringInOffset(off: number, v: string) {
    return this.setBytesInOffset(off, getTextEncoder().encode(v));
  }

  getString(fieldNum: number): string {
    const b = this.getBytes(fieldNum);
    return getTextDecoder().decode(b);
  }

  setString(fieldNum: number, v: string) {
    return this.setBytes(fieldNum, getTextEncoder().encode(v));
  }

  getUnionIndex(fieldNum: number, unionNum: number): number {
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

  isUnionIndex(fieldNum: number, unionNum: number, unionIndex: number): [boolean, number] {
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
    return [unionType === unionIndex, off];
  }

  getUint8ArrayIteratorInOffset(off: number): ArrayIterator {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint8Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint8Array);
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeUint8, this);
  }

  getUint8ArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeUint8, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint8ArrayIteratorInOffset(off);
  }

  getUint16ArrayIteratorInOffset(off: number): ArrayIterator {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint32Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint16Array);
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeUint16, this);
  }

  getUint16ArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeUint16, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint16ArrayIteratorInOffset(off);
  }

  getUint32ArrayIteratorInOffset(off: number): ArrayIterator {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint32Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint32Array);
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeUint32, this);
  }

  getUint32ArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeUint32, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint32ArrayIteratorInOffset(off);
  }

  getUint64ArrayIteratorInOffset(off: number): ArrayIterator {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeUint64Array];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeUint64Array);
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeUint64, this);
  }

  getUint64ArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeUint64, this);
    }
    const off = this.offsets[fieldNum];
    return this.getUint64ArrayIteratorInOffset(off);
  }

  getMessageArrayIteratorInOffset(off: number): ArrayIterator {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeMessageArray];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeMessageArray);
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeMessage, this);
  }

  getMessageArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeMessage, this);
    }
    const off = this.offsets[fieldNum];
    return this.getMessageArrayIteratorInOffset(off);
  }

  getBytesArrayIteratorInOffset(off: number): ArrayIterator {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeMessageArray];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeBytesArray);
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeBytes, this);
  }

  getBytesArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeBytes, this);
    }
    const off = this.offsets[fieldNum];
    return this.getBytesArrayIteratorInOffset(off);
  }

  getBytes20ArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeBytes20, this);
    }
    let off = this.offsets[fieldNum];
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeBytes20Array];
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeBytes20, this);
  }

  getBytes32ArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeBytes32, this);
    }
    let off = this.offsets[fieldNum];
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeBytes32Array];
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeBytes32, this);
  }

  getStringArrayIteratorInOffset(off: number): ArrayIterator {
    const contentSize = this.dataView.getUint32(off, true);
    off += FieldSizes[FieldTypes.TypeStringArray];
    off = alignDynamicFieldContentOffset(off, FieldTypes.TypeStringArray);
    return new ArrayIterator(off, off + contentSize, FieldTypes.TypeString, this);
  }

  getStringArrayIterator(fieldNum: number): ArrayIterator {
    if (!this.lazyCalcOffsets() || fieldNum >= Object.keys(this.offsets).length) {
      return new ArrayIterator(0, 0, FieldTypes.TypeString, this);
    }
    const off = this.offsets[fieldNum];
    return this.getStringArrayIteratorInOffset(off);
  }
}
