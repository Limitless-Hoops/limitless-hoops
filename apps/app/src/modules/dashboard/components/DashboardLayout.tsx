import React from "react";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";
import { Link } from "react-router-dom";

import { Footer } from "@lh/ui";

export default function DashboardLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen bg-background text-white flex flex-col overflow-x-hidden">
      {/* Top Navigation */}
      <header className="sticky top-0 z-40 w-full border-b border-white/10 bg-black/50 backdrop-blur-xl">
        <div className="container mx-auto flex h-16 items-center justify-between px-4 sm:px-8">
          <Link to="/dashboard" className="flex items-center gap-2 hover:opacity-80 transition-opacity">
            <img 
              src={LogoImage} 
              alt="Limitless Hoops Logo" 
              className="h-10 object-contain brightness-0 invert"
            />
          </Link>
          
          <nav className="flex items-center gap-6">
            <Link 
              to="/account" 
              className="h-8 w-8 rounded-full bg-white/10 flex items-center justify-center border border-white/20 cursor-pointer hover:bg-white/20 transition-colors"
            >
              <span className="text-xs font-bold">JS</span>
            </Link>
          </nav>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 container mx-auto px-4 sm:px-8 py-8 relative">
        {/* Subtle background glow for the dashboard */}
        <div className="absolute top-0 left-1/4 h-125 w-125 rounded-full bg-primary/10 blur-[120px] pointer-events-none" />
        
        <div className="relative z-10 pb-12">
          {children}
        </div>
      </main>

      <Footer />
    </div>
  );
}
