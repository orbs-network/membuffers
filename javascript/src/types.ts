/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

export type FieldType = number;

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
  TypeBytes32: 41,
  TypeBytes20: 42,
  TypeBytes32Array: 51,
  TypeBytes20Array: 52,
});

export const FieldSizes = Object.freeze({
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
  [FieldTypes.TypeBytes32]: 32,
  [FieldTypes.TypeBytes20]: 20,
  [FieldTypes.TypeBytes32Array]: 4,
  [FieldTypes.TypeBytes20Array]: 4,
});

export const FieldAlignment = Object.freeze({
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
  [FieldTypes.TypeBytes32]: 4,
  [FieldTypes.TypeBytes20]: 4,
  [FieldTypes.TypeBytes32Array]: 4,
  [FieldTypes.TypeBytes20Array]: 4,
});

export const FieldDynamic = Object.freeze({
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
  [FieldTypes.TypeBytes32]: false,
  [FieldTypes.TypeBytes20]: false,
  [FieldTypes.TypeBytes32Array]: true,
  [FieldTypes.TypeBytes20Array]: true,
});

export const FieldDynamicContentAlignment = Object.freeze({
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
  [FieldTypes.TypeBytes32]: 0,
  [FieldTypes.TypeBytes20]: 0,
  [FieldTypes.TypeBytes32Array]: 4,
  [FieldTypes.TypeBytes20Array]: 4,
});
