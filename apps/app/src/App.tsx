import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import LoginPage from "./modules/auth/pages/LoginPage";
import VerifyPage from "./modules/auth/pages/VerifyPage";
import OnboardingPage from "./modules/auth/pages/OnboardingPage";
import NotFound from "./modules/static/NotFound";
import TermsOfServicePage from "./modules/static/TermsOfServicePage";
import PrivacyPolicyPage from "./modules/static/PrivacyPolicyPage";
import DashboardPage from "./modules/dashboard/pages/DashboardPage";
import AccountPage from "./modules/account/pages/AccountPage";

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
        
        {/* Dashboard */}
        <Route path="/dashboard" element={<DashboardPage />} />
        <Route path="/account" element={<AccountPage />} />
      </Routes>
    </Router>
  );
}

export default App;
