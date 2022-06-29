import { deserializeTransaction } from "@stacks/transactions";
import now from "performance-now";

const stringTx = "0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a0000000000000001000135230d6a06749377212758f925871c1106fb258abed738df874a858f6f394881786d413dccfe8b134f0a8f820ce72220b8201b1cfc43ac365cf1e23f4f4003ab030200000000000516f5f1d7c482c6e1f8330b4805c5c03d8409d32452000000000000006468656c6c6f0000000000000000000000000000000000000000000000000000000000"

const rawTx = Buffer.from(stringTx, "hex");

var times = [];

for (var i = 0; i < 500000; i++) {
  var start = now()
  deserializeTransaction(rawTx);
  var end = now()
  var lat = ((end - start) * 1000000);
  times.push(lat)
  // print nanoseconds by converting from milliseconds
  console.log(lat+"ns")
}

// print average
console.log(times.reduce((a, b) => a + b, 0) / times.length)
