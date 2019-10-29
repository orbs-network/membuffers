// import jDataView, { Uint64 } from "jdataview";
// import { wrapInt64, wrapUint64, unwrapInt64, unwrapUint64 } from "./jdataview-bigint";
import { uint64ToUint8Array, uint8ArrayToUint64 } from "./bigint";

let wrapper: any = DataView;

class IOSDataViewImplemetation extends DataView {
    OFFSET_32 = 4

    setBigUint64(byteOffset: number, value: bigint, littleEndian?: boolean): void {
        let arr = uint64ToUint8Array(value);
        if (littleEndian) {
           arr = arr.reverse();
        }
    
        for (let i = 0; i < arr.length ; i++) {
            this.setUint8(byteOffset+i, arr[i])
        }
    }
    
    getBigUint64(byteOffset?: number, littleEndian?: boolean): bigint {
        let arr = new Uint8Array(8);
        for (let i = 0; i < arr.length ; i++) {
            arr[i] = this.getUint8(byteOffset+i)
        }
        
        if (littleEndian) {
            arr = arr.reverse();
        }

        return uint8ArrayToUint64(arr);
    }
}

try {
    new DataView(new ArrayBuffer(8)).setBigUint64(0, BigInt(0));
} catch (e) {
    wrapper = IOSDataViewImplemetation;
}

export const DataViewWrapper: DataViewConstructor = <DataViewConstructor>wrapper;
export const IOSDataViewWrapper: DataViewConstructor = <DataViewConstructor><unknown>IOSDataViewImplemetation;
