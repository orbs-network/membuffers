/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import "./matcher-extensions";
import { DataViewWrapper, IOSDataViewWrapper } from "./data-view-wrapper";

const HugeNumber = BigInt("29273397577908229");

test("Test DataViewWrapper polyfill", () => {
    const z = new DataViewWrapper(new ArrayBuffer(8));
    z.setBigUint64(0, HugeNumber);
    expect(z.getBigUint64(0)).toEqual(HugeNumber);
})

test("Test IOSDataViewWrapper", () => {
    const z = new IOSDataViewWrapper(new ArrayBuffer(8))
    z.setBigUint64(0, HugeNumber);
    expect(z.getBigUint64(0).toString()).toEqual(HugeNumber.toString());
})

test("Contract test: custom to native, big indian", () => {
    const buffer = new ArrayBuffer(8);
    const customWrapper = new IOSDataViewWrapper(buffer);
    customWrapper.setBigUint64(0, HugeNumber);
    expect(customWrapper.getBigUint64(0)).toEqual(HugeNumber);

    const nativeWrapper = new DataViewWrapper(buffer);
    expect(nativeWrapper.getBigUint64(0)).toEqual(HugeNumber);
})

test("Contract test: native to custom, big indian", () => {
    const buffer = new ArrayBuffer(8);
    const nativeWrapper = new DataViewWrapper(buffer);
    nativeWrapper.setBigUint64(0, HugeNumber);
    expect(nativeWrapper.getBigUint64(0)).toEqual(HugeNumber);

    const customWrapper = new IOSDataViewWrapper(buffer);
    expect(customWrapper.getBigUint64(0)).toEqual(HugeNumber);
})

test("Contract test: custom to native, little indian", () => {
    const buffer = new ArrayBuffer(8);
    const customWrapper = new IOSDataViewWrapper(buffer);
    customWrapper.setBigUint64(0, HugeNumber, true);
    expect(customWrapper.getBigUint64(0, true)).toEqual(HugeNumber);

    const nativeWrapper = new DataViewWrapper(buffer);
    expect(nativeWrapper.getBigUint64(0, true)).toEqual(HugeNumber);
})

test("Contract test: native to custom, little indian", () => {
    const buffer = new ArrayBuffer(8);
    const nativeWrapper = new DataViewWrapper(buffer);
    nativeWrapper.setBigUint64(0, HugeNumber, true);
    expect(nativeWrapper.getBigUint64(0, true)).toEqual(HugeNumber);

    const customWrapper = new IOSDataViewWrapper(buffer);
    expect(customWrapper.getBigUint64(0, true)).toEqual(HugeNumber);
})