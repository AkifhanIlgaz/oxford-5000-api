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
            Oxford 5000™ API
          </h1>
          <div className="space-y-6">
            <p className="text-2xl text-slate-300 font-display">
              Your Gateway to Advanced English Vocabulary
            </p>
            <p className="text-slate-400 text-lg font-light leading-relaxed">
              Access comprehensive data for the Oxford 5000™ word list,
              including definitions, examples, and CEFR levels. Perfect for
              educators, developers, and language learners.
            </p>
            <ul className="list-disc list-inside text-slate-400 space-y-3 text-lg font-light">
              <li>Complete Oxford 5000™ word database</li>
              <li>CEFR level classifications</li>
              <li>Usage examples and definitions</li>
              <li>Regular updates and support</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
