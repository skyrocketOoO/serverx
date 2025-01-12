import http from 'k6/http';


// smoke
export const options = {
  scenarios: {
    smoke: {
      executor: 'constant-vus',
      vus: 1, // Minimal load
      duration: '10s', // Short duration
    },
  },
};

// Average-Load
export const options = {
  scenarios: {
    average_load: {
      executor: 'constant-vus',
      vus: 50, // Expected average number of users
      duration: '5m', // Simulate normal conditions for 5 minutes
    },
  },
};

// Stress
export const options = {
  scenarios: {
    stress: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '1m', target: 100 }, // Ramp up to 100 users
        { duration: '3m', target: 200 }, // Peak load with 200 users
        { duration: '1m', target: 0 },   // Ramp down to 0 users
      ],
    },
  },
};

// Soak
export const options = {
  scenarios: {
    soak: {
      executor: 'constant-vus',
      vus: 20, // Moderate load
      duration: '12h', // Extended duration
    },
  },
};

// Spike
export const options = {
  scenarios: {
    spike: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '10s', target: 200 }, // Sudden spike to 200 users
        { duration: '1m', target: 200 },  // Maintain spike load
        { duration: '10s', target: 0 },   // Drop back to 0 users
      ],
    },
  },
};

// Breakpoint 
export const options = {
  scenarios: {
    breakpoint: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '1m', target: 50 },  // Ramp up to 50 users
        { duration: '1m', target: 100 }, // Ramp up to 100 users
        { duration: '1m', target: 200 }, // Ramp up to 200 users
        { duration: '1m', target: 300 }, // Gradually increase further
      ],
    },
  },
};

export default function () {
  http.get('http://example.com');
}
