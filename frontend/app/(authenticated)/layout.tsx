'use client';

import AppSidebar from '@/components/ui/custom/sidebar';
import { SidebarProvider } from '@/components/ui/sidebar';
import { Inter } from 'next/font/google';
import { redirect } from 'next/navigation';
import { useEffect } from 'react';

const inter = Inter({ subsets: ['latin'] });

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  useEffect(() => {
    const accessToken = localStorage.getItem('accessToken');
    if (!accessToken) {
      redirect('/');
    }
  }, []);

  return (
    <SidebarProvider>
      <div className="flex min-h-screen">
        <AppSidebar />
        <main className={`${inter.className} flex-1`}>{children}</main>
      </div>
    </SidebarProvider>
  );
}
