'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { type ChartConfig } from '@/components/ui/chart';
import { apiEndpoints } from '@/constants/api';
import { useEffect, useState } from 'react';

interface UsageData {
  date: string;
  count: number;
}

const chartConfig = {
  count: {
    label: 'API Calls',
    color: 'hsl(var(--chart-1))',
  },
} satisfies ChartConfig;

export default function UsagePage() {
  const [todayUsage, setTodayUsage] = useState(0);
  const [totalUsage, setTotalUsage] = useState(0);

  useEffect(() => {
    const fetchUsageData = async () => {
      try {
        const [todayResponse, totalResponse] = await Promise.all([
          fetch(apiEndpoints.todayUsage, {
            headers: {
              Authorization: `Bearer ${localStorage.getItem('accessToken')}`,
            },
          }),
          fetch(apiEndpoints.totalUsage, {
            headers: {
              Authorization: `Bearer ${localStorage.getItem('accessToken')}`,
            },
          }),
        ]);

        const todayData = await todayResponse.json();
        const totalData = await totalResponse.json();

        console.log(todayData, totalData);

        setTodayUsage(todayData.result.usage);
        setTotalUsage(totalData.result.usage);
      } catch (error) {
        console.error('Error fetching usage data:', error);
      }
    };

    fetchUsageData();
  }, []);

  return (
    <div className="container mx-auto p-6 space-y-6">
      <h1 className="text-2xl font-bold mb-6">API Usage Statistics</h1>

      <div className="grid gap-6 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Total API Calls</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{totalUsage}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Today&apos;s Usage</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{todayUsage}</div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
