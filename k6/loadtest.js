import http from "k6/http";
import { check, sleep } from "k6";
import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

// Test configuration
// Uncomment the scenario you want to run

// Smoke test: Verify the system works under minimal load
// export const options = {
//   stages: [
//     { duration: '1m', target: 10 },
//   ],
//   thresholds: {
//     http_req_duration: ['p(95)<500'],
//     http_req_failed: ['rate<0.01'],
//   },
// };

// Load test: Normal load performance
// export const options = {
//     stages: [
//         { duration: "30s", target: 50 }, // Ramp up to 50 users
//         { duration: "1m", target: 50 }, // Stay at 50 users
//         { duration: "30s", target: 100 }, // Ramp up to 100 users
//         { duration: "1m", target: 100 }, // Stay at 100 users
//         { duration: "30s", target: 0 }, // Ramp down to 0 users
//     ],
//     thresholds: {
//         http_req_duration: ["p(95)<1000", "p(99)<2000"],
//         http_req_failed: ["rate<0.05"],
//     },
// };

// Stress test: Find breaking point
export const options = {
    stages: [
        { duration: "1m", target: 100 },
        { duration: "2m", target: 200 },
        { duration: "2m", target: 300 },
        { duration: "2m", target: 400 },
        { duration: "1m", target: 0 },
    ],
    thresholds: {
        http_req_duration: ["p(95)<2000"],
        http_req_failed: ["rate<0.1"],
    },
};

// Spike test: Sudden traffic increase
// export const options = {
//   stages: [
//     { duration: '10s', target: 50 },
//     { duration: '1m', target: 50 },
//     { duration: '10s', target: 500 },  // Spike!
//     { duration: '1m', target: 500 },
//     { duration: '10s', target: 50 },
//     { duration: '1m', target: 50 },
//     { duration: '10s', target: 0 },
//   ],
// };

const BASE_URL = __ENV.BASE_URL || "http://localhost:1323";
const ENDPOINT = "/api/v1/favorite";

export default function () {
    const payload = JSON.stringify({
        userId: uuidv4(),
        favNum: Math.floor(Math.random() * 100) + 1, // Random number between 1-100
    });

    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };

    const response = http.post(`${BASE_URL}${ENDPOINT}`, payload, params);

    check(response, {
        "status is 200": (r) => r.status === 200,
        "response time < 500ms": (r) => r.timings.duration < 500,
        "response time < 1000ms": (r) => r.timings.duration < 1000,
    });

    sleep(0.1); // 100ms think time between requests
}

// Optional: Test with invalid data to verify error handling
export function testInvalidData() {
    const invalidPayloads = [
        { userId: "invalid-uuid", favNum: 42 },
        { userId: uuidv4(), favNum: -1 },
        { userId: uuidv4(), favNum: 0 },
        { userId: "", favNum: 42 },
    ];

    invalidPayloads.forEach((payload) => {
        const response = http.post(
            `${BASE_URL}${ENDPOINT}`,
            JSON.stringify(payload),
            {
                headers: { "Content-Type": "application/json" },
            },
        );

        check(response, {
            "invalid request returns 400": (r) => r.status === 400,
        });
    });
}
