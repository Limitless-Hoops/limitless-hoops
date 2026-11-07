import { normalizePhoneForAPI } from "@lh/utils";
import { useState } from "react";
import { Card, CardContent } from "@lh/ui";
import { useSearchParams, useNavigate } from "react-router-dom";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";
import VerifyCodeForm from "@/modules/auth/components/VerifyCodeForm";
import { supabase } from "@/lib/supabase";
import { useAuth } from "@/modules/auth/contexts/AuthContext";

export default function VerifyPage() {
  const [searchParams] = useSearchParams();
  // M2: Remove dead ?email= fallback — only read ?id=
  const identifier = searchParams.get("id");
  const navigate = useNavigate();
  const { household } = useAuth();
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  // If there's no identifier in the query string, redirect back to login
  if (!identifier) {
    navigate("/login", { replace: true });
    return null;
  }

  const isEmail = identifier.includes("@");

  const handleVerify = async (code: string) => {
    setLoading(true);
    setError("");
    
    const verifyPayload: any = {
      token: code,
      type: isEmail ? "email" : "sms",
    };
    
    if (isEmail) {
      verifyPayload.email = identifier;
    } else {
      verifyPayload.phone = normalizePhoneForAPI(identifier);
    }

    const { error } = await supabase.auth.verifyOtp(verifyPayload);
    
    setLoading(false);

    if (error) {
      setError(error.message);
      return;
    }

    // H5: Navigate to dashboard for returning users, onboarding for new ones.
    // After verifyOtp, onAuthStateChange fires and AuthContext fetches the household.
    // Give it a moment, then check.
    if (household) {
      navigate("/dashboard", { replace: true });
    } else {
      navigate("/auth/onboarding", { replace: true });
    }
  };

  const handleResend = async () => {
    setError("");
    if (isEmail) {
      await supabase.auth.signInWithOtp({ email: identifier });
    } else {
      await supabase.auth.signInWithOtp({ phone: normalizePhoneForAPI(identifier) });
    }
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
          <CardContent className="pt-8 px-6 sm:px-8 pb-8">
            <VerifyCodeForm 
              identifier={identifier} 
              onVerify={handleVerify}
              onResend={handleResend}
              loading={loading}
            />
            {error && (
              <p className="mt-4 text-center text-sm font-medium text-red-400">
                {error}
              </p>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
