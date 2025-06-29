import http from 'k6/http';
import { check, group } from 'k6';
import { Rate } from 'k6/metrics';

// Custom metrics
const errorRate = new Rate('errors');

// Test configuration
export const options = {
  stages: [
    { duration: '30s', target: 5 },  // Ramp up
    { duration: '2m', target: 10 },  // Stay at 10 users
    { duration: '30s', target: 0 },  // Ramp down
  ],
  thresholds: {
    http_req_duration: ['p(95)<2000'], // 95% of requests under 2s
    errors: ['rate<0.1'], // Error rate under 10%
  },
};

// Base URL - can be overridden with environment variable
const BASE_URL = __ENV.API_BASE_URL || 'http://localhost:8080';

// Test data
const testContent = {
  title: 'Load Test Content',
  body: 'This is content created during load testing'
};

export default function() {
  let contentId;

  group('Content API Load Test', () => {
    // Test 1: Create Content
    group('Create Content', () => {
      const createPayload = JSON.stringify(testContent);
      const createParams = {
        headers: {
          'Content-Type': 'application/json',
        },
      };

      const createResponse = http.post(
        `${BASE_URL}/content.v1.ContentService/CreateContent`,
        createPayload,
        createParams
      );

      const createSuccess = check(createResponse, {
        'create status is 200': (r) => r.status === 200,
        'create response has content': (r) => {
          try {
            const body = JSON.parse(r.body);
            return body.content && body.content.id;
          } catch (e) {
            return false;
          }
        },
      });

      if (!createSuccess) {
        errorRate.add(1);
      } else {
        errorRate.add(0);
        try {
          const responseBody = JSON.parse(createResponse.body);
          contentId = responseBody.content.id;
        } catch (e) {
          console.error('Failed to parse create response:', e);
        }
      }
    });

    // Test 2: Get Content (if we have an ID)
    if (contentId) {
      group('Get Content', () => {
        const getPayload = JSON.stringify({ id: contentId });
        const getParams = {
          headers: {
            'Content-Type': 'application/json',
          },
        };

        const getResponse = http.post(
          `${BASE_URL}/content.v1.ContentService/GetContent`,
          getPayload,
          getParams
        );

        const getSuccess = check(getResponse, {
          'get status is 200': (r) => r.status === 200,
          'get response has content': (r) => {
            try {
              const body = JSON.parse(r.body);
              return body.content && body.content.id === contentId;
            } catch (e) {
              return false;
            }
          },
        });

        errorRate.add(getSuccess ? 0 : 1);
      });

      // Test 3: Update Content
      group('Update Content', () => {
        const updatePayload = JSON.stringify({
          id: contentId,
          title: testContent.title + ' (Updated)',
          body: testContent.body + ' - Updated during load test'
        });
        const updateParams = {
          headers: {
            'Content-Type': 'application/json',
          },
        };

        const updateResponse = http.post(
          `${BASE_URL}/content.v1.ContentService/UpdateContent`,
          updatePayload,
          updateParams
        );

        const updateSuccess = check(updateResponse, {
          'update status is 200': (r) => r.status === 200,
          'update response has content': (r) => {
            try {
              const body = JSON.parse(r.body);
              return body.content && body.content.id === contentId;
            } catch (e) {
              return false;
            }
          },
        });

        errorRate.add(updateSuccess ? 0 : 1);
      });
    }

    // Test 4: List Contents
    group('List Contents', () => {
      const listPayload = JSON.stringify({
        page_limit: 10,
        page_offset: 0
      });
      const listParams = {
        headers: {
          'Content-Type': 'application/json',
        },
      };

      const listResponse = http.post(
        `${BASE_URL}/content.v1.ContentService/ListContents`,
        listPayload,
        listParams
      );

      const listSuccess = check(listResponse, {
        'list status is 200': (r) => r.status === 200,
        'list response has contents array': (r) => {
          try {
            const body = JSON.parse(r.body);
            return Array.isArray(body.contents);
          } catch (e) {
            return false;
          }
        },
      });

      errorRate.add(listSuccess ? 0 : 1);
    });

    // Test 5: Delete Content (if we have an ID)
    if (contentId) {
      group('Delete Content', () => {
        const deletePayload = JSON.stringify({ id: contentId });
        const deleteParams = {
          headers: {
            'Content-Type': 'application/json',
          },
        };

        const deleteResponse = http.post(
          `${BASE_URL}/content.v1.ContentService/DeleteContent`,
          deletePayload,
          deleteParams
        );

        const deleteSuccess = check(deleteResponse, {
          'delete status is 200': (r) => r.status === 200,
        });

        errorRate.add(deleteSuccess ? 0 : 1);
      });
    }
  });
}