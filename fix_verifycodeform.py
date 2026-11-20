import re

with open('apps/app/src/modules/auth/components/VerifyCodeForm.tsx', 'r') as f:
    content = f.read()

# Add useEffect import
content = content.replace('import { useState, useRef } from "react";', 'import { useState, useRef, useEffect } from "react";')

# Replace the two useStates with a single useEffect
old_timer = """  // Countdown timer
  useState(() => {
    if (cooldown > 0) {
      const timer = setTimeout(() => setCooldown(c => c - 1), 1000);
      return () => clearTimeout(timer);
    }
  });

  // Simpler effect using the existing pattern
  useState(() => {
    const interval = setInterval(() => {
      setCooldown(c => {
        if (c <= 0) { clearInterval(interval); return 0; }
        return c - 1;
      });
    }, 1000);
    return () => clearInterval(interval);
  });"""

new_timer = """  useEffect(() => {
    if (cooldown <= 0) return;
    const timer = setInterval(() => {
      setCooldown(c => c - 1);
    }, 1000);
    return () => clearInterval(timer);
  }, [cooldown]);"""

content = content.replace(old_timer, new_timer)

with open('apps/app/src/modules/auth/components/VerifyCodeForm.tsx', 'w') as f:
    f.write(content)

