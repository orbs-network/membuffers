import jDataView from "jdataview";
import { wrapInt64, wrapUint64, unwrapInt64, unwrapUint64 } from "./jdataview-bigint";

let wrapper: any = DataView;

class JDataViewImplemetation extends jDataView {
    setBigInt64(byteOffset: number, value: bigint, littleEndian?: boolean): void {
        this.setInt64(byteOffset, wrapInt64(value), littleEndian)
    }

    setBigUint64(byteOffset: number, value: bigint, littleEndian?: boolean): void {
        this.setUint64(byteOffset, wrapUint64(value), littleEndian)
    }

    getBigInt64(byteOffset?: number, littleEndian?: boolean): bigint {
        return unwrapInt64(this.getInt64(byteOffset, littleEndian));
    }
    
    getBigUint64(byteOffset?: number, littleEndian?: boolean): bigint {
        return unwrapUint64(this.getUint64(byteOffset, littleEndian))
    }
}

try {
    new DataView(new ArrayBuffer(8)).setBigUint64(0, BigInt(0));
} catch (e) {
    wrapper = JDataViewImplemetation;
}

export const DataViewWrapper: DataViewConstructor = <DataViewConstructor>wrapper;
export const JDataViewWrapper: DataViewConstructor = <DataViewConstructor><unknown>JDataViewImplemetation;
