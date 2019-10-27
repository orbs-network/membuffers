import jDataView from "jdataview";

export function wrapInt64(value: bigint): jDataView.Int64 {
    let lo: number = Number(value & BigInt("0xffffffff"));
    let hi: number = Number((value >> BigInt(32)) & BigInt("0xffffffff"));

    return new jDataView.Int64(lo, hi)
}

export function wrapUint64(value: bigint): jDataView.Uint64 {
    let lo: number = Number(value & BigInt("0xffffffff"));
    let hi: number = Number((value >> BigInt(32)) & BigInt("0xffffffff"));

    return new jDataView.Int64(lo, hi)
}

export function unwrapInt64(value: jDataView.Int64): bigint {
    const { lo, hi } = value;

    if (hi > Math.pow(2, 31)) {
        const p32 = BigInt(Math.pow(2, 32));
        return -((p32 - BigInt(lo)) + p32* (p32 - BigInt(1) - BigInt(hi)));
    }

    return BigInt(lo) + (BigInt(hi) << BigInt(32))
}

export function unwrapUint64(value: jDataView.Uint64): bigint {
    const { lo, hi } = value;
    return BigInt(lo) + (BigInt(hi) << BigInt(32))
}