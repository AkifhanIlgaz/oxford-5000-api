'use client';

import { Card } from '@/components/ui/card';

export default function DashboardPage() {
  return (
    <main className="container mx-auto p-6">
      <h1 className="text-3xl font-bold text-slate-100 mb-6">Dashboard</h1>
      <Card className="p-6 bg-white/5 border-slate-800">
        <p className="text-slate-300">Welcome to your dashboard</p>
      </Card>
    </main>
  );
}
