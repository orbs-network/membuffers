import {alignOffsetToType} from './message';
import {FieldTypes, FieldSizes} from './types';

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
    if (buf !== null) {
      new DataView(buf.buffer).setUint8(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint8];
  }

  writeUint16(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint16);
    if (buf !== null) {
      new DataView(buf.buffer).setUint16(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint16];
  }

  writeUint32(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint32);
    if (buf !== null) {
      new DataView(buf.buffer).setUint32(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint32];
  }

  writeUint64(buf, v) {
    this.size = alignOffsetToType(this.size, FieldTypes.TypeUint64);
    if (buf !== null) {
      new DataView(buf.buffer).setBigUint64(this.size, v, true);
    }
    this.size += FieldSizes[FieldTypes.TypeUint64];
  }

}