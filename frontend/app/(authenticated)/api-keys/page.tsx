'use client';

import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { apiEndpoints } from '@/constants/api';
import { Check } from 'lucide-react';
import { useEffect, useState } from 'react';
import { Badge } from "@/components/ui/badge";

type ApiKeyData = {
  key: string;
  uid: string;
  totalUsage: number;
  createdAt: string;
};

type ApiKeyResponse = {
  message: string;
  result: {
    apiKey: ApiKeyData;
  };
  status: string;
};

export default function ApiKeysPage() {
  const [apiKey, setApiKey] = useState<string>('');
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isVisible, setIsVisible] = useState(false);
  const [keyMetadata, setKeyMetadata] = useState({
    createdAt: new Date(),
    totalUsage: 0,
  });
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    const fetchApiKey = async () => {
      try {
        const accessToken = localStorage.getItem('accessToken');
        const response = await fetch(apiEndpoints.apiKey, {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        });

        const data: ApiKeyResponse = await response.json();

        setApiKey(data.result.apiKey.key);

        if (data.result.apiKey.key) {
          setKeyMetadata({
            createdAt: new Date(data.result.apiKey.createdAt),
            totalUsage: data.result.apiKey.totalUsage,
          });
        }
      } catch (err) {
        setError(
          err instanceof Error ? err.message : 'Failed to fetch API key'
        );
      } finally {
        setIsLoading(false);
      }
    };

    fetchApiKey();
  }, []);

  const generateNewApiKey = async () => {
    try {
      const accessToken = localStorage.getItem('accessToken');
      const response = await fetch(apiEndpoints.apiKey, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
      console.log(response);

      const data: ApiKeyResponse = await response.json();

      console.log(data);

      setApiKey(data.result.apiKey.key);
      setKeyMetadata({
        createdAt: new Date(data.result.apiKey.createdAt),
        totalUsage: 0,
      });
    } catch (err) {
      setError(
        err instanceof Error ? err.message : 'Failed to generate API key'
      );
    }
  };

  const toggleVisibility = () => {
    setIsVisible(!isVisible);
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(apiKey);
    setCopied(true);
    setTimeout(() => {
      setCopied(false);
    }, 2000);
  };

  const revokeApiKey = async () => {
    try {
      const accessToken = localStorage.getItem('accessToken');
      const response = await fetch(apiEndpoints.apiKey, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });

      if (response.ok) {
        setApiKey('');
        setKeyMetadata({
          createdAt: new Date(),
          totalUsage: 0,
        });
      } else {
        const data = await response.json();
        setError(data.message || 'Failed to revoke API key');
      }
    } catch (err) {
      setError(
        err instanceof Error ? err.message : 'Failed to revoke API key'
      );
    }
  };

  if (isLoading) {
    return (
      <div className="flex-1 p-8">
        <div className="max-w-2xl mx-auto">
          <p className="text-slate-400">Loading...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex-1 p-8">
        <div className="max-w-2xl mx-auto">
          <p className="text-red-500">Error: {error}</p>
        </div>
      </div>
    );
  }

  if (!apiKey) {
    return (
      <div className="flex-1 p-8">
        <div className="max-w-2xl mx-auto">
          <h1 className="text-3xl font-bold text-slate-100 mb-8">API Keys</h1>

          <Card className="bg-slate-900 border-slate-800">
            <CardHeader>
              <CardTitle className="text-slate-100">No API Key Found</CardTitle>
              <CardDescription>
                Generate an API key to start making authenticated requests
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Button onClick={generateNewApiKey}>Generate New API Key</Button>
            </CardContent>
          </Card>
        </div>
      </div>
    );
  }

  return (
    <div className="flex-1 p-8">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold text-slate-100 mb-8">API Keys</h1>

        <Card className="bg-slate-900 border-slate-800">
          <CardHeader>
            <CardTitle className="text-slate-100">Your API Key</CardTitle>
            <CardDescription>
              Use this key to authenticate your API requests
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="flex gap-2">
              <Input
                type={isVisible ? 'text' : 'password'}
                value={apiKey}
                readOnly
                className="font-mono bg-slate-950 border-slate-800 text-slate-100"
              />
              <Button variant="outline" onClick={toggleVisibility}>
                {isVisible ? 'Hide' : 'Show'}
              </Button>
              <Button variant="default" onClick={copyToClipboard}>
                {copied ? <Check className="h-4 w-4" /> : 'Copy'}
              </Button>
            </div>
            <div className="flex flex-wrap gap-2">
              <Badge 
                variant="secondary" 
                className="bg-blue-900/50 hover:bg-blue-900/70 text-blue-200 border border-blue-800"
              >
                Created {keyMetadata.createdAt.toLocaleDateString()}
              </Badge>
              <Badge 
                variant="secondary" 
                className="bg-purple-900/50 hover:bg-purple-900/70 text-purple-200 border border-purple-800"
              >
                {keyMetadata.totalUsage.toLocaleString()} requests
              </Badge>
            </div>
            <div className="pt-4">
              <Button 
                variant="destructive" 
                onClick={revokeApiKey}
                className="bg-red-900 hover:bg-red-800"
              >
                Revoke API Key
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
