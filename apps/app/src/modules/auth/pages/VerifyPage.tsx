import React, { useState, useRef } from "react";
import { Button, Card, CardContent } from "@lh/ui";
import { useSearchParams, useNavigate } from "react-router-dom";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";

export default function VerifyPage() {
  const [searchParams] = useSearchParams();
  const identifier = searchParams.get("id") || "your device";
  const navigate = useNavigate();

  const [otp, setOtp] = useState(["", "", "", "", "", ""]);
  const inputRefs = useRef<(HTMLInputElement | null)[]>([]);

  const handleChange = (index: number, value: string) => {
    if (!/^[0-9]*$/.test(value)) return;
    
    const newOtp = [...otp];
    newOtp[index] = value;
    setOtp(newOtp);

    // Auto-advance
    if (value !== "" && index < 5) {
      inputRefs.current[index + 1]?.focus();
    }
  };

  const handleKeyDown = (index: number, e: React.KeyboardEvent) => {
    if (e.key === "Backspace" && otp[index] === "" && index > 0) {
      inputRefs.current[index - 1]?.focus();
    }
  };

  const handleVerify = () => {
    const code = otp.join("");
    console.log("Verifying code:", code);
    // Mock successful login - route to onboarding for new user
    navigate("/auth/onboarding");
  };

  return (
    <div className="relative flex min-h-screen flex-col items-center justify-center overflow-hidden bg-background p-4 sm:p-8">
      {/* Subtle Premium Background Glow */}
      <div className="absolute top-[-10%] left-1/2 h-125 w-200 -translate-x-1/2 rounded-[100%] bg-primary/20 opacity-50 blur-[120px]" />

      <div className="z-10 flex w-full max-w-md flex-col items-center">
        {/* Brand Header */}
        <div className="mb-10 text-center flex flex-col items-center animate-in fade-in slide-in-from-bottom-4 duration-700">
          <img 
            src={LogoImage} 
            alt="Limitless Hoops Logo" 
            className="mb-8 w-48 object-contain drop-shadow-2xl sm:w-56 brightness-0 invert"
          />
        </div>

        {/* Glassmorphic Login Card */}
        <Card className="w-full border-white/10 bg-black/40 backdrop-blur-2xl shadow-2xl shadow-black/50 animate-in fade-in slide-in-from-bottom-6 duration-1000 delay-150 fill-mode-both">
          <CardContent className="pt-8 text-center px-6 sm:px-8">
            <div className="mx-auto mb-6 flex h-16 w-16 items-center justify-center rounded-full bg-white/5 border border-white/10 text-white shadow-inner">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="28"
                height="28"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <rect width="20" height="16" x="2" y="4" rx="2" />
                <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
              </svg>
            </div>
            
            <h1 className="mb-2 text-2xl font-extrabold tracking-tight text-white">
              Check your device
            </h1>
            <p className="mb-8 text-sm text-white/60 font-medium">
              We sent a 6-digit code to <br />
              <span className="font-bold text-white mt-1 block">{identifier}</span>
            </p>

            <div className="mb-8 flex justify-center gap-2 sm:gap-3">
              {otp.map((digit, idx) => (
                <input
                  key={idx}
                  ref={(el) => { inputRefs.current[idx] = el; }}
                  type="text"
                  inputMode="numeric"
                  maxLength={1}
                  value={digit}
                  onChange={(e) => handleChange(idx, e.target.value)}
                  onKeyDown={(e) => handleKeyDown(idx, e)}
                  className="h-12 w-10 sm:h-14 sm:w-12 rounded-lg border-white/10 bg-white/5 text-center text-lg sm:text-xl font-bold text-white focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary shadow-inner"
                />
              ))}
            </div>

            <Button
              className="group relative h-14 w-full overflow-hidden rounded-lg bg-linear-to-r from-primary to-red-800 text-base font-bold tracking-wide text-white shadow-lg shadow-primary/25 transition-all hover:scale-[1.02] hover:shadow-primary/40 disabled:opacity-50 disabled:hover:scale-100"
              onClick={handleVerify}
              disabled={otp.some((d) => d === "")}
            >
              <div className="absolute inset-0 flex h-full w-full justify-center transform-[skew(-12deg)_translateX(-100%)] group-hover:duration-1000 group-hover:transform-[skew(-12deg)_translateX(100%)]">
                <div className="relative h-full w-8 bg-white/20" />
              </div>
              <span>Verify Code</span>
            </Button>

            <div className="mt-8 text-sm font-medium text-white/40">
              Didn't receive it?{" "}
              <button className="text-white/80 underline hover:text-white transition-colors">
                Click to resend
              </button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
