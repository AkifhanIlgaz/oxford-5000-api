import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { CodeBlock } from '@/components/ui/code-block';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { apiExamples } from '@/constants/api';
import { appInfo, features, gettingStarted } from '@/constants/common';
import { routes } from '@/constants/navigation';

export default function Home() {
  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="container mx-auto px-4 py-20 text-center">
        <Badge className="mb-4">{appInfo.versionBadge}</Badge>
        <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl mb-4 text-white">
          {appInfo.title}
        </h1>
        <p className="text-xl text-muted-foreground max-w-2xl mx-auto mb-8">
          {appInfo.description}
        </p>
        <div className="flex gap-4 justify-center">
          <Button size="lg" asChild>
            <a href={routes.documentation}>View Documentation</a>
          </Button>
          <Button size="lg" variant="outline" asChild>
            <a href={routes.register}>Get Started</a>
          </Button>
        </div>
      </section>

      {/* Features Section */}
      <section className="container mx-auto px-4 py-16">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>{features.comprehensiveData.title}</CardTitle>
            </CardHeader>
            <CardContent>{features.comprehensiveData.description}</CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>{features.simpleIntegration.title}</CardTitle>
            </CardHeader>
            <CardContent>{features.simpleIntegration.description}</CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>{features.usageManagement.title}</CardTitle>
            </CardHeader>
            <CardContent>{features.usageManagement.description}</CardContent>
          </Card>
        </div>
      </section>

      {/* Code Example Section */}
      <section className="container mx-auto px-4 py-16">
        <h2 className="text-3xl font-bold text-center mb-8 text-white">
          Quick Start
        </h2>
        <Tabs defaultValue="request" className="max-w-3xl mx-auto">
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="request">Request</TabsTrigger>
            <TabsTrigger value="response">Response</TabsTrigger>
          </TabsList>
          <TabsContent value="request">
            <Card>
              <CardContent className="pt-6">
                <CodeBlock language="bash" code={apiExamples.searchRequest} />
              </CardContent>
            </Card>
          </TabsContent>
          <TabsContent value="response">
            <Card>
              <CardContent className="pt-6">
                <CodeBlock language="json" code={apiExamples.searchResponse} />
              </CardContent>
            </Card>
          </TabsContent>
        </Tabs>
      </section>

      {/* Getting Started Steps */}
      <section className="container mx-auto px-4 py-16">
        <h2 className="text-3xl font-bold text-center mb-8 text-white">
          Getting Started
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>{gettingStarted.createAccount.title}</CardTitle>
            </CardHeader>
            <CardContent>
              {gettingStarted.createAccount.description}
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>{gettingStarted.getApiKey.title}</CardTitle>
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
