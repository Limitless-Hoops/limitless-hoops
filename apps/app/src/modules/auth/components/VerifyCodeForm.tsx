import { useState, useRef, useEffect } from "react";
import { Button } from "@lh/ui";

interface VerifyCodeFormProps {
  // L17: make identifier required — all call sites always pass it
  identifier: string;
  onVerify: (code: string) => void;
  onResend?: () => void;
  loading?: boolean;
}

export default function VerifyCodeForm({ identifier, onVerify, onResend, loading }: VerifyCodeFormProps) {
  const [otp, setOtp] = useState(["", "", "", "", "", ""]);
  const inputRefs = useRef<(HTMLInputElement | null)[]>([]);

  // M4/L21: determine whether this is a phone or email verification
  const isPhone = !identifier.includes("@");

  const handlePaste = (e: React.ClipboardEvent) => {
    e.preventDefault();
    const pastedData = e.clipboardData.getData("text/plain").replace(/[^0-9]/g, "").slice(0, 6);
    if (!pastedData) return;

    const newOtp = [...otp];
    for (let i = 0; i < pastedData.length; i++) {
      newOtp[i] = pastedData[i];
    }
    setOtp(newOtp);

    const focusIndex = Math.min(pastedData.length, 5);
    inputRefs.current[focusIndex]?.focus();
  };

  const handleChange = (index: number, value: string) => {
    if (!/^[0-9]*$/.test(value)) return;
    
    const newOtp = [...otp];
    newOtp[index] = value;
    setOtp(newOtp);

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
    onVerify(code);
  };

  const [cooldown, setCooldown] = useState(60);

  useEffect(() => {
    if (cooldown <= 0) return;
    const timer = setInterval(() => {
      setCooldown(c => c - 1);
    }, 1000);
    return () => clearInterval(timer);
  }, [cooldown]);

  const handleResendClick = () => {
    if (cooldown > 0 || !onResend) return;
    onResend();
    setCooldown(60);
  };

  // M4/L21: Phone icon for SMS, envelope icon for email
  const PhoneIcon = () => (
    <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"
      fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07A19.5 19.5 0 0 1 4.69 12 19.79 19.79 0 0 1 1.62 3.44 2 2 0 0 1 3.6 1.27h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L7.91 8.96a16 16 0 0 0 6.13 6.13l.96-.96a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/>
    </svg>
  );

  const EmailIcon = () => (
    <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"
      fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <rect width="20" height="16" x="2" y="4" rx="2" />
      <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
    </svg>
  );

  return (
    <div className="text-center w-full">
      <div className="mx-auto mb-6 flex h-16 w-16 items-center justify-center rounded-full bg-white/5 border border-white/10 text-white shadow-inner">
        {isPhone ? <PhoneIcon /> : <EmailIcon />}
      </div>
      
      {/* M4/L21: heading matches the contact type */}
      <h1 className="mb-2 text-2xl font-extrabold tracking-tight text-white">
        {isPhone ? "Check your phone" : "Check your inbox"}
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
            onPaste={handlePaste}
            className="h-12 w-10 sm:h-14 sm:w-12 rounded-lg border border-white/10 bg-white/5 text-center text-lg sm:text-xl font-bold text-white focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary shadow-inner"
          />
        ))}
      </div>

      <Button
        className="w-full"
        onClick={handleVerify}
        disabled={otp.some((d) => d === "") || loading}
      >
        {loading ? "Verifying..." : "Verify Code"}
      </Button>

      <div className="mt-8 text-sm font-medium text-white/40">
        Didn't receive it?{" "}
        {cooldown > 0 ? (
          <span className="text-white/40">Resend in {cooldown}s</span>
        ) : (
          <button 
            onClick={handleResendClick}
            disabled={!onResend}
            className={`cursor-pointer text-white/80 underline hover:text-white transition-colors ${!onResend && "opacity-50 cursor-not-allowed"}`}
          >
            Click to resend
          </button>
        )}
      </div>
    </div>
  );
}
