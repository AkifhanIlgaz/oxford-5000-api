import AppSidebar from '@/components/ui/custom/sidebar';
import { SidebarProvider } from '@/components/ui/sidebar';
import { Inter } from 'next/font/google';

const inter = Inter({ subsets: ['latin'] });

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <SidebarProvider>
      <div className="flex min-h-screen">
        <AppSidebar />
        <main className={`flex-1 ${inter.className}`}>{children}</main>
      </div>
    </SidebarProvider>
  );
}
