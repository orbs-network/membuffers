/**
 * Copyright 2019 the orbs-client-sdk-javascript authors
 * This file is part of the orbs-client-sdk-javascript library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

import "./matcher-extensions";
import { wrapInt64, wrapUint64, unwrapInt64, unwrapUint64 } from "./jdataview-bigint";

const HUGE_POSITIVE_NUMBER = "29273397577908229";
const SMALL_NEGATIVE_NUMBER = "-292733";

test("Test jDataView int64 wrapping", () => {
    const positive = BigInt(HUGE_POSITIVE_NUMBER);
    console.log(wrapInt64(positive).toString(), "vs", positive.toString())
    expect(wrapInt64(positive).toString()).toEqual(positive.toString());

    const negative = BigInt(SMALL_NEGATIVE_NUMBER);
    console.log(wrapInt64(negative).toString(), "vs", negative.toString())
    expect(wrapInt64(negative).toString()).toEqual(negative.toString());
})

test("Test jDataView uint64 wrapping", () => {
    const positive = BigInt(HUGE_POSITIVE_NUMBER);
    console.log(wrapUint64(positive).toString(), "vs", positive.toString())
    expect(wrapUint64(positive).toString()).toEqual(positive.toString());
})

test("Test jDataView int64 unwrapping", () => {
    const positive = BigInt(HUGE_POSITIVE_NUMBER);
    const wrappedPositive = wrapInt64(positive);
    console.log(unwrapInt64(wrappedPositive).toString(), "vs", positive.toString())
    expect(unwrapInt64(wrappedPositive).toString()).toEqual(positive.toString());

    const negative = BigInt(SMALL_NEGATIVE_NUMBER);
    const wrappedNegative = wrapInt64(negative);
    console.log(unwrapInt64(wrappedNegative).toString(), "vs", negative.toString())
    expect(unwrapInt64(wrappedNegative).toString()).toEqual(negative.toString());
})

test("Test jDataView uint64 unwrapping", () => {
    const positive = BigInt(HUGE_POSITIVE_NUMBER);
    const wrappedPositive = wrapUint64(positive);
    console.log(unwrapUint64(wrappedPositive).toString(), "vs", positive.toString())
    expect(unwrapUint64(wrappedPositive).toString()).toEqual(positive.toString());
})