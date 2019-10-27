import jDataView from "jdataview";

let wrapper: any = DataView;

class JDataViewImplemetation extends jDataView {
    setBigInt64(byteOffset: number, value: bigint, littleEndian?: boolean): void {
        let lo: number = Number(value & BigInt("0xffffffff"));
        let hi: number = Number((value >> BigInt(32)) & BigInt("0xffffffff"));

        this.setInt64(byteOffset, new jDataView.Int64(lo, hi), littleEndian)
    }

    setBigUint64(byteOffset: number, value: bigint, littleEndian?: boolean): void {
        let lo: number = Number(value & BigInt("0xffffffff"));
        let hi: number = Number((value >> BigInt(32)) & BigInt("0xffffffff"));

        this.setUint64(byteOffset, new jDataView.Uint64(lo, hi), littleEndian)
    }

    getBigInt64(byteOffset?: number, littleEndian?: boolean): bigint {
        const { lo, hi } = this.getInt64(byteOffset, littleEndian);

        if (hi > Math.pow(2, 31)) {
            const p32 = BigInt(Math.pow(2, 32));
            return -((p32 - BigInt(lo)) + p32* (p32 - BigInt(1) - BigInt(hi)));
        }

        return BigInt(lo) + (BigInt(hi) << BigInt(32))
    }
    
    getBigUint64(byteOffset?: number, littleEndian?: boolean): bigint {
        const { lo, hi } = this.getInt64(byteOffset, littleEndian);

        return BigInt(lo) + (BigInt(hi) << BigInt(32))
    }
}

try {
    new DataView(new ArrayBuffer(8)).setBigUint64(0, BigInt(0));
} catch (e) {
    wrapper = JDataViewImplemetation;
}

export const DataViewWrapper: DataViewConstructor = <DataViewConstructor>wrapper;
export const JDataViewWrapper: DataViewConstructor = <DataViewConstructor><unknown>JDataViewImplemetation;
