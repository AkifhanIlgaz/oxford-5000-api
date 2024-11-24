'use client';

import { AnalyticsIcon } from '@/components/icons/analytics';
import { BillingIcon } from '@/components/icons/billing';
import { DocumentationIcon } from '@/components/icons/documentation';
import { KeyIcon } from '@/components/icons/key';
import { SignOutIcon } from '@/components/icons/signout';
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar';
import { appInfo } from '@/constants/common';
import { cn } from '@/lib/utils';
import Link from 'next/link';
import { redirect, usePathname } from 'next/navigation';

const items = [
  {
    title: 'Getting Started',
    items: [
      { title: 'API Keys', url: '/api-keys', icon: KeyIcon },
      { title: 'Usage Statistics', url: '/usage', icon: AnalyticsIcon },
    ],
  },
  {
    title: 'Billing',
    items: [{ title: 'Billing', url: '/billing', icon: BillingIcon }],
  },
];

const AppSidebar = () => {
  const pathname = usePathname();

  const handleSignOut = async () => {
    try {
      localStorage.removeItem('accessToken');
      localStorage.removeItem('refreshToken');
      redirect('/');
      // Handle successful sign out (e.g., redirect to login)
    } catch (error) {
      console.error('Sign out failed:', error);
    }
  };

  return (
    <Sidebar className="">
      <SidebarContent className="pb-12 bg-gradient-to-b from-slate-950 via-slate-900 to-slate-950">
        <div className="px-6 py-4">
          <h2 className="text-2xl font-bold text-foreground text-white">
            {appInfo.title}
          </h2>
        </div>

        {/* Documentation Button */}

        {/* Navigation Groups */}
        {items.map((group) => (
          <SidebarGroup key={group.title} className="px-3 py-2  text-slate-300">
            <SidebarGroupLabel className="mb-2 px-4 text-sm font-semibold text-muted-foreground">
              {group.title}
            </SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {group.items.map((item) => (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton
                      asChild
                      className={cn(
                        'w-full rounded-lg px-4 py-2 text-sm font-medium hover:bg-muted',
                        pathname === item.url &&
                          'bg-primary text-primary-foreground hover:bg-primary/90'
                      )}
                    >
                      <Link href={item.url} className="text-slate-200">
                        <item.icon />
                        <span className="flex-grow font-inter text-[14px]  font-semibold leading-6 whitespace-nowrap">
                          {item.title}
                        </span>
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        ))}

        <SidebarGroup className="mt-auto px-3 py-2">
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem>
                <SidebarMenuButton
                  asChild
                  className="w-full py-6 rounded-2xl inline-flex items-center justify-center relative box-border cursor-pointer select-none appearance-none font-semibold font-inter text-sm leading-7 min-w-[64px] text-[#f6f6f6] hover:text-[#f6f6f6] bg-[#5F61F2] shadow-[0_3px_1px_-2px_rgba(0,0,0,0.2),0_2px_2px_0_rgba(0,0,0,0.14),0_1px_5px_0_rgba(0,0,0,0.12)]  transition-all duration-250 hover:bg-[rgb(67,56,202)]"
                >
                  <Link href="/documentation">
                    <DocumentationIcon />
                    Documentation
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
              <SidebarMenuItem>
                <SidebarMenuButton
                  onClick={handleSignOut}
                  className="w-full py-6 rounded-2xl inline-flex items-center justify-center relative box-border cursor-pointer select-none appearance-none font-semibold font-inter text-sm leading-7 min-w-[64px] text-[#f6f6f6] hover:text-[#f6f6f6] bg-rose-500 shadow-[0_3px_1px_-2px_rgba(0,0,0,0.2),0_2px_2px_0_rgba(0,0,0,0.14),0_1px_5px_0_rgba(0,0,0,0.12)] transition-all duration-250 hover:bg-rose-600"
                >
                  <SignOutIcon />
                  Sign Out
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
};

export default AppSidebar;
