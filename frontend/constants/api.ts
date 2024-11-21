export const apiEndpoints = {
  login: '/api/auth/login',
  register: '/api/auth/register'
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
}`
}; 