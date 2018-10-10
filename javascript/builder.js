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
      new DataView(buf.buffer).setUint8(this.size, v);
    }
    this.size += FieldSizes[FieldTypes.TypeUint8];
  }

}