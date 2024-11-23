'use client';

import { Alert, AlertDescription } from '@/components/ui/alert';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Eye, EyeOff, LogIn } from 'lucide-react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { authMessages } from '@/constants/auth';
import { routes } from '@/constants/navigation';
import { apiEndpoints } from '@/constants/api';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);

    try {
      const response = await fetch(apiEndpoints.login, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (response.ok) {
        router.push(routes.dashboard);
      } else {
        const data = await response.json();
        setError(data.message || authMessages.loginFailed);
      }
    } catch (error) {
      console.log(error);
      setError(authMessages.loginError);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <>
      <div className="text-center mb-12">
        <h2 className="text-3xl font-display font-bold text-slate-200 mb-3">
          {authMessages.welcomeBack}
        </h2>
      </div>

      <Card className="w-full bg-transparent border-none">
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            {error && (
              <Alert variant="destructive">
                <AlertDescription className="font-medium">
                  {error}
                </AlertDescription>
              </Alert>
            )}
            <div className="space-y-2">
              <Label htmlFor="email" className="text-slate-200 font-medium">
                {authMessages.emailLabel}
              </Label>
              <Input
                id="email"
                type="email"
                placeholder={authMessages.emailPlaceholder}
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
                className="bg-white border-slate-200 font-light"
              />
            </div>
            <div className="space-y-2 pb-2">
              <Label htmlFor="password" className="text-slate-200">
                {authMessages.passwordLabel}
              </Label>
              <div className="relative">
                <Input
                  id="password"
                  type={showPassword ? 'text' : 'password'}
                  value={password}
                  placeholder={authMessages.passwordPlaceholder}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                  className="bg-white border-slate-200 pr-10"
                />
                <Button
                  type="button"
                  variant="ghost"
                  size="sm"
                  className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? (
                    <EyeOff className="h-4 w-4 text-slate-400" />
                  ) : (
                    <Eye className="h-4 w-4 text-slate-400" />
                  )}
                </Button>
              </div>
            </div>

            <Button
              type="submit"
              className="w-full bg-blue-600 hover:bg-blue-700 text-white "
              disabled={isLoading}
            >
              {isLoading ? (
                authMessages.loggingIn
              ) : (
                <>
                  <LogIn className="mr-2 h-4 w-4" /> {authMessages.logIn}
                </>
              )}
            </Button>
          </form>
        </CardContent>
        <CardFooter className="flex flex-col gap-4">
          <Button
            variant="outline"
            className="w-full border-slate-200 text-slate-900 hover:bg-slate-100 font-medium"
            asChild
          >
            <Link href="/register">{authMessages.noAccount}</Link>
          </Button>
        </CardFooter>
      </Card>
    </>
  );
}
