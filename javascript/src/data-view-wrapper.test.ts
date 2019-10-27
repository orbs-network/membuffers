/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import "./matcher-extensions";
import { DataViewWrapper, JDataViewWrapper } from "./data-view-wrapper";
import jDataView from "jdataview";

test("Test jDataView", () => {
    expect(() => {
        const Wrapper: any = jDataView;
        const z = new Wrapper(new ArrayBuffer(8));
        z.setUint64(0, 29273397577908224);
    }).not.toThrow();
})

test("Test DataViewWrapper polyfill", () => {
    const z = new DataViewWrapper(new ArrayBuffer(8))
    z.setBigUint64(0, BigInt("29273397577908224"));
    expect(z.getBigUint64(0)).toEqual(BigInt("29273397577908224"))
})

test("Test JDataViewWrapper", () => {
    const z = new JDataViewWrapper(new ArrayBuffer(8))
    z.setBigUint64(0, BigInt("29273397577908224"));
    expect(z.getBigUint64(0)).toEqual(BigInt("29273397577908224"))
})

test("Test DataViewWrapper polyfill (negative)", () => {
    const z = new DataViewWrapper(new ArrayBuffer(8))
    z.setBigInt64(0, BigInt("-292733"));
    expect(z.getBigInt64(0)).toEqual(BigInt("-292733"))
})

test("Test JDataViewWrapper (negative)", () => {
    const z = new JDataViewWrapper(new ArrayBuffer(8))
    z.setBigInt64(0, BigInt("-292733"));
    expect(z.getBigInt64(0)).toEqual(BigInt("-292733"))
})
