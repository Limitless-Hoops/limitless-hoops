import React from 'react';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
}

export function Button({ children, ...props }: ButtonProps) {
  return (
    <button 
      className="h-12 px-6 rounded-md text-white text-lg font-bold uppercase tracking-wide bg-brand-accent shadow-[0_2px_4px_rgba(45,35,66,0.4),0_7px_13px_-3px_rgba(45,35,66,0.3),inset_0_-3px_0_rgba(58,65,111,0.5)] transition-all hover:-translate-y-0.5 active:translate-y-0.5"
      {...props}
    >
      {children}
    </button>
  );
}