import React, { useState } from "react";
import { Button, Input, Card, CardContent } from "@lh/ui";
import { useNavigate, Link } from "react-router-dom";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";
import { isValidEmailOrPhone } from "@lh/utils";

export default function LoginPage() {
  const [identifier, setIdentifier] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSendCode = (e: React.SubmitEvent) => {
    e.preventDefault();
    if (!identifier.trim()) {
      setError("Email address or phone number is required");
      return;
    }
    if (!isValidEmailOrPhone(identifier)) {
      setError("Please enter a valid email or 10-digit phone number");
      return;
    }
    navigate(`/auth/verify?id=${encodeURIComponent(identifier)}`);
  };

  const handleGoogleSignIn = () => {
    console.log("Trigger Google OAuth");
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
          <h1 className="text-3xl font-extrabold uppercase tracking-tight text-white sm:text-4xl">
            BE YOUR OWN HERO.
          </h1>
          <p className="mt-3 text-sm text-white/60 sm:text-base font-medium">
            Sign in or create an account to manage your household.
          </p>
        </div>

        {/* Glassmorphic Login Card */}
        <Card className="w-full border-white/10 bg-black/40 backdrop-blur-2xl shadow-2xl shadow-black/50 animate-in fade-in slide-in-from-bottom-6 duration-1000 delay-150 fill-mode-both">
          <CardContent className="p-6 sm:p-8">
            {/* Social Auth */}
            <Button
              variant="outline"
              className="h-14 w-full border border-white/10 bg-black text-white transition-all hover:bg-white/10 font-semibold shadow-sm flex items-center justify-center"
              onClick={handleGoogleSignIn}
            >
              <svg
                className="mr-3 h-5 w-5"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                  fill="#4285F4"
                />
                <path
                  d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                  fill="#34A853"
                />
                <path
                  d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                  fill="#FBBC05"
                />
                <path
                  d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                  fill="#EA4335"
                />
              </svg>
              Continue with Google
            </Button>

            {/* Divider */}
            <div className="relative my-8">
              <div className="absolute inset-0 flex items-center">
                <div className="w-full border-t border-white/10" />
              </div>
              <div className="relative flex justify-center text-xs uppercase tracking-wider font-semibold">
                <span className="bg-[#0f0f11] px-3 text-white/40">
                  Or continue with code
                </span>
              </div>
            </div>

            {/* Magic OTP Form */}
            <form onSubmit={handleSendCode} noValidate className="space-y-5">
              <div className="space-y-2">
                <Input
                  type="text"
                  placeholder="Email address or phone number"
                  className="h-14 bg-white/5 px-4 text-base text-white placeholder:text-white/30 focus:ring-1 transition-colors border-white/10 focus:border-primary focus:ring-primary"
                  value={identifier}
                  error={error}
                  onChange={(e) => {
                    setIdentifier(e.target.value);
                    if (error) setError("");
                  }}
                />
              </div>
              <Button
                type="submit"
                className="group relative h-14 w-full overflow-hidden rounded-lg bg-linear-to-r from-primary to-red-800 text-base font-bold tracking-wide text-white shadow-lg shadow-primary/25 transition-all hover:scale-[1.02] hover:shadow-primary/40"
              >
                <div className="absolute inset-0 flex h-full w-full justify-center transform-[skew(-12deg)_translateX(-100%)] group-hover:duration-1000 group-hover:transform-[skew(-12deg)_translateX(100%)]">
                  <div className="relative h-full w-8 bg-white/20" />
                </div>
                <span>Send Login Code</span>
              </Button>
            </form>
          </CardContent>
        </Card>
        
        <p className="mt-8 max-w-sm text-center text-xs font-medium text-white/40 sm:text-sm animate-in fade-in duration-1000 delay-300 fill-mode-both">
          By continuing, you agree to our{" "}
          <Link to="/terms" className="text-white/60 underline underline-offset-2 transition-colors hover:text-white">
            Terms of Service
          </Link>{" "}
          and{" "}
          <Link to="/privacy" className="text-white/60 underline underline-offset-2 transition-colors hover:text-white">
            Privacy Policy
          </Link>.
        </p>
      </div>
    </div>
  );
}
