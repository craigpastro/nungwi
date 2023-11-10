import http from "k6/http";
import exec from "k6/execution";
import { check } from "k6";

const baseUrl = __ENV.API_BASE_URL || "http://localhost:8080";
const scenario = __ENV.SCENARIO || "constantArrivalRate";
const rate = __ENV.RATE || 600;
const numFolders = __ENV.NUM_FOLDERS || 100;
const numDocumentsPerFolder = __ENV.NUM_DOCUMENTS_PER_FOLDER || 100;
const maxTuplesPerWrite = __ENV.MAX_TUPLES_PER_WRITE || 1000;

const scenarios = {
  constantArrivalRate: {
    executor: "constant-arrival-rate",
    duration: "5m",
    rate,
    timeUnit: "1s", // rate per timeUnit
    preAllocatedVUs: 10,
    maxVUs: 500,
  },
  rampingArrivalRate: {
    executor: "ramping-arrival-rate",
    startRate: 0,
    timeUnit: "1s",
    preAllocatedVUs: 10,
    maxVUs: 500,
    stages: [
      { target: rate, duration: "1m" }, // ramp up to target rps
      { target: rate, duration: "5m" }, // stay at target rps
      { target: 0, duration: "1m" }, // ramp down to 0 rps
    ],
  },
};

export const options = {
  summaryTrendStats: ["avg", "min", "med", "max", "p(95)", "p(99)"],
  setupTimeout: "10m",
  teardownTimeout: "10m",
  scenarios: {
    [scenario]: scenarios[scenario],
  },
};

const params = {
  headers: {"Content-Type": "application/json"}
}

export function setup() {
  // Write the schema
  const configs = [
    {"namespace": "folder", "relation": "parent", "rewrite": "self"},
    {"namespace": "folder", "relation": "viewer", "rewrite": "union(self, tupleToUserset(parent, viewer))"},
    {"namespace": "document", "relation": "parent", "rewrite": "self"},
    {"namespace": "document", "relation": "viewer", "rewrite": "union(self, tupleToUserset(parent, viewer))"},
  ];

  let res = http.post(
    `${baseUrl}/nungwi.v1alpha.NungwiService/WriteSchema`,
    JSON.stringify({ configs }),
    params,
  );
  check(res, { "write schema status was 200": (r) => r.status === 200 });
  if (res.status !== 200) {
    exec.test.abort("failed to write schema");
  }

  const tuples = [];
  const checks = [];

  for (let i = 0; i < numFolders; i++) {
    tuples.push(newTuple("folder", `${i}`, "viewer", "anya"));

    for (let j = 0; j < numDocumentsPerFolder; j++) {
      tuples.push(newTuple("document", `${i}-${j}`, "parent", `object(folder, ${i})`));
      checks.push(newCheck("document", `${i}-${j}`, "viewer", "anya"));
    }
  }

  // Chunk in the tuples.
  for (let i = 0; i < tuples.length; i += maxTuplesPerWrite) {
    const chunk = tuples.slice(i, i + maxTuplesPerWrite);

    res = http.post(
      `${baseUrl}/nungwi.v1alpha.NungwiService/WriteTuples`,
      JSON.stringify({ tuples: chunk }),
      params,
    );
    check(res, { "write tuples status was 200": (r) => r.status === 200 });
    if (res.status !== 200) {
      exec.test.abort("failed to write tuples");
    }
  }

  return { tuples, checks };
}

export default function (data) {
  // Choose a random check request
  const req = data.checks[Math.floor(Math.random() * data.checks.length)];

  const res = http.post(
    `${baseUrl}/nungwi.v1alpha.NungwiService/Check`,
    JSON.stringify(req),
    params,
  );
  check(res, { "check status was 200": (r) => r.status === 200 });

  const allowed = res.json().allowed;
  check(allowed, { "allowed was true": (allowed) => allowed === true });

  return;
}

export function teardown(data) {
  const tuples = data.tuples;

  for (let i = 0; i < tuples.length; i += maxTuplesPerWrite) {
    const chunk = tuples.slice(i, i + maxTuplesPerWrite);

    const res = http.post(
      `${baseUrl}/nungwi.v1alpha.NungwiService/DeleteTuples`,
      JSON.stringify({ tuples: chunk }),
      params,
    );
    check(res, { "delete tuples status was 200": (r) => r.status === 200 });
  }

  const res = http.post(
    `${baseUrl}/nungwi.v1alpha.NungwiService/DeleteSchema`,
    JSON.stringify({}),
    params,
  );
  check(res, { "delete schema status was 200": (r) => r.status === 200 });
}

function newTuple(namespace, id, relation, user) {
  return { namespace, id, relation, user };
}

function newCheck(namespace, id, relation, user) {
  return { namespace, id, relation, user };
}
