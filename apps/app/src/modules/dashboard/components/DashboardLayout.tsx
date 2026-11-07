import React from "react";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";
import { Link } from "react-router-dom";

import { Footer } from "@lh/ui";
import { useAuth } from "@/modules/auth/contexts/AuthContext";

export default function DashboardLayout({ children }: { children: React.ReactNode }) {
  const { household, signOut } = useAuth();
  
  const user = household?.users?.[0];
  const initials = user ? `${user.first_name?.[0] || ""}${user.last_name?.[0] || ""}`.toUpperCase() : "U";

  const handleSignOut = async () => {
    await signOut();
  };

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
          
          <nav className="flex items-center gap-6 relative group">
            <div className="h-8 w-8 rounded-full bg-white/10 flex items-center justify-center border border-white/20 cursor-pointer hover:bg-white/20 transition-colors">
              <span className="text-xs font-bold">{initials}</span>
            </div>

            {/* Dropdown Menu (Shows on hover) */}
            <div className="absolute right-0 top-full mt-2 w-48 rounded-md bg-zinc-900 border border-white/10 shadow-2xl opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-50">
              <div className="py-1">
                <Link
                  to="/account"
                  className="block px-4 py-2 text-sm text-white/80 hover:bg-white/10 hover:text-white"
                >
                  Profile
                </Link>
                <button
                  onClick={handleSignOut}
                  className="block w-full text-left px-4 py-2 text-sm text-red-400 hover:bg-white/10 hover:text-red-300"
                >
                  Sign Out
                </button>
              </div>
            </div>
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
