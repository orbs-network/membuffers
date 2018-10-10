import {alignOffsetToType, alignDynamicFieldContentOffset} from './message';
import {FieldTypes, FieldSizes} from './types';
import {getTextEncoder} from './text';

export class InternalBuilder {

  constructor() {
    this.size = 0;
  }

  reset() {
    this.size = 0;
  }

  calcRequiredSize() {
    this.write(null);
    return this.getSize();
  }

  getSize() {
    return this.size;
  }

  writeUint8(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint8);
    if (buf) {
      new DataView(buf.buffer).setUint8(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint8];
  }

  writeUint16(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint16);
    if (buf) {
      new DataView(buf.buffer).setUint16(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint16];
  }

  writeUint32(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint32);
    if (buf) {
      new DataView(buf.buffer).setUint32(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint32];
  }

  writeUint64(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint64);
    if (buf) {
      new DataView(buf.buffer).setBigUint64(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint64];
  }

  writeBytes(buf, v) {
    const dataView = (buf) ? new DataView(buf.buffer) : undefined;
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

  writeString(buf, v) {
    this.writeBytes(buf, getTextEncoder().encode(v));
  }

  writeUnionIndex(buf, unionIndex) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUnion);
    if (buf) {
      new DataView(buf.buffer).setUint16(this.size, unionIndex, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUnion];
  }

  writeUint8Array(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint8Array);
    if (buf) {
      new DataView(buf.buffer).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint8], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint8Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint8Array);
    for (const vv of v) {
      this.writeUint8(buf, vv);
    }
  }

  writeUint16Array(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint16Array);
    if (buf) {
      new DataView(buf.buffer).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint16], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint16Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint16Array);
    for (const vv of v) {
      this.writeUint16(buf, vv);
    }
  }

  writeUint32Array(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint32Array);
    if (buf) {
      new DataView(buf.buffer).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint32], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint32Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint32Array);
    for (const vv of v) {
      this.writeUint32(buf, vv);
    }
  }

  writeUint64Array(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint64Array);
    if (buf) {
      new DataView(buf.buffer).setUint32(this.size, v.length * FieldSizes[FieldTypes.TypeUint64], true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint64Array];
    this.size = alignDynamicFieldContentOffset(this.size, FieldTypes.TypeUint64Array);
    for (const vv of v) {
      this.writeUint64(buf, vv);
    }
  }

}