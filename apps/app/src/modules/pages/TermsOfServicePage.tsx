import { Card, CardContent } from "@lh/ui";
import { Link } from "react-router-dom";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";

export default function TermsOfServicePage() {
  return (
    <div className="relative flex min-h-screen flex-col items-center justify-center overflow-hidden bg-background p-4 sm:p-8">
      <div className="absolute top-[-10%] left-1/2 h-125 w-200 -translate-x-1/2 rounded-[100%] bg-primary/20 opacity-30 blur-[120px]" />

      <div className="z-10 flex w-full max-w-4xl flex-col items-center pt-8 pb-12">
        <Link to="/" className="mb-10 text-center flex flex-col items-center animate-in fade-in slide-in-from-bottom-4 duration-700 hover:opacity-80 transition-opacity">
          <img 
            src={LogoImage} 
            alt="Limitless Hoops Logo" 
            className="w-48 object-contain drop-shadow-2xl sm:w-56 brightness-0 invert"
          />
        </Link>

        <Card className="w-full border-white/10 bg-black/40 backdrop-blur-2xl shadow-2xl shadow-black/50 animate-in fade-in slide-in-from-bottom-6 duration-1000 delay-150 fill-mode-both">
          <CardContent className="p-8 sm:p-12 text-white/80 prose prose-invert max-w-none">
            <h1 className="text-3xl font-extrabold uppercase tracking-tight text-white mb-8">
              Terms of Service
            </h1>
            
            <div className="space-y-6 text-sm sm:text-base leading-relaxed">
              <p>
                <strong>Last Updated:</strong> September 2026
              </p>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">1. Acceptance of Terms</h2>
                <p>
                  By accessing or using the Limitless Hoops platform, you agree to be bound by these Terms of Service. If you do not agree to all the terms and conditions, you may not access or use the service.
                </p>
              </section>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">2. User Accounts</h2>
                <p>
                  To use certain features, you must register for an account. You agree to provide accurate, current, and complete information and to keep your account information updated. You are responsible for all activities that occur under your account.
                </p>
              </section>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">3. Acceptable Use</h2>
                <p>
                  You agree not to use the platform for any unlawful purpose or in any way that interrupts, damages, or impairs the service. You are solely responsible for your conduct and any data you submit to the platform.
                </p>
              </section>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">4. Modification of Terms</h2>
                <p>
                  We reserve the right to modify these terms at any time. We will notify you of any material changes by posting the new terms on the platform. Your continued use of the service after such modifications will constitute acknowledgment and agreement of the modified terms.
                </p>
              </section>

              <div className="pt-8 border-t border-white/10 text-center text-sm text-white/50">
                <Link to="/login" className="underline hover:text-white transition-colors">
                  Return to Login
                </Link>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
