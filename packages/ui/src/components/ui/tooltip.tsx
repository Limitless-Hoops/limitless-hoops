import * as React from "react";
import { useState } from "react";

export function CustomTooltip({ 
  children, 
  content, 
  align = "center" 
}: { 
  children: React.ReactNode; 
  content: string; 
  align?: "center" | "right" | "left";
}) {
  const [active, setActive] = useState(false);
  
  let alignClass = "left-1/2 -translate-x-1/2";
  if (align === "right") alignClass = "right-0";
  if (align === "left") alignClass = "left-0";

  return (
    <div 
      className="relative inline-flex items-center"
      onMouseEnter={() => setActive(true)}
      onMouseLeave={() => setActive(false)}
      onClick={(e) => {
        e.stopPropagation();
        setActive(!active);
      }}
    >
      {children}
      <div 
        className={`absolute bottom-full mb-2 ${alignClass} w-max max-w-[220px] p-2.5 bg-neutral-900/95 backdrop-blur-sm border border-white/10 text-xs text-white leading-relaxed rounded-md shadow-xl transition-all duration-150 z-50 whitespace-normal text-left ${
          active ? 'opacity-100 visible translate-y-0' : 'opacity-0 invisible translate-y-1'
        }`}
      >
        {content}
      </div>
    </div>
  );
}
