import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { CodeBlock } from '@/components/ui/code-block';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';

export default function Home() {
  const searchExample = `curl -X GET "https://api.oxford-dictionary.com/api/word/search/abandon" \\
  -H "Authorization: Bearer YOUR_API_KEY"`;

  const responseExample = `{
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
}`;

  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="container mx-auto px-4 py-20 text-center">
        <Badge className="mb-4">v1.0 Now Available</Badge>
        <h1 className="text-white scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl mb-4">
          The Oxford 5000™ API
        </h1>
        <p className="text-xl text-muted-foreground max-w-2xl mx-auto mb-8">
          Access comprehensive word definitions, examples, and pronunciations
          from the Oxford 5000™ List
        </p>
        <div className="flex gap-4 justify-center">
          <Button size="lg" asChild>
            <a href="/api/swagger">View Documentation</a>
          </Button>
          <Button size="lg" variant="outline" asChild>
            <a href="/register">Get Started</a>
          </Button>
        </div>
      </section>

      {/* Features Section */}
      <section className="container mx-auto px-4 py-16">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Comprehensive Data</CardTitle>
            </CardHeader>
            <CardContent>
              Access detailed word definitions, examples, CEFR levels, and audio
              pronunciations.
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Simple Integration</CardTitle>
            </CardHeader>
            <CardContent>
              RESTful API with straightforward authentication and clear
              documentation.
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Usage Management</CardTitle>
            </CardHeader>
            <CardContent>
              Track API usage with built-in rate limiting and usage analytics.
            </CardContent>
          </Card>
        </div>
      </section>

      {/* Code Example Section */}
      <section className="container mx-auto px-4 py-16">
        <h2 className="text-3xl font-bold text-center mb-8">Quick Start</h2>
        <Tabs defaultValue="request" className="max-w-3xl mx-auto">
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="request">Request</TabsTrigger>
            <TabsTrigger value="response">Response</TabsTrigger>
          </TabsList>
          <TabsContent value="request">
            <Card>
              <CardContent className="pt-6">
                <CodeBlock language="bash" code={searchExample} />
              </CardContent>
            </Card>
          </TabsContent>
          <TabsContent value="response">
            <Card>
              <CardContent className="pt-6">
                <CodeBlock language="json" code={responseExample} />
              </CardContent>
            </Card>
          </TabsContent>
        </Tabs>
      </section>

      {/* Getting Started Steps */}
      <section className="container mx-auto px-4 py-16">
        <h2 className="text-3xl font-bold text-center mb-8">Getting Started</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>1. Create Account</CardTitle>
            </CardHeader>
            <CardContent>
              Register for an account to get started with the Oxford Dictionary
              API.
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>2. Get API Key</CardTitle>
            </CardHeader>
            <CardContent>
              Generate your API key from the dashboard to authenticate your
              requests.
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>3. Make Requests</CardTitle>
            </CardHeader>
            <CardContent>
              Start making requests to our endpoints using your API key.
            </CardContent>
          </Card>
        </div>
      </section>
    </div>
  );
}
