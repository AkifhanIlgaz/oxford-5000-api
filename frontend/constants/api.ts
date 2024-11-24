const apiBaseUrl = 'http://localhost:8080';

export const apiEndpoints = {
  login: `${apiBaseUrl}/api/auth/login`,
  register: `${apiBaseUrl}/api/auth/register`,
  apiKey: `${apiBaseUrl}/api/user/api-key`,
  documentation: `${apiBaseUrl}/api/swagger/index.html`,
  todayUsage: `${apiBaseUrl}/api/user/api-key/usage/today`,
  totalUsage: `${apiBaseUrl}/api/user/api-key/usage/total`,
};

export const apiExamples = {
  searchRequest: `curl -X GET "https://api.oxford-dictionary.com/api/word/search/abandon" \\
  -H "Authorization: Bearer YOUR_API_KEY"`,
  searchResponse: `{
  "success": true,
  "message": "Word found",
  "data": {
    "word": "example",
    "header": {
      "partOfSpeech": "noun",
      "cefrLevel": "B1",
      "audio": {
        "uk": "uk_pronunciation_url",
        "us": "us_pronunciation_url"
      }
    },
    "definitions": [
      {
        "meaning": "something that shows what something is like",
        "examples": ["This painting is a perfect example of her early work"]
      }
    ]
  }
}`,
};

export const routes = {
  home: '/',
  login: '/login',
  register: '/register',
  dashboard: '/dashboard',
  documentation: '/api/swagger',
  forgotPassword: '/forgot-password',
};
