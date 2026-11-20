import { normalizePhoneForAPI, isValidEmailOrPhone, formatPhone } from "@lh/utils";
import { useState } from "react";
import { Button, Input, Card, CardContent } from "@lh/ui";
import { useNavigate, Link } from "react-router-dom";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";
import { supabase } from "@/lib/supabase";

export default function LoginPage() {
  const [identifier, setIdentifier] = useState("");
  const isPhoneInput = /^\d{10}$/.test(identifier.replace(/\D/g, "")) && !/[a-zA-Z@]/.test(identifier);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSendCode = async (e: React.FormEvent) => {
    e.preventDefault();
    const val = identifier.trim();
    if (!val) {
      setError("Email address or phone number is required");
      return;
    }
    // M1: validate format before hitting Supabase
    if (!isValidEmailOrPhone(val)) {
      setError("Please enter a valid email address or US/Canadian phone number");
      return;
    }
    
    setLoading(true);
    let authError = null;

    if (val.includes("@")) {
      const { error } = await supabase.auth.signInWithOtp({ email: val });
      authError = error;
    } else {
      const { error } = await supabase.auth.signInWithOtp({ phone: normalizePhoneForAPI(val) });
      authError = error;
    }
    
    setLoading(false);

    if (authError) {
      setError(authError.message);
      return;
    }

    navigate(`/auth/verify?id=${encodeURIComponent(val)}`);
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
            
            {/* Social Auth Grid — 3 columns: Google, Facebook, Apple */}
            <div className="grid grid-cols-3 gap-3">
              {[ 
                { id: 'google', label: 'Google', icon: (
                  <svg className="h-4 w-4" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                    <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                    <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                    <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                  </svg>
                ), disabled: false },
                { id: 'facebook', label: 'Facebook', icon: (
                  <svg className="h-4 w-4" viewBox="0 0 24 24" fill="#1877F2" xmlns="http://www.w3.org/2000/svg">
                    <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.469h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.469h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                  </svg>
                ), disabled: false },
                { id: 'apple', label: 'Apple', icon: (
                  <svg className="h-4 w-4" viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                    <path d="M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.546 9.103 1.519 12.09 1.013 1.454 2.208 3.09 3.792 3.039 1.52-.065 2.09-.987 3.935-.987 1.831 0 2.35.987 3.96.948 1.637-.026 2.62-1.458 3.579-2.91 1.163-1.696 1.64-3.342 1.66-3.428-.039-.013-3.182-1.222-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.502 1.09zM15.53 3.83c.843-1.012 1.4-2.427 1.245-3.83-1.207.052-2.662.805-3.532 1.818-.78.896-1.454 2.338-1.273 3.714 1.338.104 2.715-.688 3.56-1.702z"/>
                  </svg>
                ), disabled: true }
              ].map(provider => (
                <Button
                  key={provider.id}
                  type="button"
                  variant="outline"
                  aria-disabled={provider.disabled}
                  onClick={() => {
                    if (!provider.disabled) {
                      supabase.auth.signInWithOAuth({ provider: provider.id as any, options: { redirectTo: window.location.origin + "/dashboard" }});
                    }
                  }}
                  className={`group relative h-12 w-full border border-white/10 bg-black transition-all font-semibold shadow-sm overflow-hidden ${provider.disabled ? 'text-white/40 cursor-not-allowed' : 'text-white hover:bg-white/10'}`}
                >
                  <div className={`flex items-center justify-center gap-2 transition-transform duration-300 ${provider.disabled ? 'group-hover:-translate-y-12' : ''}`}>
                    {provider.icon}
                    {provider.label}
                  </div>
                  {provider.disabled && (
                    <div className="absolute inset-0 flex items-center justify-center translate-y-12 transition-transform duration-300 group-hover:translate-y-0 text-white/70">
                      Unavailable
                    </div>
                  )}
                </Button>
              ))}
            </div>

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
                  className="h-14 text-base"
                  value={identifier}
                  error={error}
                  prefixText={isPhoneInput ? "+1" : undefined}
                  spellCheck={false}
                  autoComplete="username"
                  onChange={(e) => {
                    const val = e.target.value;
                    const hasLettersOrAt = /[a-zA-Z@]/.test(val);
                    
                    if (hasLettersOrAt) {
                      setIdentifier(val);
                      if (error) setError("");
                      return;
                    }
                    
                    let stripped = val.replace(/[\(\)\-\s]/g, "");
                    if (stripped.startsWith("+1")) {
                      stripped = stripped.replace("+1", "1");
                    }
                    
                    if (/^\d+$/.test(stripped)) {
                      if (stripped.length === 11 && stripped.startsWith("1")) {
                        stripped = stripped.substring(1);
                      }
                      if (stripped.length === 10) {
                        setIdentifier(formatPhone(stripped));
                      } else {
                        setIdentifier(stripped);
                      }
                    } else {
                      setIdentifier(stripped);
                    }
                    
                    if (error) setError("");
                  }}
                />
                <p className="text-[11px] text-white/40 pt-1">Only US/Canadian numbers are allowed.</p>
              </div>
              <Button
                type="submit"
                size="lg"
                className="w-full"
                disabled={loading}
              >
                {loading ? "Sending Code..." : "Send Login Code"}
              </Button>
            </form>
          </CardContent>
        </Card>
        
        <p className="mt-8 max-w-sm text-center text-xs font-medium text-white/40 sm:text-sm animate-in fade-in duration-1000 delay-300 fill-mode-both">
          By continuing, you agree to our{" "}
          <Link to="/terms" className="cursor-pointer text-white/60 underline underline-offset-2 transition-colors hover:text-white">
            Terms of Service
          </Link>{" "}
          and{" "}
          <Link to="/privacy" className="cursor-pointer text-white/60 underline underline-offset-2 transition-colors hover:text-white">
            Privacy Policy
          </Link>.
        </p>
      </div>
    </div>
  );
}
