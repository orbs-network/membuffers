/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import { FieldTypes, FieldSizes } from "./types";

// TODO https://github.com/orbs-network/membuffers/issues/35 - make this more efficient if can drop the hex string conversion
export function bigIntToUint8Array(v: bigint): Uint8Array {
    let hex = v.toString(16);
    if (64 < hex.length) {
        throw new Error("size mismatch");
    }
    if (hex.length % 2 !== 0) {
        hex = '0' + hex;
    }
    const startIndex = FieldSizes[FieldTypes.TypeUint256] - hex.length/2;
    const arr = new Uint8Array(FieldSizes[FieldTypes.TypeUint256]);
    for (let i = 0; i < hex.length; i+=2) {
        arr[startIndex + i/2] = parseInt(hex.substring(i, i+2), 16);
    }
    return arr;
}

export function uint8ArrayToBigInt(v: Uint8Array): bigint {
    let hexStr = '0x';
    for (let i = 0; i < FieldSizes[FieldTypes.TypeUint256]; i++) {
        hexStr += ('0'+v[i].toString(16)).slice(-2)
    }
    return BigInt(hexStr);
}

export function uint64ToUint8Array(v: bigint): Uint8Array {
    let hex = v.toString(16);
    // if (16 < hex.length) {
    //     throw new Error("size mismatch");
    // }
    if (hex.length % 2 !== 0) {
        hex = '0' + hex;
    }
    const startIndex = FieldSizes[FieldTypes.TypeUint64] - hex.length/2;
    const arr = new Uint8Array(FieldSizes[FieldTypes.TypeUint64]);
    for (let i = 0; i < hex.length; i+=2) {
        arr[startIndex + i/2] = parseInt(hex.substring(i, i+2), 16);
    }
    return arr;
}

export function uint8ArrayToUint64(v: Uint8Array): bigint {
    let hexStr = '0x';
    for (let i = 0; i < FieldSizes[FieldTypes.TypeUint64]; i++) {
        hexStr += ('0'+v[i].toString(16)).slice(-2)
    }
    return BigInt(hexStr);
}