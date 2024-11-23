import { appInfo, features } from '@/constants/common';
import { ReactNode } from 'react';

export default function AuthLayout({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen flex items-center justify-center p-4 ">
      <div className="w-full max-w-7xl flex items-center gap-20 px-8">
        {/* Left Column */}
        <div className="flex-1 max-w-md">{children}</div>

        {/* Divider */}
        <div className="h-96 w-px bg-slate-800" />

        {/* Right Column - Project Info */}
        <div className="flex-1 max-w-md text-slate-200 space-y-8">
          <h1 className="text-5xl font-bold font-display tracking-tight">
            {appInfo.title}
          </h1>
          <div className="space-y-6">
            <p className="text-2xl text-slate-300 font-display">
              {appInfo.subtitle}
            </p>
            <p className="text-slate-400 text-lg font-light leading-relaxed">
              {appInfo.description}
            </p>
            <ul className="list-disc list-inside text-slate-400 space-y-3 text-lg font-light">
              <li>{features.comprehensiveData.description}</li>
              <li>{features.simpleIntegration.description}</li>
              <li>{features.usageManagement.description}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
