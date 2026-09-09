import { Card, CardContent } from "@lh/ui";
import { Link } from "react-router-dom";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";

export default function PrivacyPolicyPage() {
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
              Privacy Policy
            </h1>
            
            <div className="space-y-6 text-sm sm:text-base leading-relaxed">
              <p>
                <strong>Last Updated:</strong> September 2026
              </p>

              <p>
                At Limitless Hoops, we take your privacy seriously. This Privacy Policy explains how we collect, use, disclose, and safeguard your information when you visit our application.
              </p>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">1. Information We Collect</h2>
                <p>
                  We may collect personal information that you voluntarily provide to us when you register on the platform, express an interest in obtaining information about us or our products and services, or otherwise contact us. This includes names, phone numbers, and email addresses.
                </p>
              </section>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">2. How We Use Your Information</h2>
                <p>
                  We use the information we collect or receive to facilitate account creation and the logon process, to send administrative information to you, and to fulfill and manage your profile and team roster spots.
                </p>
              </section>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">3. Data Security</h2>
                <p>
                  We have implemented appropriate technical and organizational security measures designed to protect the security of any personal information we process. However, no electronic transmission over the internet or information storage technology can be guaranteed to be 100% secure.
                </p>
              </section>

              <section>
                <h2 className="text-xl font-bold text-white mb-2">4. Your Privacy Rights</h2>
                <p>
                  Depending on your location, you may have certain rights regarding your personal information, such as the right to request access to or deletion of your data. You may review, change, or terminate your account at any time.
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
