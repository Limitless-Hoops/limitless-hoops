import * as React from "react";

export function ListCard({ children, className = "" }: { children: React.ReactNode, className?: string }) {
  return (
    <div className={`flex flex-col sm:flex-row sm:items-center justify-between bg-white/5 border border-white/10 rounded-lg p-4 gap-3 ${className}`}>
      {children}
    </div>
  );
}
