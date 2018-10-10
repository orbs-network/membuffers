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

}