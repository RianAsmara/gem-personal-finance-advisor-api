
import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
  stages: [
    { duration: '30s', target: 100000 }, // ramp up to 100k virtual users (VUs)
    { duration: '1m', target: 100000 },  // sustain 100k VUs for 1 minute
    { duration: '30s', target: 0 },      // ramp down
  ],
  thresholds: {
    'http_req_duration': ['p(95)<500'],  // 95% of requests must complete below 500ms
  },
};

export default function () {
  http.get('http://localhost:9999/livez');
  sleep(0.001); // Minimal sleep to allow for more requests per VU
}
