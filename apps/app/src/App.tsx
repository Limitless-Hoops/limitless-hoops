import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import LoginPage from "./modules/auth/pages/LoginPage";
import VerifyPage from "./modules/auth/pages/VerifyPage";
import OnboardingPage from "./modules/auth/pages/OnboardingPage";
import NotFound from "./modules/pages/NotFound";
import TermsOfServicePage from "./modules/pages/TermsOfServicePage";
import PrivacyPolicyPage from "./modules/pages/PrivacyPolicyPage";

function App() {
  return (
    <Router>
      <Routes>
        {/* Default route redirect to login for now */}
        <Route path="/" element={<Navigate to="/login" replace />} />
        <Route path="*" element={<NotFound />} />
        
        {/* Legal */}
        <Route path="/terms" element={<TermsOfServicePage />} />
        <Route path="/privacy" element={<PrivacyPolicyPage />} />

        {/* Auth Module */}
        <Route path="/login" element={<LoginPage />} />
        <Route path="/auth/verify" element={<VerifyPage />} />
        <Route path="/auth/onboarding" element={<OnboardingPage />} />
        
        {/* Mock Dashboard */}
        <Route path="/dashboard" element={
          <div className="flex min-h-screen flex-col items-center justify-center bg-background text-foreground">
            <div className="mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-primary text-3xl shadow-lg">
              🏀
            </div>
            <h1 className="mt-4 text-3xl font-bold uppercase">Dashboard</h1>
            <p className="text-muted-foreground mt-2">Authentication successful. You are in.</p>
          </div>
        } />
      </Routes>
    </Router>
  );
}

export default App;
