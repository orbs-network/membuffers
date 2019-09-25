/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import { alignOffsetToType, alignDynamicFieldContentOffset } from "./message";
import { FieldTypes, FieldSizes } from "./types";
import { getTextEncoder } from "./text";

interface MessageWriter {
  write(buf: Uint8Array): void;
  getSize(): number;
  calcRequiredSize(): number;
}

export class InternalBuilder {
  public size: number;

  constructor() {
    this.size = 0;
  }

  reset(): void {
    this.size = 0;
  }

  getSize(): number {
    return this.size;
  }

  writeUint8(buf: Uint8Array, v: number): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint8);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint8(this.size, v);
    }
    this.size += FieldSizes[FieldTypes.TypeUint8];
  }

  writeUint16(buf: Uint8Array, v: number): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint16);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint16(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint16];
  }

  writeUint32(buf: Uint8Array, v: number): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint32);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint32];
  }

  writeUint64(buf: Uint8Array, v: bigint): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint64);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setBigUint64(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint64];
  }

  writeBytes(buf: Uint8Array, v: Uint8Array): void {
    const dataView = buf ? new DataView(buf.buffer, buf.byteOffset) : undefined;
    this.size = alignOffsetToType(this.size, FieldTypes.TypeBytes);
    if (buf) {
      if (v) {
        dataView.setUint32(this.size, v.byteLength, true);
      } else {
        dataView.setUint32(this.size, 0, true);
      }
    }
    this.size += FieldSizes[FieldTypes.TypeBytes];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeBytes);
    if (v) {
      if (buf) {
        buf.set(v, this.size);
      }
      this.size += v.byteLength;
    }
  }

  writeBytes20(buf: Uint8Array, v: Uint8Array): void {
    if (!v || 20 !== v.byteLength) {
      throw new Error("size mismatch");
    }
    this.size = alignOffsetToType(this.size, FieldTypes.TypeBytes20);
    if (v) {
      if (buf) {
        buf.set(v, this.size);
      }
      this.size += v.byteLength;
    }
  }

  writeBytes32(buf: Uint8Array, v: Uint8Array): void {
    if (!v || 32 !== v.byteLength) {
      throw new Error("size mismatch");
    }
    this.size = alignOffsetToType(this.size, FieldTypes.TypeBytes32);
    if (v) {
      if (buf) {
        buf.set(v, this.size);
      }
      this.size += v.byteLength;
    }
  }

  writeString(buf: Uint8Array, v: string): void {
    this.writeBytes(buf, getTextEncoder().encode(v));
  }

  writeUnionIndex(buf: Uint8Array, unionIndex: number): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUnion);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint16(this.size, unionIndex, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUnion];
  }

  writeUint8Array(buf: Uint8Array, v: number[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint8Array);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint8], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint8Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint8Array);
    for (const vv of v) {
      this.writeUint8(buf, vv);
    }
  }

  writeUint16Array(buf: Uint8Array, v: number[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint16Array);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint16], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint16Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint16Array);
    for (const vv of v) {
      this.writeUint16(buf, vv);
    }
  }

  writeUint32Array(buf: Uint8Array, v: number[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint32Array);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint32], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint32Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint32Array);
    for (const vv of v) {
      this.writeUint32(buf, vv);
    }
  }

  writeUint64Array(buf: Uint8Array, v: bigint[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint64Array);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint64], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint64Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint64Array);
    for (const vv of v) {
      this.writeUint64(buf, vv);
    }
  }

  writeBytesArray(buf: Uint8Array, v: Uint8Array[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeBytesArray);
    const sizePlaceholderOffset = this.size;
    this.size += FieldSizes[FieldTypes.TypeBytesArray];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeBytesArray);
    const contentSizeStartOffset = this.size;
    for (const vv of v) {
      this.writeBytes(buf, vv);
    }
    const contentSize = this.size - contentSizeStartOffset;
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(sizePlaceholderOffset, contentSize, true);
    }
  }

  writeBytes20Array(buf: Uint8Array, v: Uint8Array[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeBytes20Array);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(this.size, v.length*20, true);
    }
    this.size += FieldSizes[FieldTypes.TypeBytes20Array];
    for (const vv of v) {
      this.writeBytes20(buf, vv);
    }
  }

  writeBytes32Array(buf: Uint8Array, v: Uint8Array[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeBytes32Array);
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(this.size, v.length*32, true);
    }
    this.size += FieldSizes[FieldTypes.TypeBytes32Array];
    for (const vv of v) {
      this.writeBytes32(buf, vv);
    }
  }

  writeStringArray(buf: Uint8Array, v: string[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeStringArray);
    const sizePlaceholderOffset = this.size;
    this.size += FieldSizes[FieldTypes.TypeStringArray];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeStringArray);
    const contentSizeStartOffset = this.size;
    for (const vv of v) {
      this.writeString(buf, vv);
    }
    const contentSize = this.size - contentSizeStartOffset;
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(sizePlaceholderOffset, contentSize, true);
    }
  }

  writeMessage(buf: Uint8Array, v: MessageWriter): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeMessage);
    const sizePlaceholderOffset = this.size;
    this.size += FieldSizes[FieldTypes.TypeMessage];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeMessage);
    if (buf) {
      v.write(buf.subarray(this.size));
    } else {
      v.write(null);
    }
    const contentSize = v.getSize();
    this.size += contentSize;
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(sizePlaceholderOffset, contentSize, true);
    }
  }

  writeMessageArray(buf: Uint8Array, v: MessageWriter[]): void {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeMessageArray);
    const sizePlaceholderOffset = this.size;
    this.size += FieldSizes[FieldTypes.TypeMessageArray];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeMessageArray);
    const contentSizeStartOffset = this.size;
    for (const vv of v) {
      this.writeMessage(buf, vv);
    }
    const contentSize = this.size - contentSizeStartOffset;
    if (buf) {
      new DataView(buf.buffer, buf.byteOffset).setUint32(sizePlaceholderOffset, contentSize, true);
    }
  }
}
